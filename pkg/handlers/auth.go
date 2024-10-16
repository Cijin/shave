package handlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"shave/internal/database"
	"shave/pkg/data"
	"shave/views/unauthorized"

	"github.com/go-chi/chi"
	"golang.org/x/oauth2"
)

func (h *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	sessionVerifier, err := h.store.SaveSessionVerfier(w, r)
	if err != nil {
		slog.Error("Unable to save state", "SESSION_ERROR", err)

		InternalError(w, r)
		return
	}

	url, err := h.authenticator.AuthCodeURL(sessionVerifier.State.String(), r, oauth2.S256ChallengeOption(sessionVerifier.Verifier))
	if err != nil {
		slog.Error("Check if provider is registered", "AUTHENTICATOR_ERROR", err)

		InternalError(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *HttpHandler) AuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	sessionVerifier, err := h.store.GetSessionVerfier(r)
	if err != nil {
		slog.Error("Unable to get state for session", "SESSION_ERROR", err)

		InternalError(w, r)
		return
	}

	stateParam := r.URL.Query().Get("state")
	if stateParam != sessionVerifier.State.String() {
		renderComponent(w, r, unauthorized.Index("Unable to login, request was tampered"))
		return
	}

	token, sessionUser, err := h.authenticator.Authenticate(r, oauth2.VerifierOption(sessionVerifier.Verifier))
	if err != nil {
		slog.Error("Unable to login", "AUTHENTICATOR_ERROR", err)

		renderComponent(w, r, unauthorized.Index("Login failed"))
		return
	}

	var user database.User
	user, err = h.dbQueries.GetUser(r.Context(), sessionUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			createUserParams := database.CreateUserParams{
				ID:            getUUID(),
				Email:         sessionUser.Email,
				Sub:           sessionUser.Sub,
				Name:          sessionUser.Name,
				EmailVerified: sessionUser.EmailVerified,
				CreatedAt:     time.Now().UTC(),
				UpdatedAt:     time.Now().UTC(),
			}

			user, err = h.dbQueries.CreateUser(r.Context(), createUserParams)
			if err != nil {
				slog.Error("Unable to create user", "DB_ERROR", err)

				InternalError(w, r)
				return
			}
		}
	} else {
		slog.Error("Unable to get user", "DB_ERROR", err)

		InternalError(w, r)
		return
	}

	createSessionParams := database.CreateSessionParams{
		ID:           getUUID(),
		UserID:       user.ID,
		Email:        sessionUser.Email,
		Provider:     provider,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	_, err = h.dbQueries.CreateSession(r.Context(), createSessionParams)
	if err != nil {
		slog.Error("Unable to save session info", "DB_ERROR", err)

		InternalError(w, r)
		return
	}

	session := data.Session{
		AccessToken: token.AccessToken,
		Expiry:      token.Expiry,
		Provider:    provider,
	}
	err = h.store.SaveSession(w, r, session)
	if err != nil {
		slog.Error("Unable to save session info", "SESSION_ERROR", err)

		InternalError(w, r)
		return
	}

	err = h.store.SaveSessionUser(w, r, sessionUser)
	if err != nil {
		slog.Error("Unable to save session user info", "SESSION_ERROR", err)

		InternalError(w, r)
		return
	}

	fmt.Println(token, sessionUser)
}
