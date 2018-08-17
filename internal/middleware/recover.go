package middleware

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
	"github.com/mozey/logutil"
)

// Recover returns a middleware which recovers from panics anywhere in the chain
// and handles the control to the centralized HTTPErrorHandler.
func Recover() echo.MiddlewareFunc {
	return RecoverWithConfig(m.DefaultRecoverConfig)
}

// RecoverWithConfig returns a Recover middleware with config.
// See: `Recover()`.
func RecoverWithConfig(config m.RecoverConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = m.DefaultRecoverConfig.Skipper
	}
	if config.StackSize == 0 {
		config.StackSize = m.DefaultRecoverConfig.StackSize
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)
					if !config.DisablePrintStack {
						// Don't print stack trace inside JSON,
						// it's really hard to read without newlines!
						c.Logger().Printf("[PANIC RECOVER] %v\n", err)
						logutil.Debug(string(stack[:length]))
					}
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
