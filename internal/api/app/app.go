package app

import (
	"compress/gzip"
	"fmt"
	gh "github.com/gorilla/handlers"
	"github.com/mozey/gateway/internal/api/routes"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/api/handlers"
	"github.com/mozey/gateway/pkg/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"runtime/debug"
	"time"
)

// CreateRouter creates a new router.
// It also returns a cleanup function that
// must be called before the server exits
func CreateRouter(conf *config.Config) (h *handlers.Handler, cleanup func()) {
	// Namespace for handlers and services
	h = handlers.NewHandler(conf)

	// Routes
	routes.Console(h)
	routes.Misc(h)

	// Router setup
	h.Router.PanicHandler = middleware.PanicHandler(
		&middleware.PanicHandlerOptions{
			PrintStack: true,
		})
	h.Router.NotFound = middleware.NotFound()

	// Logger
	SetupLogger(conf)

	// Middleware
	SetupMiddleware(h)

	return h, h.Cleanup
}

// SetupLogger configures the logger
func SetupLogger(conf *config.Config) {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.With().Caller().Logger()

	// Prod
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if conf.AwsProfile() == "aws-local" {
		// Dev
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zerolog.ErrorStackMarshaler = func (err error) interface{} {
			// TODO Option for ConsoleWriter to format stack traces?
			fmt.Println(string(debug.Stack()))
			return nil
		}
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    false,
			TimeFormat: time.RFC3339,
		})
	}

	// Add contextual fields to the global logger
	// https://github.com/rs/zerolog#add-contextual-fields-to-the-global-logger
	log.Logger = log.With().Str("global_ctx", "foo").Logger()
}

// SetupMiddleware configures the middleware given a route handler
func SetupMiddleware(h *handlers.Handler) {
	// Middleware in reverse order,
	h.Handler = middleware.Auth(h.Handler)
	//h.Handler = middleware.RequestLogger(h.Handler)
	h.Handler = gh.CompressHandlerLevel(h.Handler, gzip.BestSpeed)
	h.Handler = middleware.RequestID(h.Handler)
}
