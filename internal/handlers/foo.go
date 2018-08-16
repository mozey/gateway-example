package handlers

import (
	"net/http"
	"log"
	"fmt"
	"github.com/labstack/echo"
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
	resp := Response{
		Message: fmt.Sprintf("foo: %v user: %v", foo, c.Get("user"))}
	return c.JSON(http.StatusOK, resp)
}
