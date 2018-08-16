package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"fmt"
)

type Response struct {
	Message string
}

func Validator(apiKey string, c echo.Context) (bool, error) {
	if apiKey == "123" {
		c.Set("user", "joe")
		return true, nil
	}
	return false, echo.NewHTTPError(
		http.StatusUnauthorized, fmt.Sprintf("invalid api_key"))
}

func Skipper(c echo.Context) bool {
	path := c.Path()

	if path == "/v1" ||
		path == "/v1/bar" {
		// These endpoints do not require api_key
		return true
	}
	return false
}

func Auth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "query:api_key",
		Validator: Validator,
		Skipper:   Skipper,
	})
}
