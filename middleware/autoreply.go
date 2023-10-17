package middleware

import (
	"net/http"
	"path"
)

// auto reply to paths listed in the paths slice with given code
//
// given paths are pre-compiled to be clean, incoming paths are cleaned and checked against the cleaned paths
func AutoReply(paths []string, code int) func(http.Handler) http.Handler {
	var pathCodes = make(map[string]int, len(paths))

	for _, p := range paths {
		pathCodes[p] = code
	}

	return AutoReplyMap(pathCodes)
}

// auto reply to paths with the paths given status code
//
// given paths are pre-compiled to be clean, incoming paths are cleaned and checked against the cleaned paths
func AutoReplyMap(pathCodes map[string]int) func(http.Handler) http.Handler {
	var pathCodesClean = make(map[string]int, len(pathCodes))

	for p, c := range pathCodes {
		var cleanPath = path.Clean(p)

		pathCodesClean[cleanPath] = c
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var urlPathClean = path.Clean(r.URL.Path)

			if code, ok := pathCodesClean[urlPathClean]; ok {
				w.WriteHeader(code)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
