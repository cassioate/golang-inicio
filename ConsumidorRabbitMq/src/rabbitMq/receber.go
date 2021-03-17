package rabbitMq

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

// Irá estabelecer uma conexão com a MENSAGERIA, então conectará um canal que consome uma Queue(fila)
// sendo repassada para um for onde toda as mensages recebidas serão interceptadas, em caso de erro elá irá retorna-lo.
// Uma vez iniciado, ele irá permanecer com a fila aberta recebendo todas as novas mensages em tempo real.
func ReceberMensagem() error {

	url := os.Getenv("AMQP_URL")

	connection, erro := conectar(url)
	defer connection.Close()

	channel, erro := connection.Channel()
	defer channel.Close()
	retornaErro(erro)

	messages, erro := buscarMensagem(channel)
	retornaErro(erro)

	for m := range messages {
		//Retorna o Body do struct acima que é um []byte
		SalvarArquivoJson(m.Body)
		fmt.Println("Arquivo recebido")
	}

	return erro
}

// Estabelece conexão com a MENSAGERIA e retorna um erro em caso negativo.
func conectar(url string) (*amqp.Connection, error) {
	connection, erro := amqp.Dial(url)
	for connection == nil {
		fmt.Println("Aguardando conexão com RabbitMQ")
		time.Sleep(time.Second * 5)
		connection, erro = amqp.Dial(url)
	}
	return connection, erro
}

// Busca a fila de mensagens para que possa ser utilizada.
func buscarMensagem(channel *amqp.Channel) (<-chan amqp.Delivery, error) {
	messages, erro := channel.Consume(
		"Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for messages == nil {
		fmt.Println("Aguardando menssagens")
		time.Sleep(time.Second * 5)
		messages, erro = channel.Consume(
			"Queue",
			"",
			true,
			false,
			false,
			false,
			nil,
		)
	}
	return messages, erro
}

// Salva a mensagem recebida em um arquivo .Json. É definido o diretorio em que será salvo o arquivo,
// então é verificado se existe algum arquivo com o nome "cliente" caso exista irá criar um "cliente_1",
// caso exista irá criar "cliente_2" e assim por diante, se utilizando de um for para fazer tal verificação,
// após isso irá salvar o arquivo no diretorio escolhido.
// Em caso de erro irá retorna-lo.
func SalvarArquivoJson(message []byte) error {

	os.Chdir(os.Getenv("NOVOS_CLIENTES"))
	fileInfo, erro := os.Stat("cliente.json")
	var arquivo *os.File
	var novoNomeProArquivo string

	if fileInfo != nil {
		var i int64
		i = 0
		for fileInfo != nil {
			novoNomeProArquivo = "cliente_" + strconv.FormatInt(i, 10) + ".json"
			i++
			fileInfo, _ = os.Stat(novoNomeProArquivo)
		}
		arquivo, erro = os.OpenFile(novoNomeProArquivo, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	} else {
		arquivo, erro = os.OpenFile("cliente.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	}
	retornaErro(erro)
	arquivo.Write(message)
	arquivo.Close()
	return erro
}

// Simplificação de um retorno de erro muito utilzado
func retornaErro(erro error) error {
	if erro != nil {
		return erro
	}
	return nil
}
