package middleware

import (
	"net/http"
	"path/filepath"
	"strings"
)

// adds a trailing slash to the request path, uses http.StatusMovedPermanently
//
// must come before CleanPath, or PrefixRemove
func AddTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/") && filepath.Ext(r.URL.Path) == "" {
			var path = r.URL.Path
			var qs = r.URL.RawQuery
			path += "/"
			var uri = path
			if qs != "" {
				uri += "?" + qs
			}

			http.Redirect(w, r, sanitizeURI(uri), http.StatusMovedPermanently)

			r.RequestURI = path
			r.URL.Path = path
			return
		}

		next.ServeHTTP(w, r)
	})
}
