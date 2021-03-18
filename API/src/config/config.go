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
	StringConexaoBanco     = ""
	StringConexaoBancoTest = ""
	// Porta da API rodando
	Porta = 0
)

// Iinicia as variaveis de ambiente e define a porta a ser utilizada, também da inicio ao carregamento das tabelas pelo GoSideKick.
// Esse metodo também possui a obrigação de enviar a primeira mensagem para a MENSAGERIA, para que assim possa ativar uma fila antes
// que o serviço que irá receber a mensagem se conecte, permitindo assim que ele ingresse na fila corretamente do outro lado.
// Evitando que o receptor da mensagem final não consiga ingressar.
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

// Carrega as tabelas utilizando o GoSideKick/Migration, informa também quantos e quais foram os arquivos utilizados.
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
