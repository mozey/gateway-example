package handlers

import (
	"github.com/labstack/echo"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/pkg/handler"
)

// Handler for  mozey/gateway
type Handler struct {
	*hutil.Handler
	Config *config.Config
}

// Response to be used or extended by handlers
type Response struct {
	Message string `json:"message"`
}

// NewHandler creates a new handler and initialises services
// that are shared between handlers.
// Remember to close services like the database connection by
// calling h.Cleanup before the application exits
func NewHandler(e *echo.Echo, conf *config.Config) (h *Handler) {
	h = &Handler{}
	h.Config = conf
	h.Handler = hutil.NewHandler(e, &hutil.Config{
		Debug:          conf.Debug(),
		Region:         conf.Region(),
		AwsProfile:     conf.AwsProfile(),
	})
	return h
}


