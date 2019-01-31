package routes

import "github.com/mozey/gateway/internal/api/handlers"

func Misc(h *handlers.Handler) {
	h.Router.HandlerFunc("GET", "/v1", h.Index)
	h.Router.HandlerFunc("GET", "/v1/foo/:foo", h.Foo)
	h.Router.HandlerFunc("GET", "/v1/bar", h.Bar)
	h.Router.HandlerFunc("GET", "/v1/status", h.Status)
}
