package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"shave/views/unauthorized"

	"github.com/go-chi/chi"
)

func (h *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	state, err := h.store.SaveState(w, r)
	if err != nil {
		slog.Error("Unable to save state", "SESSION_ERROR", err)

		InternalError(w, r)
		return
	}

	url, err := h.authenticator.AuthCodeURL(state, r)
	if err != nil {
		slog.Error("Check if provider is registered", "AUTHENTICATOR_ERROR", err)

		InternalError(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (h *HttpHandler) AuthCallback(w http.ResponseWriter, r *http.Request) {
	stateParam := chi.URLParam(r, "state")

	state, err := h.store.GetState(r)
	if err != nil {
		slog.Error("Unable to get state for session", "SESSION_ERROR", err)

		InternalError(w, r)
		return
	}

	if stateParam != state {
		renderComponent(w, r, unauthorized.Index("Unable to login, request was tampered"))
		return
	}

	token, idToken, err := h.authenticator.Authenticate(r)
	if err != nil {
		renderComponent(w, r, unauthorized.Index("Login failed"))
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		InternalError(w, r)
		return
	}

	fmt.Println(token, profile)
}
