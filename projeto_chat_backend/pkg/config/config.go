package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "your_host"
	port     = 5432
	user     = "your_user"
	password = "your_password"
	dbname   = "your_dbname"
)

// GetDBConnection retorna uma nova conexão com o banco de dados
func GetDBConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}
	return db, nil
}
