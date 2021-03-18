package util

import (
	"math/rand"
	"strings"
	"time"
)

const letrasENumeros = "abcdefghijklmnopqrstuvxyz0123456789"

// define a seed para realizar o random em sequencia
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(letrasENumeros)

	for i := 0; i < n; i++ {
		c := letrasENumeros[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
