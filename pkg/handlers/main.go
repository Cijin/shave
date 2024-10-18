package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"

	"shave/internal/database"
	"shave/pkg/authenticator"
	"shave/pkg/authenticator/providers/google"
	"shave/pkg/store"
	"shave/views/internalError"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

type HttpHandler struct {
	db            *sql.DB
	dbQueries     *database.Queries
	store         *store.Store
	schemaDecoder *schema.Decoder
	authenticator *authenticator.Authenticator
}

func NewHttpHandler(db *sql.DB) (*HttpHandler, error) {
	store, err := store.New()
	if err != nil {
		return nil, err
	}

	googleProvider, err := google.New()
	if err != nil {
		return nil, err
	}

	authenticator := authenticator.New(true, googleProvider)

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	dbQueries := database.New(db)

	return &HttpHandler{
		db:            db,
		dbQueries:     dbQueries,
		store:         store,
		schemaDecoder: decoder,
		authenticator: authenticator,
	}, nil
}

// Helpers
func getUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		slog.Error("Unable to create uuid v7, defaulting to uuid v4", "Error", err)
		id = uuid.New()
	}

	return id
}

// helpers
// render component should be used when redering items that are part of a page
// easy way to differentiate, is they usually don't render nav within
func renderComponent(w http.ResponseWriter, r *http.Request, component templ.Component) {
	err := component.Render(r.Context(), w)
	if err != nil {
		slog.Error("Unable to render component: ", "Error", err)
	}
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("HX-Push-Url", "500")
	w.Header().Add("HX-Retarget", "body")
	renderComponent(w, r, internalError.Index())
}
