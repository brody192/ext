package exthandler

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/brody192/ext/extutil"
	"github.com/brody192/ext/extvar"

	"github.com/go-chi/chi/v5"
)

// sets up a static file server at the given path with and accepts a sub dir
// will panic if can't sub
func FileServerSub(r chi.Router, path string, fsys fs.FS, dir string) {
	FileServer(r, path, extutil.MustSubFS(fsys, dir))
}

// sets up a static file server from an embeded fs at the given path with and accepts a sub dir
// will panic if can't sub
func FileServerEmbeded(r chi.Router, path string, embfs embed.FS, dir string) {
	FileServer(r, path, extutil.MustSubFS(embfs, dir))
}

// sets up a static file server at the given path
func FileServer(r chi.Router, path string, root fs.FS) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(http.FS(root)))
		fs.ServeHTTP(w, r)
	})
}

// adds matching routes to the router with methods specified in the methods slice
func Match(r chi.Router, methods []string, pattern string, handler http.HandlerFunc) {
	for _, method := range methods {
		if extutil.ContainsStrings(extvar.Methods, method) == false {
			panic("method: " + method + " is not a valid method")
		}
		r.Method(method, pattern, handler)
	}
}
