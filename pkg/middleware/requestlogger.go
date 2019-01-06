package middleware

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := w.Header().Get("X-Request-ID")

		// Log the request (status code is not set yet)
		log.Info().
			Str("method", r.Method).
			Str("request_uri", string(r.RequestURI)).
			Str("request_id", string(requestID)).
			Msg("-")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
