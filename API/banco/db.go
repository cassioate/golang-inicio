package banco

import (
	"database/sql"
	"modulo/src/config"

	_ "github.com/lib/pq"
)

func Conectar() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConexaoBanco)
	if err != nil {
		panic(err.Error())
	}

	return db, err
}
