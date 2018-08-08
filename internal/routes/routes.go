package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mozey/gateway/internal/controllers"
	"github.com/mozey/gateway/internal/middleware"
)

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"/", controllers.Index,
	},

	Route{
		"/foo{ignore:.*}", controllers.Foo,
	},
	Route{
		"/bar{ignore:.*}", controllers.Bar,
	},

	// Echo request when no match
	Route{
		"/{everything:.*}", controllers.Echo,
	},
}


func NewRouter() *mux.Router {
	// StrictSlash must be false otherwise index is not loaded consistently
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		// TODO Logger middleware?
		//handler = middleware.Logger(handler)
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
