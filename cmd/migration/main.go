// This is custom goose binary with postgres support only.
package main

import (
	"context"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

const command = "up"

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set, unable to run migrations")
		return
	}

	migrationDir := os.Getenv("MIGRATION_DIR")
	if migrationDir == "" {
		log.Fatal("MIGRATION_DIR not set, unable to run migrations")
		return
	}

	db, err := goose.OpenDBWithDriver("turso", dbURL)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.RunContext(context.Background(), command, db, migrationDir); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
