package extutil

import (
	"io/fs"
	"net/http"
	"os"
	"sort"
	"strings"
)

// binary search for string x in slice a
//
// NOTE: slice a must be sorted
func ContainsString(a []string, x string) bool {
	var i = sort.SearchStrings(a, x)
	return i < len(a) && a[i] == x
}

// binary search for int x in slice a
//
// NOTE: slice a must be sorted
func ContainsInt(a []int, x int) bool {
	var i = sort.SearchInts(a, x)
	return i < len(a) && a[i] == x
}

// binary search for float x in slice a
//
// NOTE: slice a must be sorted
func ContainsFloat64(a []float64, x float64) bool {
	var i = sort.SearchFloat64s(a, x)
	return i < len(a) && a[i] == x
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
