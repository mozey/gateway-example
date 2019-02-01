package routes

import "github.com/mozey/gateway/internal/console/handlers"

func Console(h *handlers.Handler) {
	h.Router.HandlerFunc("GET", "/console", h.Console)
	h.Router.HandlerFunc("GET", "/console/static/*file", h.Static)
}
