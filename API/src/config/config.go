package config

import (
	"context"
	"fmt"
	"log"
	"modulo/src/model"
	"modulo/src/repositorios"
	"os"
	"strconv"

	"github.com/gosidekick/migration/v3"
	"github.com/joho/godotenv"
)

var (
	// conexão do banco
	StringConexaoBanco = ""
	// Porta da API rodando
	Porta = 0
)

//vai iniciar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_NOME"),
	)

	err := CarregarTabelas()
	if err != nil {
		log.Fatal(err)
	}

	repositorios.EnviarRabbitMQ(model.Cliente{})
}

func CarregarTabelas() error {
	quantidade, listaDeTabelasUtilizadas, err := migration.Run(context.Background(), "./banco/fixtures", StringConexaoBanco, "up")
	if err != nil {
		return err
	}
	if quantidade > 0 {
		fmt.Printf("Foram realizadas: %d\n", quantidade)
		fmt.Printf("Foram utilizadas as seguintes tabelas: %s\n", listaDeTabelasUtilizadas)
	}
	return nil
}
