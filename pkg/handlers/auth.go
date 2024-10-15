package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"shave/views/unauthorized"

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

	// save session
	// gzip??
	// save session to db
	// save provider??

	// save session user

	// redirect to logged in page
	fmt.Println(token, sessionUser)
}
