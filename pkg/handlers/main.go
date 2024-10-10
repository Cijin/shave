package handlers

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"paper-chase/internal/database"
	"paper-chase/pkg/store"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

type HttpHandler struct {
	db            *sql.DB
	dbQueries     *database.Queries
	store         *store.Store
	schemaDecoder *schema.Decoder
	awsService    *aws.AwsService
	Cancels       []context.CancelFunc
}

func NewHttpHandler(db *sql.DB) (*HttpHandler, error) {
	// AWS Service -------------------------
	awsService, err := aws.New()
	if err != nil {
		return nil, err
	}

	var cancels []context.CancelFunc
	ctx, cancel := context.WithCancel(context.Background())
	go awsService.RotateJWKS(ctx)

	cancels = append(cancels, cancel)

	store, err := store.New()
	if err != nil {
		return nil, err
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	dbQueries := database.New(db)

	return &HttpHandler{
		db:            db,
		dbQueries:     dbQueries,
		store:         store,
		schemaDecoder: decoder,
		awsService:    awsService,
		Cancels:       cancels,
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

func getFormatedDate(t time.Time) string {
	// BKK is UTC + 7
	offset := 7 * time.Hour
	return t.Add(offset).Format("2 Jan 2006")
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
