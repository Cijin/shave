package main

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

/*
This helper does what it says. Easy to use. Ensure all public
assets are in the `/public` folder. You do not need to include
the `/public` string when accessing public assets.
*/
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("fileserver does not permit any parameters")
	}

	// add trailing '/' if not existing
	if path != "/" && path[len(path)-1] != '/' {
		// letting the caller know that resource has moved from path to path/
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		// ex: public/*
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		// removes public from /public/*
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))

		w.Header().Set("Cache-Control", "max-age=31536000, public, must-revalidate")
		fs.ServeHTTP(w, r)
	})
}
