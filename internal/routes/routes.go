package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mozey/gateway/internal/controllers"
	"github.com/mozey/gateway/internal/middleware"
)

// Route defines a route
type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes array used to create a router
type Routes []Route

var routes = Routes{
	Route{
		"/", controllers.Index,
	},

	Route{
		"/v1/foo{ignore:.*}", controllers.Foo,
	},
	Route{
		"/v1/bar{ignore:.*}",
		middleware.WithAuth(
			middleware.Auth{}, http.HandlerFunc(controllers.Bar)),
	},

	// Echo request when no match
	Route{
		"/{everything:.*}", controllers.Echo,
	},
}

// NewRouter returns a new router instance
func NewRouter() *mux.Router {
	// StrictSlash must be false otherwise index is not loaded consistently
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.Logger(handler)
		handler = middleware.ResponseHeaders(handler)
		router.
			// Paths should not be case sensitive
			Path(route.Path).
			// Don't allow duplicate paths in the routes definition
			Name(route.Path).
			Handler(handler)
	}
	return router
}
