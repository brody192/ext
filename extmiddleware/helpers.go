package extmiddleware

import (
	"net/http"
	"strings"
)

func sanitizeURI(uri string) string {
	// double slash `\\`, `//` or even `\/` is absolute uri for browsers and by redirecting request to that uri
	// we are vulnerable to open redirect attack. so replace all slashes from the beginning with single slash
	if len(uri) > 1 && (uri[0] == '\\' || uri[0] == '/') && (uri[1] == '\\' || uri[1] == '/') {
		uri = "/" + strings.TrimLeft(uri, `/\`)
	}
	return uri
}

// return the given query parameter with all leading and trailing white space removed, as defined by Unicode.
func TrimmedQParam(r *http.Request, q string) string {
	var qp = r.URL.Query().Get(q)

	if qp != "" {
		return strings.TrimSpace(qp)
	}

	return qp
}
