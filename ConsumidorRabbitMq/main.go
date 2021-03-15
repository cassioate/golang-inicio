package main

import (
	"fmt"
	"modulo/src/config"
	"modulo/src/rabbitMq"
	"net/http"
)

func main() {
	fmt.Println("Iniciou")
	config.Carregar()
	rabbitMq.ReceberMensagem()
	print("to aqui")
	http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), nil)
}
