package extutil

import (
	"io/fs"
	"net/http"
	"os"
	"strings"
)

// panics if fs.sub fails
func MustSubFS(fsys fs.FS, dir string) fs.FS {
	fsub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}

	return fsub
}

// use environment variable PORT or specified port
//
// returned value is prefixed with `:`
func EnvPortOr(port string) string {
	if envPort, envExists := os.LookupEnv("PORT"); envExists {
		return ":" + envPort
	}
	return ":" + strings.TrimPrefix(port, ":")
}

// checks if provided method is valid as defined by the functions switch cases
func IsValidMethod(method string) bool {
	switch method {
	case
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodConnect,
		http.MethodDelete,
		http.MethodHead,
		http.MethodPatch,
		http.MethodOptions,
		http.MethodTrace:
		return true
	}
	return false
}

// return the given query parameter with all leading and trailing white space removed, as defined by Unicode.
func TrimmedQParam(r *http.Request, q string) string {
	var qp = r.URL.Query().Get(q)

	if qp != "" {
		return strings.TrimSpace(qp)
	}

	return qp
}
