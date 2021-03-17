package model

import (
	"errors"
	"strings"
)

type Cliente struct {
	Uuid          string `json:"Uuid,omitempty"`
	Nome          string `json:"Nome,omitempty"`
	Endereco      string `json:"Endereco,omitempty"`
	Cadastrado_em string `json:"Cadastrado_em,omitempty"`
	Atualizado_em string `json:"Atualizado_em,omitempty"`
}

// Faz um preparo para que seja realizada a validação e a formatação dos valores que entraram
// no caso ele recebe um objeto do tipo Cliente e os atributos que estão dentro dele serão validados e formatados.
func (cliente *Cliente) Preparar() error {
	if erro := cliente.validar(); erro != nil {
		return erro
	}
	cliente.formatar()
	return nil
}

// Executa uma validação para informar que o nome é obrigatorio.
func (cliente *Cliente) validar() error {
	if cliente.Nome == "" {
		return errors.New("O nome é obrigatorio")
	}
	return nil
}

// Faz a formatação para que seja removido os espaços do inicio e do fim do nome e endereço
func (cliente *Cliente) formatar() {
	cliente.Nome = strings.TrimSpace(cliente.Nome)
	cliente.Endereco = strings.TrimSpace(cliente.Endereco)
}
