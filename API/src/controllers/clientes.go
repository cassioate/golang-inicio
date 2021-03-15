package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo/banco"
	"modulo/src/model"
	"modulo/src/repositorios"
	"modulo/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarTodosClientes(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	clientes, erro := repositorio.BuscarTodosClientes()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	respostas.JSON(w, http.StatusOK, clientes)
}

func CriarCliente(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cliente model.Cliente

	if erro := json.Unmarshal(corpoRequest, &cliente); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = cliente.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	clienteId, erro := repositorio.Criar(cliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db.Close()
	respostas.JSON(w, http.StatusCreated, clienteId)
}

func BuscarClientePorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID, erro := strconv.ParseUint(parametros["UUID"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	cliente, erro := repositorio.BuscarClientePorId(clienteID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	respostas.JSON(w, http.StatusOK, cliente)
}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID, erro := strconv.ParseUint(parametros["UUID"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cliente model.Cliente
	if erro = json.Unmarshal(corpoRequest, &cliente); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = cliente.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	erro = repositorio.Atualizar(clienteID, cliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db.Close()
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID, erro := strconv.ParseUint(parametros["UUID"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeCliente(db)
	erro = repositorio.Deletar(clienteID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db.Close()
	respostas.JSON(w, http.StatusNoContent, nil)
}
