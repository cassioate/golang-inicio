package repositorios

import (
	"database/sql"
	"fmt"
	"modulo/src/model"
)

// Representa um repositorio de cliente
type Cliente struct {
	db *sql.DB
}

// Cria um repositorio de cliente
func NovoRepositorioDeCliente(db *sql.DB) *Cliente {
	return &Cliente{db}
}

// Insere um cliente no banco de dados e retorna o ID dele
func (repositorio Cliente) Criar(cliente model.Cliente) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"insert into cliente (uuid, nome, endereco) values($1, $2, $3) RETURNING uuid",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	var id uint64
	err := statement.QueryRow(cliente.Uuid, cliente.Nome, cliente.Endereco).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, erro
	}

	return id, erro
}

// Busca todos os clientes
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

// Busca os clientes por ID
func (repositorio Cliente) BuscarClientePorId(Id uint64) (model.Cliente, error) {

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

func (repositorio Cliente) Atualizar(Id uint64, cliente model.Cliente) error {

	statement, erro := repositorio.db.Prepare(
		`update cliente set nome = $1, endereco = $2, atualizado_em = to_char(current_timestamp, 'DD/MM/YYYY HH24:MI:SS') where uuid = $3`,
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, err := statement.Exec(cliente.Nome, cliente.Endereco, Id)
	if err != nil {
		return erro
	}
	return erro
}

func (repositorio Cliente) Deletar(Id uint64) error {

	statement, erro := repositorio.db.Prepare(
		"delete from cliente where uuid = $1",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, err := statement.Exec(Id)
	if err != nil {
		return erro
	}

	return erro
}
