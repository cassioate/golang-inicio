package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo/banco"
	"modulo/src/model"
	"modulo/src/repositorios"
	"modulo/src/respostas"
	"net/http"

	"github.com/gorilla/mux"
)

// Recebe o GET(ALL) e Tem a função de abrir a conexão com o banco de dados, e chamar o repositorio para que seja enviado essa conexão
// de forma que o repositorio possa fazer suas tratativas, após o retorno do repositorio será realizado o fechamento da conexão com o banco
// e então será enviada a resposta em formato Json, com o status http.
func BuscarTodosClientes(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	seErro(w, http.StatusUnprocessableEntity, erro)

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	clientes, erro := repositorio.BuscarTodosClientes()
	seErro(w, http.StatusInternalServerError, erro)

	defer db.Close()

	respostas.JSON(w, http.StatusOK, clientes)
}

// Recebe o POST e tem a função de abrir a conexão com o banco de dados, e chamar o repositorio para que seja enviado essa conexão
// de forma que o repositorio possa fazer suas tratativas, após o retorno do repositorio será realizado o fechamento da conexão com o banco
// e então será enviada a resposta em formato Json, com o status http.
func CriarCliente(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	seErro(w, http.StatusUnprocessableEntity, erro)

	var cliente model.Cliente
	erro = json.Unmarshal(corpoRequest, &cliente)
	seErro(w, http.StatusBadRequest, erro)

	erro = cliente.Preparar()
	seErro(w, http.StatusBadRequest, erro)

	db, erro := banco.Conectar()
	seErro(w, http.StatusInternalServerError, erro)

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	clienteId, erro := repositorio.Criar(cliente)
	seErro(w, http.StatusInternalServerError, erro)

	db.Close()
	respostas.JSON(w, http.StatusCreated, clienteId)
}

// Recebe o GET(ID) e tem a função de abrir a conexão com o banco de dados, e chamar o repositorio para que seja enviado essa conexão
// de forma que o repositorio possa fazer suas tratativas, após o retorno do repositorio será realizado o fechamento da conexão com o banco
// e então será enviada a resposta em formato Json, com o status http.
func BuscarClientePorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID := string(parametros["UUID"])

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	cliente, erro := repositorio.BuscarClientePorId(clienteID)
	seErro(w, http.StatusInternalServerError, erro)

	defer db.Close()
	respostas.JSON(w, http.StatusOK, cliente)
}

// Recebe o PUT e tem a função de abrir a conexão com o banco de dados, e chamar o repositorio para que seja enviado essa conexão
// de forma que o repositorio possa fazer suas tratativas, após o retorno do repositorio será realizado o fechamento da conexão com o banco
// e então será enviada a resposta em formato Json, com o status http.
func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID := string(parametros["UUID"])

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	seErro(w, http.StatusUnprocessableEntity, erro)

	var cliente model.Cliente
	erro = json.Unmarshal(corpoRequest, &cliente)
	seErro(w, http.StatusBadRequest, erro)

	erro = cliente.Preparar()
	seErro(w, http.StatusBadRequest, erro)

	db, erro := banco.Conectar()
	seErro(w, http.StatusInternalServerError, erro)

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	erro = repositorio.Atualizar(clienteID, cliente)
	seErro(w, http.StatusInternalServerError, erro)

	db.Close()
	respostas.JSON(w, http.StatusNoContent, nil)
}

// Recebe o DELETE e tem a função de abrir a conexão com o banco de dados, e chamar o repositorio para que seja enviado essa conexão
// de forma que o repositorio possa fazer suas tratativas, após o retorno do repositorio será realizado o fechamento da conexão com o banco
// e então será enviada a resposta em formato Json, com o status http.
func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID := string(parametros["UUID"])

	db, erro := banco.Conectar()
	seErro(w, http.StatusInternalServerError, erro)

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	erro = repositorio.Deletar(clienteID)
	seErro(w, http.StatusInternalServerError, erro)

	db.Close()
	respostas.JSON(w, http.StatusNoContent, nil)
}

// Tem a função de encurtar o codigo em caso de erro no controller
func seErro(w http.ResponseWriter, statusCode int, erro error) {
	if erro != nil {
		respostas.Erro(w, statusCode, erro)
		return
	}
}
