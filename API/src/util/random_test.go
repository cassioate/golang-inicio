package util

import "testing"

func TestRandom(t *testing.T) {
	sb := RandomString(8)

	if sb == "" {
		t.Fatal("Não conseguiu gerar um numero aleatorio no teste")
	}
}
