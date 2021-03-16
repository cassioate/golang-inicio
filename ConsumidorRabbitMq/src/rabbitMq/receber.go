package rabbitMq

import (
	"fmt"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func ReceberMensagem() error {

	url := os.Getenv("AMQP_URL")
	connection, erro := amqp.Dial(url)
	retornaErro(erro)
	defer connection.Close()

	channel, erro := connection.Channel()
	retornaErro(erro)
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
	retornaErro(erro)

	for m := range messages {
		//Retorna o Body do struct acima que é um []byte
		SalvarArquivoTxt(m.Body)
		fmt.Println("Arquivo recebido")
	}

	return erro
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
