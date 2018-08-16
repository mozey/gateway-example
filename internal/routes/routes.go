package routes

import (
	"github.com/labstack/echo"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/mozey/gateway/internal/middleware"
)

// CreateMux creates a new instance of echo
func CreateMux() *echo.Echo {
	e := echo.New()

	middleware.Setup(e)

	e.GET("/v1", handlers.Index)
	e.GET("/v1/foo", handlers.Foo)
	e.GET("/v1/bar", handlers.Bar)

	return e
}
