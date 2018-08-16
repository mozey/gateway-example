package routes

import (
	"github.com/labstack/echo"
	"github.com/mozey/gateway/internal/controllers"
	"github.com/mozey/gateway/internal/middleware"
)

func CreateMux() *echo.Echo {
	e := echo.New()

	middleware.Setup(e)

	e.GET("/v1", controllers.Index)
	e.GET("/v1/foo", controllers.Foo)
	e.GET("/v1/bar", controllers.Bar)

	return e
}
