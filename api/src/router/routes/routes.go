package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents all routes
type Routes struct {
	Uri     string
	Methods string
	Func    func(http.ResponseWriter, *http.Request)
	Auth    bool
}

// Configure adds all routes on the router
func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Func).Methods(route.Methods)
	}
	return r
}
