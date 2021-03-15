package model

import (
	"errors"
	"strings"
)

type Cliente struct {
	Uuid          string
	Nome          string
	Endereco      string
	Cadastrado_em string
	Atualizado_em string
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
