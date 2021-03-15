package rabbitMq

import (
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func ReceberMensagem() ([]byte, error) {

	url := os.Getenv("AMQP_URL")

	connection, erro := amqp.Dial(url)
	if erro != nil {
		return nil, erro
	}
	defer connection.Close()

	channel, erro := connection.Channel()
	if erro != nil {
		return nil, erro
	}

	defer channel.Close()

	messages, erro := channel.Consume(
		"Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if erro != nil {
		return nil, erro
	}

	//Recebe um Struct que contem um BODY
	message := <-messages
	SalvarArquivoTxt(message.Body)

	//Retorna o Body do struct acima que é um []byte
	return message.Body, erro
}

func SalvarArquivoTxt(message []byte) error {

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

	if erro != nil {
		return erro
	}

	arquivo.Write(message)
	arquivo.Close()
	return erro
}

func blockForever() {
	select {}
}
