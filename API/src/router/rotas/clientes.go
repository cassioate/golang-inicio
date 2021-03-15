package rotas

import (
	"modulo/src/controllers"
	"net/http"
)

var rotasClientes = []Rota{
	{
		Uri:                "/cliente",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodosClientes,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/cliente/{UUID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarClientePorId,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/cliente",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCliente,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/cliente/{UUID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCliente,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/cliente/{UUID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCliente,
		RequerAutenticacao: false,
	},
}
