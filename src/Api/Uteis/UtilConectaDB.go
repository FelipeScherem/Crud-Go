package util

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectDB cria e retorna uma conexão com o banco de dados PostgreSQL usando o GORM
func ConectaDB() *gorm.DB {
	// Busca as infos do banco no .ENV
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Constroi a string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Abre a conexão
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Retorna a conexão
	return database
}

// FechaDB fecha a conexão com o banco de dados
func FechaDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Erro ao obter o objeto sql.DB: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Erro ao fechar a conexão com o banco de dados: %v", err)
	}
}
