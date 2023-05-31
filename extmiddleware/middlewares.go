package extmiddleware

import (
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

// sets all cors headers to accept anything
func CorsAny(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}

// disallow a fragment specified in list from appearing in path
//
// returns with the status code specified by code
func DisallowInPath(list []string, code int) func(http.Handler) http.Handler {
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

// disallow a requests with headers specified in headers
//
// returns with the status code specified by code
func BlockHeaders(headers []string, code int) func(http.Handler) http.Handler {
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

// adds a trailing slash to the request path, uses http.StatusMovedPermanently
//
// must come before CleanPath, or PrefixRemove
func AddTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") == false && filepath.Ext(r.URL.Path) == "" {
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
