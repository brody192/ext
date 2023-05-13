package exthandler

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/brody192/ext/extutil"

	"github.com/go-chi/chi/v5"
)

// sets up a static file server at the given path with and accepts a sub dir
//
// will panic if can't sub
func FileServerSub(r chi.Router, path string, fsys fs.FS, dir string, browse bool) {
	FileServer(r, path, extutil.MustSubFS(fsys, dir), browse)
}

// sets up a static file server from an embeded fs at the given path with and accepts a sub dir
//
// will panic if can't sub
func FileServerEmbeded(r chi.Router, path string, embfs embed.FS, dir string, browse bool) {
	FileServer(r, path, extutil.MustSubFS(embfs, dir), browse)
}

// sets up a static file server at the given path
func FileServer(r chi.Router, path string, root fs.FS, browse bool) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	if strings.HasSuffix(path, "*") == false {
		path += "*"
	}

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var rctx = chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		var fs = http.StripPrefix(pathPrefix, http.FileServer(extutil.FileSystem{Fs: http.FS(root)}))
		if browse {
			fs = http.StripPrefix(pathPrefix, http.FileServer(http.FS(root)))
		}
		fs.ServeHTTP(w, r)
	})
}

// adds matching routes to the router with methods specified in the methods slice
func Match(r chi.Router, methods []string, pattern string, handler http.HandlerFunc) {
	for _, method := range methods {
		if extutil.IsValidMethod(method) == false {
			panic("method: " + method + " is not a valid method")
		}
		r.Method(method, pattern, handler)
	}
}

// a MethodNotAllowed handler that returns http.StatusMethodNotAllowed body text and status code
func MethodNotAllowedStatusText(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
