package middleware

import (
	"net/http"
	"time"
)

const (
	HeaderContentType = "Content-Type"
	HeaderContentTypeJson = "application/json; charset=utf-8"
	HeaderExecutionStart = "X-Execution-Start"
	HeaderExecutionStartFormat = time.RFC3339
	HeaderExecutionDurationS = "X-Execution-Duration-S"
)

// Headers set by default for the response
func ResponseHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentType, HeaderContentTypeJson)
		w.Header().Set(HeaderExecutionStart,
			time.Now().UTC().Format(HeaderExecutionStartFormat))
		inner.ServeHTTP(w, r)
	})
}
