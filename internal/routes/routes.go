package routes

import (
	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/mozey/gateway/internal/middleware"
)

// CreateMux creates a new instance of echo
func CreateMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Auth())
	e.Use(m.Logger())
	//e.Use(m.LoggerWithConfig(m.LoggerConfig{
	//	Format: "{\"method\"=${method}, \"uri\"=${uri}, \"status\"=${status}}\n",
	//}))
	e.Use(middleware.Recover())

	e.GET("/v1", handlers.Index)
	e.GET("/v1/foo", handlers.Foo)
	e.GET("/v1/bar", handlers.Bar)

	return e
}
