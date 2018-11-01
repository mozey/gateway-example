package hutil

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	stdLog "log"
	"path"
	"runtime"
)

// Use a struct with fields for db, s3, sqs etc.
// Handlers should be methods on this struct
// https://github.com/labstack/echo/issues/568#issuecomment-227532114

// "Context is at the request level so it doesn’t
// make sense to store application level data in it"
// https://forum.labstack.com/t/where-to-keep-db-connection/350/4

// LogFormat header for writing and parsing logs.
// Unfortunately github.com/labstack/gommon/log doesn't allow
// customising the skip parameter making it hard to get the line right.
// Also see the PrependLogArgs function
//const LogFormat = "${time_rfc3339_nano} ${level} ${short_file} ${line}"
const LogFormat = "${time_rfc3339_nano} ${level}"
const CallDepth = 2
const LogIndent = "    "

// Config for shared handler
type Config struct {
	Debug          string
	Region         string
	AwsProfile     string
}

// Handler with shared fields.
// Embed this type in internal handlers
type Handler struct {
	Config       *Config
}

// NewHandler creates a new shared handler instance
func NewHandler(e *echo.Echo, config *Config) (h *Handler) {
	h = &Handler{}
	h.Config = config
	h.SetupLogging(e)
	return h
}

// Cleanup function must be called before the application exits
func (h *Handler) Cleanup() {
	// Close db connections etc
}

// SetupLogging given an echo instance
func (h *Handler) SetupLogging(e *echo.Echo) {
	logger := e.Logger
	if l, ok := logger.(*log.Logger); ok {
		l.SetHeader(LogFormat)
	}
	// Probably not required, but set the standard logger flags anyway.
	// It might be useful for tracking down errant calls to stdLog
	stdLog.SetFlags(
		stdLog.Ldate | stdLog.Ltime | stdLog.LUTC | stdLog.Lshortfile)
}

// LogContext is a helper function to prepend logger args with a request ID.
// A single request can print multiple log lines possibly not in sequence,
// and therefore the request ID is essential.
func LogContext(c echo.Context) (prefix string) {
	// Using the same id logic as github.com/labstack/echo/middleware/logger.go
	id := c.Request().Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = c.Response().Header().Get(echo.HeaderXRequestID)
	}
	if id == "" {
		id = "0"
	}

	return fmt.Sprintf("%s", id)
}

// PrependLogArgs uses LogContext to generate an ID prefix for the log line
func PrependLogArgs(c echo.Context, args ...interface{}) (newArgs []interface{}) {
	newArgs = make([]interface{}, 0, len(args)+1)
	// Get file and line
	_, file, line, ok := runtime.Caller(CallDepth)
	if !ok {
		file = "???"
		line = 0
	}
	// Prepend context
	newArgs = append(newArgs, fmt.Sprintf(
		"%s %d %s", path.Base(file), line, LogContext(c)))
	// Add the rest of the args
	return append(newArgs, args...)
}

// Debug log
func (h *Handler) Debug(c echo.Context, args ...interface{}) {
	newArgs := PrependLogArgs(c, args...)
	c.Logger().Debug(newArgs...)
}

// Debug log with string formatting
func (h *Handler) Debugf(c echo.Context, format string, args ...interface{}) {
	newArgs := PrependLogArgs(c, args...)
	c.Logger().Debugf("%s"+format, newArgs...)
}

// Debugj works the same as Debug,
// except it take a single argument that is marshaled to JSON.
// All newlines are removed in production,
// they create multiple log lines in AWS Lambda
func (h *Handler) Debugj(c echo.Context, msg string, arg interface{}) {
	var b []byte
	var err error
	if h.Config.AwsProfile == "aws-local" {
		b, err = json.MarshalIndent(arg, "", LogIndent)
	} else {
		b, err = json.Marshal(arg)
	}
	if err == nil {
		newArgs := PrependLogArgs(c, msg, " ", string(b))
		c.Logger().Debug(newArgs...)
	} else {
		// Marshal failed, pass argument as is
		newArgs := PrependLogArgs(c, msg, " ", arg)
		c.Logger().Debug(newArgs...)
	}
}

// Error log
func (h *Handler) Error(c echo.Context, args ...interface{}) {
	newArgs := PrependLogArgs(c, args...)
	c.Logger().Error(newArgs...)
}

// Info log
func (h *Handler) Info(c echo.Context, args ...interface{}) {
	newArgs := PrependLogArgs(c, args...)
	c.Logger().Info(newArgs...)
}
