package routes

import (
	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/mozey/gateway/internal/middleware"
)

// CreateMux creates a new instance of echo
// It also returns a cleanup function that
// must be called before the app exits
func CreateMux() (e *echo.Echo, cleanup func()) {
	e = echo.New()

	h := handlers.NewHandler(config.New())
	middleware.Setup(e, h)

	e.GET("/v1", handlers.Index)
	e.GET("/v1/foo", handlers.Foo)
	e.GET("/v1/bar", handlers.Bar)

	return e, h.Cleanup
}
