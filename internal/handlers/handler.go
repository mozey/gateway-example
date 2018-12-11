package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/pkg/services"
	"net/http"
)

// Handler for mozey/gateway
type Handler struct {
	*services.Services
	Config *config.Config
	Router *httprouter.Router
	Handler http.Handler
}

// NewHandler creates a new handler and initialises services
// that are shared between handlers.
// Remember to close services like the database connection by
// calling h.Cleanup before the application exits
func NewHandler(conf *config.Config) (h *Handler) {
	h = &Handler{}
	h.Config = conf
	h.Services = services.NewServices(&services.Options{
		Debug:      conf.Debug(),
		Region:     conf.Region(),
		AwsProfile: conf.AwsProfile(),
	})
	h.Router = httprouter.New()
	// Remember to assign returned handler
	// when wrapping middleware
	h.Handler = h.Router
	return h
}
