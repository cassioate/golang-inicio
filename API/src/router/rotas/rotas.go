package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa as rotas da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Congirurar recebe uma rota sem nada dentro (r) passa por um for e devolve com as rotas preenchidas com os valores que estavam no slise rota.
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasClientes

	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
