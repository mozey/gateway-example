package handlers

import (
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/handler"
)

// Handler for gateway/console
type Handler struct {
	*handler.Handler
}

// NewHandler creates a new top level handler
func NewHandler(conf *config.Config) (h *Handler) {
	h = &Handler{}
	h.Handler = handler.NewHandler(conf)
	return h
}
