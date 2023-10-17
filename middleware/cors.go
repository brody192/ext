package middleware

import (
	"fmt"
	"net/http"

	"github.com/brody192/ext/variables"
)

// sets all cors headers to accept anything
func CorsAny(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(variables.HeaderAccessControlAllowCredentials, variables.HeaderTrue)
		w.Header().Set(variables.HeaderAccessControlAllowOrigin, "*")
		w.Header().Set(variables.HeaderAccessControlAllowMethods,
			fmt.Sprintf(
				"%s, %s, %s, %s, %s, %s, %s, %s, %s",
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodConnect,
				http.MethodDelete,
				http.MethodHead,
				http.MethodPatch,
				http.MethodOptions,
				http.MethodTrace,
			))
		w.Header().Set(variables.HeaderAccessControlAllowHeaders, variables.HeaderWildcard)
		next.ServeHTTP(w, r)
	})
}
