package middleware

import (
	"github.com/segmentio/ksuid"
	"net/http"
)

const HeaderXRequestID = "X-Request-ID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use existing header if available
		requestID := r.Header.Get(HeaderXRequestID)

		if requestID == "" {
			// Generate new id
			id, err := ksuid.NewRandom()
			if err != nil {
				requestID = err.Error()
			}
			requestID = id.String()
		}

		// Header
		w.Header().Set(HeaderXRequestID, requestID)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
