package model

import (
	"modulo/src/util"
	"testing"
)

func TestPreparar(t *testing.T) {
	cliente := Cliente{
		Uuid:          util.RandomString(10),
		Nome:          util.RandomString(10),
		Endereco:      util.RandomString(20),
		Cadastrado_em: "18/03/2020 10:30:30",
		Atualizado_em: "18/03/2020 10:30:30",
	}

	cliente.Preparar()
	if cliente.Nome == "" {
		t.Fatalf("Reprovado no teste de validação do nome")
	}

	if cliente.Nome[0:1] == "" {
		t.Fatalf("Reprovado no teste de formatação e remoção dos espaços")
	}

	cliente2 := Cliente{
		Uuid:          util.RandomString(10),
		Endereco:      util.RandomString(20),
		Cadastrado_em: "18/03/2020 10:30:30",
		Atualizado_em: "18/03/2020 10:30:30",
	}

	err := cliente2.Preparar()

	if err == nil {
		t.Fatalf("Reprovado no teste de obrigatoriedade do nome")
	}
}
