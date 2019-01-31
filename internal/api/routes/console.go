package routes

import "github.com/mozey/gateway/internal/api/handlers"

func Console(h *handlers.Handler) {
	h.Router.HandlerFunc("GET", "/console", h.Console)
}
