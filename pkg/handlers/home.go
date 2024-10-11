package handlers

import "net/http"

func HomePage(w http.ResponseWriter, r *http.Request) {
	renderComponent(w, r, home.Index())
}
