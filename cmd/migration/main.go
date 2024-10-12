package main

import (
	"database/sql"
	"embed"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	dbURL := os.Getenv("SQL_URL")
	if dbURL == "" {
		log.Fatal("No SQL_URL set in env")
	}

	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		log.Fatal("error opening database: ", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal("Unable to set goose dialect:", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}
