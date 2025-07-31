package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

var DB *sql.DB

// Função init executa ao importar o pacote
func init() {
	var err error

	// Dados de conexão (ajuste conforme seu ambiente)
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	// Abre a conexão com o banco
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %v", err)
	}

	// Testa a conexão
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao pingar o banco: %v", err)
	}

	log.Println("Conexão com banco de dados estabelecida.")
}
