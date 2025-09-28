package middleware

import (
	"context"
	"net/http"
)

func IsHtmxRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Header.Get("HX-Request") == "true" {
			// SA1029: using string key for context is acceptable here
			//nolint:staticcheck
			ctx = context.WithValue(ctx, "HX-Request", true)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
