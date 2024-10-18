package handlers

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

var db *sql.DB

func TestMain(m *testing.M) {
	dir, err := testTempDir()
	if err != nil {
		slog.Error("Error creating temporary directory", "TEST_OS_ERROR", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, "test.db")
	db, err = sql.Open("libsql", dbPath)
	if err != nil {
		slog.Error("Failed to open db:", "TEST_DB_ERROR", err)
		os.Exit(1)
	}
	defer db.Close()

	err = runMigrations(db)
	if err != nil {
		slog.Error("Migrations failed", "TEST_MIGRATION_ERROR", err)
		os.Exit(1)
	}

	m.Run()
}

func testTempDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	tempDirPath := filepath.Join(wd, "tmp")
	dir, err := os.MkdirTemp(tempDirPath, "test-db")
	if err != nil {
		return "", err
	}

	return dir, nil
}

func runMigrations(db *sql.DB) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	migrationsDir := filepath.Join(wd, "cmd", "migration", "migrations")
	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	return nil
}
