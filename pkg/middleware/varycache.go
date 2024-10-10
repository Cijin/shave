package middleware

import "net/http"

/*
 * This lets the browser know that a request with the same URL without
 * the `HX-Request` header is not the same
 */
func VaryCache(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "HX-Request")

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
