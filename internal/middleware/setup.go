package middleware

import (
	"fmt"
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/mozey/gateway/internal/handlers"
	"net/http"
)

func Validator(apiKey string, c echo.Context) (bool, error) {
	if apiKey == "123" {
		c.Set("user", "joe")
		return true, nil
	}
	return false, echo.NewHTTPError(
		http.StatusUnauthorized, fmt.Sprintf("invalid api_key"))
}

// Skipper lists endpoints that do not require validation
func Skipper(c echo.Context) bool {
	path := c.Path()

	// Skip auth validator for these routes
	if path == "/" ||
		path == "/v1" ||
		path == "/v1/bar" ||
		path == "/v1/status" {
		return true
	}

	return false
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := http.StatusText(http.StatusInternalServerError)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = fmt.Sprintf("%s", he.Message)
	}
	type response struct {
		Message string `json:"message"`
	}
	resp := response{
		Message: message,
	}
	if err := c.JSON(code, resp); err != nil {
		c.Logger().Error(err)
	}
	// Request logger already prints this?
	//c.Logger().Error(err)
}

// Setup middleware
func Setup(e *echo.Echo, h *handlers.Handler) {
	// Create unique ID for every request
	e.Use(em.RequestID())
	// Request logger
	loggerConfig := em.DefaultLoggerConfig
	loggerConfig.Format = `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		`"method":"${method}","uri":"${uri}","path":"${path}","status":${status},"error":"${error}","latency":${latency},` +
		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out}}` + "\n"
	e.Use(em.LoggerWithConfig(loggerConfig))

	// Auth
	e.Use(em.KeyAuthWithConfig(em.KeyAuthConfig{
		KeyLookup: "query:api_key",
		Validator: Validator,
		Skipper:   Skipper,
	}))

	// Modified recover middleware so printed stack trace is easier to read
	recoverConfig := em.DefaultRecoverConfig
	if h.Config.AwsProfile() == "aws-local" {
		// Don't print stack for other go-routines
		recoverConfig.DisableStackAll = true
		e.Use(em.RecoverWithConfig(recoverConfig))
	} else {
		// Don't print any stack trace
		recoverConfig.DisablePrintStack = true
		e.Use(em.RecoverWithConfig(recoverConfig))
	}

	// Custom error handling
	// https://echo.labstack.com/guide/error-handling
	e.HTTPErrorHandler = customHTTPErrorHandler
}