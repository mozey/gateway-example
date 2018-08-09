package middleware

import (
	"net/http"
	"time"
	"fmt"
	"log"
	"encoding/json"
	"context"
)

const (
	ContextExecutionStart       = "X-Execution-Start"
	ContextExecutionStartFormat = time.RFC3339

	HeaderContentType        = "Content-Type"
	HeaderContentTypeJson    = "application/json; charset=utf-8"
	HeaderExecutionDurationS = "X-Execution-Duration-S"
)

// ResponseMsg is a simple example of a response message type
type ResponseMsg struct {
	Message string `json:"message"`
}

// Headers set by default for the response
func ResponseHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentType, HeaderContentTypeJson)
		ctx := context.WithValue(
			r.Context(), ContextExecutionStart,
			time.Now().UTC().Format(ContextExecutionStartFormat))
		inner.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Respond with StatusOK
func Respond(w http.ResponseWriter, r *http.Request, msg interface{}) {
	RespondWithCode(http.StatusOK, w, r, msg)
}

// RespondWithCode specified
func RespondWithCode(statusCode int, w http.ResponseWriter, r *http.Request, msg interface{}) {
	startRaw := r.Context().Value(ContextExecutionStart)
	if startRaw != nil {
		start, err := time.Parse(
			ContextExecutionStartFormat, startRaw.(string))
		if err == nil {
			diff := time.Since(start)
			w.Header().Set(HeaderExecutionDurationS,
				fmt.Sprintf("%v", diff.Seconds()))
		} else {
			log.Print(err)
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}
