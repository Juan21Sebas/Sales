package main

import (
	"log"

	"github.com/clubhub/sales/cmd/config/db"
	server "github.com/clubhub/sales/internal/infrastructure/api/candidates"
)

func main() {
	dbConfig := db.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Database: "survey",
		User:     "postgres",
		Password: "postgres",
		SSLMode:  "disable",
	}

	database, err := db.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	server.RunServer(database)
}
