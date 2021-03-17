package repositorios

import (
	"database/sql"
	"encoding/json"
	"modulo/src/model"
	"os"

	"github.com/streadway/amqp"
)

// Representa um repositorio de cliente
type Cliente struct {
	db *sql.DB
}

// Cria um repositorio de cliente
func NovoRepositorioDeCliente(db *sql.DB) *Cliente {
	return &Cliente{db}
}

// Tem a função de realizar a inserção de um Cliente no banco de dados utilizando a Query preparada.
// Também irá realizar a chamada do metodo que enviará o arquivo que foi salvo no banco de dados em formato
// Json para a MENSAGERIA
func (repositorio Cliente) Criar(cliente model.Cliente) (string, error) {

	statement, erro := repositorio.db.Prepare(
		"insert into cliente (uuid, nome, endereco) values($1, $2, $3) RETURNING uuid, nome, endereco, cadastrado_em, atualizado_em",
	)
	if erro != nil {
		return "", erro
	}

	defer statement.Close()

	var clienteRetornado model.Cliente
	erro = statement.QueryRow(cliente.Uuid, cliente.Nome, cliente.Endereco).Scan(&clienteRetornado.Uuid, &clienteRetornado.Nome,
		&clienteRetornado.Endereco, &clienteRetornado.Cadastrado_em, &clienteRetornado.Atualizado_em)
	if erro != nil {
		return "", erro
	}

	erro = EnviarRabbitMQ(clienteRetornado)
	if erro != nil {
		return "", erro
	}

	return clienteRetornado.Uuid, erro
}

// Tem a função de realizar a Busca todos os Clientes no banco de dados utilizando a Query preparada.
func (repositorio Cliente) BuscarTodosClientes() ([]model.Cliente, error) {

	linhas, erro := repositorio.db.Query(
		"select * from cliente",
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var clientes []model.Cliente

	for linhas.Next() {
		var cliente model.Cliente

		if erro = linhas.Scan(
			&cliente.Uuid,
			&cliente.Nome,
			&cliente.Endereco,
			&cliente.Cadastrado_em,
			&cliente.Atualizado_em,
		); erro != nil {
			return nil, erro
		}

		clientes = append(clientes, cliente)
	}

	return clientes, erro
}

// Tem a função de realizar a Busca por ID no banco de dados utilizando a Query preparada.
func (repositorio Cliente) BuscarClientePorId(Id string) (model.Cliente, error) {

	linha, erro := repositorio.db.Query(
		"select * from cliente where uuid = $1",
		Id,
	)
	if erro != nil {
		return model.Cliente{}, erro
	}
	defer linha.Close()

	var cliente model.Cliente
	if linha.Next() {
		if erro = linha.Scan(
			&cliente.Uuid,
			&cliente.Nome,
			&cliente.Endereco,
			&cliente.Cadastrado_em,
			&cliente.Atualizado_em,
		); erro != nil {
			return model.Cliente{}, erro
		}
	}

	return cliente, erro
}

// Tem a função de realizar a Atualização no banco de dados utilizando a Query preparada.
func (repositorio Cliente) Atualizar(Id string, cliente model.Cliente) error {

	statement, erro := repositorio.db.Prepare(
		`update cliente set nome = $1, endereco = $2, atualizado_em = to_char(current_timestamp, 'DD/MM/YYYY HH24:MI:SS') where uuid = $3`,
	)
	seErroRepositorio(erro)
	defer statement.Close()

	_, erro = statement.Exec(cliente.Nome, cliente.Endereco, Id)
	seErroRepositorio(erro)

	return erro
}

// Tem a função de realizar a Deleção no banco de dados utilizando a Query preparada.
func (repositorio Cliente) Deletar(Id string) error {

	statement, erro := repositorio.db.Prepare(
		"delete from cliente where uuid = $1",
	)
	seErroRepositorio(erro)

	defer statement.Close()

	_, erro = statement.Exec(Id)
	seErroRepositorio(erro)

	return erro
}

// Esse metodo possui a obrigação de enviar a mensagem para a MENSAGERIA,
// permitindo que a fila seja criada e utilizada. Nele é estabelecida uma conexão,
// um canal e uma fila, ao final ele converte o objeto para um Json e depois publica
// ele na MENSAGERIA.
func EnviarRabbitMQ(cliente model.Cliente) error {

	url := os.Getenv("AMQP_URL")

	connection, erro := amqp.Dial(url)
	seErroRepositorio(erro)
	defer connection.Close()

	channel, erro := connection.Channel()
	seErroRepositorio(erro)
	defer channel.Close()

	_, erro = channel.QueueDeclare(
		"Queue",
		false,
		false,
		false,
		false,
		nil,
	)

	json, erro := json.Marshal(cliente)
	seErroRepositorio(erro)

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(json),
	}

	erro = channel.Publish("", "Queue", false, false, message)

	return erro
}

// Tem a função de encurtar o codigo em caso de erro no controller
func seErroRepositorio(erro error) error {
	if erro != nil {
		return erro
	}
	return nil
}
