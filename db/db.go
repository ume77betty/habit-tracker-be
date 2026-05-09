package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open database", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal("failed to connect database", err)
	}

	return database
}
