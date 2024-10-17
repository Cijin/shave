package handlers

import (
	"net/http"

	"shave/pkg/data"
	"shave/views/home"
)

func (h *HttpHandler) HomePage(w http.ResponseWriter, r *http.Request, _ data.SessionUser) {
	renderComponent(w, r, home.Index())
}
