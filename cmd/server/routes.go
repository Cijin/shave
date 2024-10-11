package main

import (
	"shave/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func registerRoutes(r chi.Router, h *handlers.HttpHandler) {
	r.Get("/", h.HomePage)
}
