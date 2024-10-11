package handlers

import (
	"net/http"

	"shave/views/home"
)

func (h *HttpHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	renderComponent(w, r, home.Index())
}
