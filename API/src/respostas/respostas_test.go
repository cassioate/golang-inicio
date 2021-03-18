package respostas

import (
	"database/sql"
	"modulo/src/model"
	"modulo/src/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	// respostas.JSON(w, http.StatusOK, clientes)

	cliente := model.Cliente{
		Uuid:          util.RandomString(10),
		Nome:          util.RandomString(10),
		Endereco:      util.RandomString(20),
		Cadastrado_em: "18/03/2020 10:30:30",
		Atualizado_em: "18/03/2020 10:30:30",
	}

	w := httptest.NewRecorder()

	JSON(w, http.StatusOK, cliente)

	if w.Result().Status != "200 OK" {
		t.Fatal("Erro no httpStatus")
	}

}

func TestErro(t *testing.T) {
	// respostas.JSON(w, http.StatusOK, clientes)
	_, err := sql.Open("postgres", "TesteErro")
	w := httptest.NewRecorder()

	Erro(w, http.StatusInternalServerError, err)

}
