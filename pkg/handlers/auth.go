package handlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"time"

	"shave/internal/database"
	"shave/pkg/data"
	"shave/views/home"
	"shave/views/unauthorized"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

const tokenExpiryThreshold = time.Minute * (-5)

type authedHandler func(w http.ResponseWriter, r *http.Request, sessionUser data.SessionUser)

func (h *HttpHandler) CheckAuthoziation(w http.ResponseWriter, r *http.Request) (data.SessionUser, error) {
	var user data.SessionUser

	session, err := h.store.GetSession(r)
	if err != nil {
		return user, err
	}

	user, err = h.store.GetSessionUser(r)
	if err != nil {
		return user, err
	}

	savedSession, err := h.dbQueries.GetSession(r.Context(), user.Email)
	if err != nil {
		return data.SessionUser{}, err
	}

	if savedSession.AccessToken != session.AccessToken || user.UserId.String() != savedSession.UserID {
		return data.SessionUser{}, err
	}

	if session.Expiry.Before(time.Now().Add(tokenExpiryThreshold)) {
		return h.refreshToken(w, r, user, session)
	}

	return user, nil
}

func (h *HttpHandler) Authorize(next authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionUser, err := h.CheckAuthoziation(w, r)
		if err != nil {
			slog.Error("Session or User data is malformed or non existent", "AUTHORIZE_ERROR", err)

			if r.URL.Path == "/" {
				h.HomePage(w, r, data.SessionUser{})
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
			return
		}

		if r.URL.Path == "/" {
			renderComponent(w, r, home.SessionedHome(sessionUser))
			return
		}

		next(w, r, sessionUser)
	})
}

func (h *HttpHandler) refreshToken(w http.ResponseWriter, r *http.Request, u data.SessionUser, s data.Session) (data.SessionUser, error) {
	var user data.SessionUser

	savedSession, err := h.dbQueries.GetSession(r.Context(), u.Email)
	if err != nil {
		return user, err
	}

	if savedSession.UserID != u.UserId.String() {
		return user, fmt.Errorf("user ID from database=%s, does not match session=%s", savedSession.UserID, u.UserId.String())
	}

	if savedSession.Provider != s.Provider {
		return user, fmt.Errorf("provider from database=%s does not match session=%s", savedSession.Provider, s.Provider)
	}

	token, err := h.authenticator.RefreshToken(r.Context(), s.Provider, savedSession.RefreshToken)
	if err != nil {
		return user, err
	}

	updateSessionParams := database.UpdateSessionParams{
		RefreshToken: token.RefreshToken,
		AccessToken:  token.AccessToken,
		ID:           savedSession.ID,
	}

	err = h.dbQueries.UpdateSession(r.Context(), updateSessionParams)
	if err != nil {
		return user, err
	}

	s.AccessToken = token.AccessToken
	s.Expiry = token.Expiry

	err = h.store.SaveSession(w, r, s)
	if err != nil {
		return user, err
	}

	return u, nil
}

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
	if !slices.Contains(data.SupportedProviders, provider) {
		slog.Error("Unrecognized provider", "ERROR", fmt.Errorf("unrecognized provider '%s'", provider))

		renderComponent(w, r, unauthorized.Index("Login failed"))
		return
	}

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
				ID:            getUUID().String(),
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
		} else {
			slog.Error("Unable to get user", "DB_ERROR", err)

			InternalError(w, r)
			return
		}
	}
	sessionUser.UserId, _ = uuid.Parse(user.ID)

	createSessionParams := database.CreateSessionParams{
		ID:           getUUID().String(),
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
