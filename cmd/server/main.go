package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"shave/internal/version"
	"shave/pkg/handlers"

	"shave/pkg/middleware"

	m "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/tursodatabase/go-libsql"

	"github.com/go-chi/httprate"
)

const (
	defaultPort  = "8080"
	syncInterval = time.Minute
	tmpDir       = "tmp"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info("Env file not found, using enviornment values")
	}

	// Db Setup -------------------------
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		slog.Error("DB_NAME cannot be empty", "ENV_ERROR", "missing db name")
	}

	// in production, fly will mount a volume in
	// the current directory: /app/tmp, check fly.toml
	// /app is the workdir in the final docker image
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("Unable to get working directory", "OS_ERROR", err)
	}
	tempDirPath := filepath.Join(wd, tmpDir)

	dir, err := os.MkdirTemp(tempDirPath, "libsql-*")
	if err != nil {
		slog.Error("Error creating temporary directory", "OS_ERROR", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)
	primaryURL := os.Getenv("DB_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryURL,
		libsql.WithAuthToken(authToken),
		libsql.WithSyncInterval(syncInterval),
	)
	if err != nil {
		slog.Error("Error creating connector", "LIBSQL_ERROR", err)
		os.Exit(1)
	}
	defer connector.Close()

	db := sql.OpenDB(connector)
	defer db.Close()

	// HttpHandler -------------------------
	h, err := handlers.NewHttpHandler(db)
	if err != nil {
		slog.Error("Error setting up handler.", "Error", err)
		return
	}

	// Port -------------------------
	port := os.Getenv("PORT")
	if port == "" {
		slog.Info("Port not set, using defaults")
		port = defaultPort
	}

	// Router -------------------------
	router := chi.NewRouter()
	router.Use(m.Logger)
	router.Use(httprate.LimitByIP(100, 1*time.Minute))
	router.Use(middleware.Cors())
	router.Use(middleware.VaryCache)
	//	router.Use(m.Recoverer)

	// Create a route along / that will serve contents from
	// the public folder
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	fileServer(router, "/", filesDir)

	registerRoutes(router, h)

	// Server -------------------------
	server := http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           router,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 20,
		IdleTimeout:       time.Minute * 2,
	}

	// Get migration info
	migrationInfo, err := getLatestMigrationInfo(db)
	if err != nil {
		slog.Error("Unable to retrieve migration info", "Error", err)
	}

	slog.Info("Migration", "Version", migrationInfo.VersionId, "Successful", migrationInfo.IsApplied, "Timestamp", migrationInfo.Tstamp)
	slog.Info("Server starting", "Version", version.Version, "Port", port)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed, shutting down", "error", err)

		if err := server.Shutdown(context.Background()); err != nil {
			slog.Error("Server shutdown failed", "error", err)
		}
	}
}
