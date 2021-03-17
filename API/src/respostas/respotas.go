package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Faz a conversão do objeto para um JSON e retorna uma reposta para requisição.
// Também envia o status code da operação.
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// Trata um erro e o envia para ser transforamdo em um JSON e envia também o status code.
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	},
	)
}
