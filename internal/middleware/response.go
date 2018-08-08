package middleware

import (
	"net/http"
	"time"
)

// Headers set by default for the response
func ResponseHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Execution-Start", time.Now().UTC().Format(time.RFC3339))
		inner.ServeHTTP(w, r)
	})
}
