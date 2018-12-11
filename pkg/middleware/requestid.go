package middleware

import (
	"github.com/segmentio/ksuid"
	"net/http"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const headerXRequestID = "X-Request-ID"

		// Use existing header if available
		requestID := r.Header.Get(headerXRequestID)

		if requestID == "" {
			// Generate new id
			id, err := ksuid.NewRandom()
			if err != nil {
				requestID = err.Error()
			}
			requestID = id.String()
		}

		// Set header
		w.Header().Set(headerXRequestID, requestID)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
