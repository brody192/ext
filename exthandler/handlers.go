package exthandler

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
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
func MatchMethods(r chi.Router, methods []string, pattern string, handler http.HandlerFunc) {
	for _, method := range methods {
		if extutil.IsValidMethod(method) == false {
			panic("method: " + method + " is not a valid method")
		}

		r.Method(method, pattern, handler)
	}
}

// adds matching routes to the router with patterns specified in the patterns slice
func MatchPatterns(r chi.Router, method string, patterns []string, handler http.HandlerFunc) {
	if extutil.IsValidMethod(method) == false {
		panic("method: " + method + " is not a valid method")
	}

	for _, pattern := range patterns {
		r.Method(method, pattern, handler)
	}
}

// adds matching routes and methods to the router with methods specified in the methods and patterns slice
func MatchMethodsPatterns(r chi.Router, methods []string, patterns []string, handler http.HandlerFunc) {
	for _, pattern := range patterns {
		for _, method := range methods {
			if extutil.IsValidMethod(method) == false {
				panic("method: " + method + " is not a valid method")
			}

			r.Method(method, pattern, handler)
		}
	}
}

// a MethodNotAllowed handler that returns http.StatusMethodNotAllowed body text and status code
func MethodNotAllowedStatusText(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

// registers additional methods with a trailing slash that don't already have a trailing slash
func RegisterTrailing(r *chi.Mux) {
	if len(r.Routes()) == 0 {
		panic("no routes registered on mux")
	}

	registerTrailingWithPrefix(r, r.Routes(), "")
}

func registerTrailingWithPrefix(r *chi.Mux, routes []chi.Route, prefix string) {
	var flatRoutes = make(map[string]bool, len(routes))

	for _, route := range routes {
		flatRoutes[route.Pattern] = true
	}

	for _, route := range routes {
		if route.SubRoutes != nil {
			var cleanPrefix = strings.TrimRight(route.Pattern, "*")
			var newPrefix = path.Join(prefix, cleanPrefix)
			registerTrailingWithPrefix(r, route.SubRoutes.Routes(), newPrefix)
		}

		if flatRoutes[route.Pattern+"/"] || flatRoutes[route.Pattern+"/*"] ||
			strings.HasSuffix(route.Pattern, "/") || strings.HasSuffix(route.Pattern, "/*") {
			continue
		}

		for method, handler := range route.Handlers {
			var path = path.Join(prefix, route.Pattern)
			r.Method(method, path+"/", handler)
		}
	}
}
