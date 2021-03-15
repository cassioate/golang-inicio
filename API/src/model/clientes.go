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

func (cliente *Cliente) Preparar() error {
	if erro := cliente.validar(); erro != nil {
		return erro
	}
	cliente.formatar()
	return nil
}

func (cliente *Cliente) validar() error {
	if cliente.Nome == "" {
		return errors.New("O nome é obrigatorio")
	}
	return nil
}

func (cliente *Cliente) formatar() {
	cliente.Nome = strings.TrimSpace(cliente.Nome)
	cliente.Endereco = strings.TrimSpace(cliente.Endereco)
}
