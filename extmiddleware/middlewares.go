package extmiddleware

import (
	"net/http"
	"os"
	"strings"
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

// use environment variable PORT or specified port
// returned value is prefixed with `:`
func EnvPortOr(port string) string {
	if envPort, envExists := os.LookupEnv("PORT"); envExists {
		return ":" + envPort
	}
	return ":" + strings.TrimPrefix(port, ":")
}

// removes specified prefix from path
func PrefixRemove(prefix string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = strings.Replace(r.URL.Path, "/"+prefix, "/", 1)
			r.URL.Path = strings.Replace(r.URL.Path, "//", "/", 1)
			next.ServeHTTP(w, r)
		})
	}
}

// disallow a requests with headers specified in headers
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
