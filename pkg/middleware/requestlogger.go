package middleware

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		log.Info().
			Str("method", r.Method).
			Str("request_uri", string(r.RequestURI)).
			Msg("-")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
