package middleware

import (
	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
)

func Setup(e *echo.Echo) {
	e.Use(Auth())
	e.Use(m.Logger())
	//e.Use(m.LoggerWithConfig(m.LoggerConfig{
	//	Format: "{\"method\"=${method}, \"uri\"=${uri}, \"status\"=${status}}\n",
	//}))
	e.Use(m.Recover())
}
