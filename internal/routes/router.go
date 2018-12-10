package routes

import (
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/handlers"
)

// CreateRouter creates a new router.
// It also returns a cleanup function that
// must be called before the app exits
func CreateRouter(conf *config.Config) (h *handlers.Handler) {
	h = handlers.NewHandler(conf)
	//middleware.Setup(e, h)

	h.Router.HandlerFunc("GET", "/v1", h.Index)
	h.Router.HandlerFunc("GET", "/v1/foo", h.Foo)
	h.Router.HandlerFunc("GET", "/v1/bar", h.Bar)
	h.Router.HandlerFunc("GET", "/v1/status", h.Status)

	return h
}
