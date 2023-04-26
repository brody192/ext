package extutil

import (
	"io/fs"
	"net/http"
	"os"
	"sort"
	"strings"
)

// binary search if v is in s
//
// NOTE: s must be sorted
func ContainsString(s []string, v string) bool {
	var i = sort.SearchStrings(s, v)
	return i < len(s) && s[i] == v
}

// binary search if v is in s
//
// NOTE: s must be sorted
func ContainsInt(s []int, v int) bool {
	var i = sort.SearchInts(s, v)
	return i < len(s) && s[i] == v
}

// binary search if v is in s
//
// NOTE: s must be sorted
func ContainsFloat64(s []float64, v float64) bool {
	var i = sort.SearchFloat64s(s, v)
	return i < len(s) && s[i] == v
}

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
//
// returns empty string if the query parameter does not exist
func TrimmedQParam(r *http.Request, q string) string {
	var qp = r.URL.Query().Get(q)

	if qp != "" {
		return strings.TrimSpace(qp)
	}

	return qp
}
