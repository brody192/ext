package middleware

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
)

// cleans path of all double slashes, uses path.Clean internally
func CleanPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rctx = chi.RouteContext(r.Context())

		r.URL.Path = path.Clean(r.URL.Path)
		r.URL.Path = strings.Replace(r.URL.Path, "https:/", "https://", 1)

		if r.URL.RawPath != "" {
			r.URL.RawPath = path.Clean(r.URL.RawPath)
			r.URL.RawPath = strings.Replace(r.URL.RawPath, "https:/", "https://", 1)
		}

		rctx.RoutePath = r.URL.Path

		next.ServeHTTP(w, r)
	})
}
