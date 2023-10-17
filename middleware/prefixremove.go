package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// removes specified prefix from path
func PrefixRemove(prefix string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var rctx = chi.RouteContext(r.Context())

			r.URL.Path = strings.Replace(r.URL.Path, "/"+prefix, "/", 1)
			r.URL.Path = strings.Replace(r.URL.Path, "//", "/", 1)

			if r.URL.RawPath != "" {
				r.URL.RawPath = strings.Replace(r.URL.RawPath, "/"+prefix, "/", 1)
				r.URL.RawPath = strings.Replace(r.URL.RawPath, "//", "/", 1)
			}

			rctx.RoutePath = r.URL.Path

			next.ServeHTTP(w, r)
		})
	}
}
