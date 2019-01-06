package middleware

import (
	"compress/gzip"
	gh "github.com/gorilla/handlers"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/mozey/gateway/pkg/middleware"
)

// Setup middleware
func Setup(h *handlers.Handler) {
	// Middleware in reverse order,
	h.Handler = middleware.Auth(h.Handler)
	//h.Handler = middleware.RequestLogger(h.Handler)
	h.Handler = gh.CompressHandlerLevel(h.Handler, gzip.BestSpeed)
	h.Handler = middleware.RequestID(h.Handler)
}
