package routes

import (
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/mozey/gateway/internal/middleware"
	pm "github.com/mozey/gateway/pkg/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// CreateRouter creates a new router.
// It also returns a cleanup function that
// must be called before the server exits
func CreateRouter(conf *config.Config) (h *handlers.Handler, cleanup func()) {
	// Namespace for handlers and services
	h = handlers.NewHandler(conf)

	// Routes
	h.Router.HandlerFunc("GET", "/v1", h.Index)
	h.Router.HandlerFunc("GET", "/v1/foo/:foo", h.Foo)
	h.Router.HandlerFunc("GET", "/v1/bar", h.Bar)
	h.Router.HandlerFunc("GET", "/v1/status", h.Status)

	// Router setup
	h.Router.PanicHandler = pm.PanicHandler(&pm.PanicHandlerOptions{
		// Logging stack traces in prod not supported yet
		PrintStack: conf.AwsProfile() == "aws-local",
	})
	h.Router.NotFound = pm.NotFound()

	// Logger setup
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger()
	if conf.AwsProfile() == "aws-local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    false,
			TimeFormat: time.RFC3339,
		})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Middleware
	middleware.Setup(h)

	return h, h.Cleanup
}
