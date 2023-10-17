package middleware

import (
	"net/http"
	"strings"
)

// disallow a fragment specified in list from appearing in path
//
// returns with the status code specified by code
func DisallowPaths(list []string, code int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, txt := range list {
				if strings.Contains(r.URL.Path, txt) {
					http.Error(w, http.StatusText(code), code)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

// disallow a requests with headers specified in headers
//
// returns with the status code specified by code
func DisallowHeaders(headers []string, code int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, h := range headers {
				if r.Header.Get(h) != "" {
					http.Error(w, http.StatusText(code), code)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
