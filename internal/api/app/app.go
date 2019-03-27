package app

import (
	"compress/gzip"
	gh "github.com/gorilla/handlers"
	"github.com/mozey/gateway/internal/api/handlers"
	"github.com/mozey/gateway/internal/api/routes"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/pkg/log"
	"github.com/mozey/gateway/pkg/middleware"
)

// CreateRouter creates a new router.
// It also returns a cleanup function that
// must be called before the server exits
func CreateRouter(conf *config.Config) (h *handlers.Handler, cleanup func()) {
	// Namespace for handlers and services
	h = handlers.NewHandler(conf)

	// Routes
	routes.Misc(h)

	// Router setup
	h.Router.PanicHandler = middleware.PanicHandler(
		&middleware.PanicHandlerOptions{
			PrintStack: true,
		})
	h.Router.NotFound = middleware.NotFound()

	// Logger
	logutil.SetupLogger(conf.AwsProfile() == "aws-local")

	// Middleware
	SetupMiddleware(h)

	return h, h.Cleanup
}

// SetupMiddleware configures the middleware given a route handler
func SetupMiddleware(h *handlers.Handler) {
	// Middleware in reverse order,
	h.HTTPHandler = middleware.Auth(h.HTTPHandler)
	//h.HTTPHandler = middleware.RequestLogger(h.HTTPHandler)
	h.HTTPHandler = gh.CompressHandlerLevel(h.HTTPHandler, gzip.BestSpeed)
	h.HTTPHandler = middleware.RequestID(h.HTTPHandler)
}
