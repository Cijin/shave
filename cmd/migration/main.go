package main

import (
	"database/sql"
	"embed"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

const command = "up"

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	sqlURL := os.Getenv("SQL_URL")
	if sqlURL == "" {
		log.Fatal("No SQL_URL set in env")
	}

	db, err := sql.Open("libsql", sqlURL)
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
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}
