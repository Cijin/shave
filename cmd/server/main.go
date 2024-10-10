package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"paper-chase/internal/version"
	"paper-chase/pkg/handlers"
	"paper-chase/pkg/middleware"
	"path/filepath"
	"time"

	m "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/go-chi/httprate"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info("Env file not found, using enviornment values")
	}

	// Db Setup -------------------------
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		slog.Error("Failed to open db.", "Error", err)
	}
	defer db.Close()

	// Port -------------------------
	port := os.Getenv("PORT")
	if port == "" {
		slog.Info("Port not set, using defaults")
		port = defaultPort
	}

	// Register handler
	h, err := handlers.NewHttpHandler(db)
	if err != nil {
		slog.Error("Error setting up handler.", "Error", err)
		return
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

		for _, cancel := range h.Cancels {
			cancel()
		}

		if err := server.Shutdown(context.Background()); err != nil {
			slog.Error("Server shutdown failed", "error", err)
		}
	}
}
