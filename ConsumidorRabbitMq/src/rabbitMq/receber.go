package rabbitMq

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

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

func conectar(url string) (*amqp.Connection, error) {
	connection, erro := amqp.Dial(url)
	for connection == nil {
		fmt.Println("Aguardando conexão com RabbitMQ")
		time.Sleep(time.Second * 5)
		connection, erro = amqp.Dial(url)
	}
	return connection, erro
}

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

func retornaErro(erro error) error {
	if erro != nil {
		return erro
	}
	return nil
}
