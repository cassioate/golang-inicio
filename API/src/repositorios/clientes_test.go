package repositorios

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"modulo/src/model"
	"modulo/src/util"
	"os"
	"testing"

	"github.com/gosidekick/migration/v3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestCriar(t *testing.T) {
	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	conexao := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USUARIO_TEST"),
		os.Getenv("DB_SENHA_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_PORTA_TEST"),
		os.Getenv("DB_NOME_TEST"),
	)

	db, err := sql.Open("postgres", conexao)

	migration.Run(context.Background(), "./banco/fixtures", conexao, "up")

	if err != nil {
		t.Fatalf("Não conseguiu criar um DATABASE %v", err)
	}

	repositorio := NovoRepositorioDeCliente(db)

	cliente := model.Cliente{
		Uuid:          util.RandomString(10),
		Nome:          util.RandomString(10),
		Endereco:      util.RandomString(20),
		Cadastrado_em: "18/03/2020 10:30:30",
		Atualizado_em: "18/03/2020 10:30:30",
	}

	clienteID, err := repositorio.Criar(cliente)

	if err != nil {
		t.Fatalf("Não conseguiu criar um cliente | erro: %v", err)
	}

	if clienteID != cliente.Uuid {
		t.Fatalf("o id encontrado foi: %s, mas o ID esperado era: %s", clienteID, cliente.Uuid)
	}

	if clienteID == "" {
		t.Fatalf("o id encontrado foi: %s, mas o ID esperado era: %s", clienteID, cliente.Uuid)
	}

	defer db.Close()

}

func TestBuscarTodosClientes(t *testing.T) {

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	conexao := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USUARIO_TEST"),
		os.Getenv("DB_SENHA_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_PORTA_TEST"),
		os.Getenv("DB_NOME_TEST"),
	)

	db, err := sql.Open("postgres", conexao)

	migration.Run(context.Background(), "./banco/fixtures", conexao, "up")

	if err != nil {
		t.Fatalf("Não conseguiu criar um DATABASE %v", err)
	}

	repositorio := NovoRepositorioDeCliente(db)

	_, err = repositorio.BuscarTodosClientes()

	if err != nil {
		t.Fatalf("Não conseguiu criar um cliente | erro: %v", err)
	}

	defer db.Close()

}

func TestBuscarClientePorId(t *testing.T) {

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	conexao := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USUARIO_TEST"),
		os.Getenv("DB_SENHA_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_PORTA_TEST"),
		os.Getenv("DB_NOME_TEST"),
	)

	db, err := sql.Open("postgres", conexao)

	migration.Run(context.Background(), "./banco/fixtures", conexao, "up")

	if err != nil {
		t.Fatalf("Não conseguiu criar um DATABASE %v", err)
	}

	cliente := model.Cliente{
		Uuid:          util.RandomString(10),
		Nome:          util.RandomString(10),
		Endereco:      util.RandomString(20),
		Cadastrado_em: "18/03/2020 10:30:30",
		Atualizado_em: "18/03/2020 10:30:30",
	}

	repositorio := NovoRepositorioDeCliente(db)

	clienteID, err := repositorio.Criar(cliente)

	if err != nil {
		t.Fatalf("Não conseguiu criar um cliente | erro: %v", err)
	}

	clienteRetornado, err := repositorio.BuscarClientePorId(cliente.Uuid)

	if err != nil {
		t.Fatalf("Não conseguiu buscar um cliente | erro: %v", err)
	}

	if clienteID != cliente.Uuid {
		t.Fatalf("Cliente encontrado diferente do cliente enviado | erro: %v", err)
	}

	if cliente.Uuid != clienteRetornado.Uuid ||
		cliente.Nome != clienteRetornado.Nome ||
		cliente.Endereco != clienteRetornado.Endereco {
		t.Fatalf("Erro no teste de comparação entre os clientes | erro: %v", err)
	}

	defer db.Close()

}

func TestAtualizar(t *testing.T) {

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	conexao := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USUARIO_TEST"),
		os.Getenv("DB_SENHA_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_PORTA_TEST"),
		os.Getenv("DB_NOME_TEST"),
	)

	db, err := sql.Open("postgres", conexao)

	migration.Run(context.Background(), "./banco/fixtures", conexao, "up")

	if err != nil {
		t.Fatalf("Não conseguiu criar um DATABASE %v", err)
	}

	cliente := model.Cliente{
		Uuid:     util.RandomString(10),
		Nome:     util.RandomString(10),
		Endereco: util.RandomString(20),
	}

	repositorio := NovoRepositorioDeCliente(db)

	clienteID, err := repositorio.Criar(cliente)

	if err != nil {
		t.Fatalf("Não conseguiu criar um cliente | erro: %v", err)
	}

	clienteRetornado, err := repositorio.BuscarClientePorId(clienteID)

	if err != nil {
		t.Fatalf("Não conseguiu retornar um cliente | erro: %v", err)
	}

	err = repositorio.Atualizar(cliente.Uuid, clienteRetornado)

	if err != nil {
		t.Fatalf("Não conseguiu criar um cliente | erro: %v", err)
	}

	defer db.Close()

}

func TestDeletar(t *testing.T) {

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	conexao := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USUARIO_TEST"),
		os.Getenv("DB_SENHA_TEST"),
		os.Getenv("DB_HOST_TEST"),
		os.Getenv("DB_PORTA_TEST"),
		os.Getenv("DB_NOME_TEST"),
	)

	db, err := sql.Open("postgres", conexao)

	migration.Run(context.Background(), "./banco/fixtures", conexao, "up")

	if err != nil {
		t.Fatalf("Não conseguiu criar um DATABASE %v", err)
	}

	cliente := model.Cliente{
		Uuid:     util.RandomString(10),
		Nome:     util.RandomString(10),
		Endereco: util.RandomString(20),
	}

	repositorio := NovoRepositorioDeCliente(db)

	clienteID, err := repositorio.Criar(cliente)

	if err != nil {
		t.Fatalf("Não conseguiu criar um cliente | erro: %v", err)
	}

	err = repositorio.Deletar(clienteID)

	if err != nil {
		t.Fatalf("Não conseguiu deçetar o cliente | erro: %v", err)
	}

	defer db.Close()

}
