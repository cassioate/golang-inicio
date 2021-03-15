package main

import (
	"fmt"
	"log"
	"modulo/src/config"
	"modulo/src/router"
	"net/http"
)

func main() {
	fmt.Println("Rodando a API")
	config.Carregar()
	r := router.Gerar()
	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
