package routes

import (
	"github.com/labstack/echo"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/mozey/gateway/internal/middleware"
)

// CreateMux creates a new instance of echo.
// It also returns a cleanup function that
// must be called before the app exits
func CreateMux(conf *config.Config) (e *echo.Echo, cleanup func()) {
	e = echo.New()

	e.Debug = conf.Debug() == "true"

	h := handlers.NewHandler(e, conf)
	middleware.Setup(e, h)

	e.GET("/v1", h.Index)
	e.GET("/v1/foo", h.Foo)
	e.GET("/v1/bar", h.Bar)

	return e, h.Cleanup
}
