package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

// InitializeDatabase inicializa a conexão com o banco de dados
func InitializeDatabase() {
	// Obtém a URL do banco de dados das variáveis de ambiente
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL não está definido nas variáveis de ambiente")
	}

	// Cria uma conexão com o pool do PostgreSQL
	var err error
	dbPool, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados PostgreSQL estabelecida!")
}

// GetDatabase retorna a conexão compartilhada do banco de dados
func GetDatabase() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("A conexão com o banco de dados ainda não foi inicializada")
	}
	return dbPool
}

// CloseDatabase fecha a conexão com o banco de dados
func CloseDatabase() {
	if dbPool != nil {
		dbPool.Close()
		log.Println("Conexão com o banco de dados PostgreSQL encerrada.")
	}
}
