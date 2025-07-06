package muxX

import (
	"net/http"
	"strings"
)

// HandleStaticDir sets up a Handler for serving static files.
// I always forget how this works, so I wrote this to reduce the boilerplate
// toil.
func HandleStaticDir(mux *http.ServeMux, routePrefix string, serverPrefix string, handleNotFound func(w http.ResponseWriter, r *http.Request), serveDir bool) {
	route := "GET " + routePrefix + "/"
	mux.Handle(route, http.StripPrefix(routePrefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !serveDir && strings.HasSuffix(r.URL.Path, "/") {
			handleNotFound(w, r)
			return
		}

		http.FileServer(http.Dir(serverPrefix)).ServeHTTP(w, r)
	})))
}
