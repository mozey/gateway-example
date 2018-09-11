package handlers

import (
	"net/http"
	"log"
	"fmt"
	"github.com/labstack/echo"
	"github.com/mozey/gateway/internal/config"
)

// Foo route handler
func Foo(c echo.Context) error {
	foo := c.QueryParam("foo")
	if foo == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"missing foo in the query string")
	}
	if foo == "panic" {
		//time.Sleep(1 * time.Second)
		// Pass in foo=panic to see the middleware.RecoveryHandler in action
		log.Panic("oops!")
	}
	if foo == "config" {
		conf := config.New()
		resp := Response{
			Message: fmt.Sprintf("conf.Debug %v", conf.Debug)}
		return c.JSON(http.StatusOK, resp)
	}

	// Auth middleware sets "user" on the echo context
	resp := Response{
		Message: fmt.Sprintf("foo: %v user: %v", foo, c.Get("user"))}
	return c.JSON(http.StatusOK, resp)
}
