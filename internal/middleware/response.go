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
	Message string   `json:"message"`
}

// ResponseStack is a response with a stack trace
type ResponseStack struct {
	Stack   []string `json:"stack"`
}

// ResponseMsgWithStack is a response message with a stack trace
type ResponseWithStack struct {
	Message string   `json:"message"`
	ResponseStack
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

func ExecutionDurationSeconds(r *http.Request) string {
	duration := r.Context().Value(HeaderExecutionDurationS)
	if duration != nil {
		// Inner handler has already called RespondWithCode
		return duration.(string)
	}
	startRaw := r.Context().Value(ContextExecutionStart)
	if startRaw != nil {
		start, err := time.Parse(
			ContextExecutionStartFormat, startRaw.(string))
		if err == nil {
			diff := time.Since(start)
			return fmt.Sprintf("%v", diff.Seconds())
		} else {
			log.Print(err)
		}
	}
	return ""
}

// Respond with StatusOK
func Respond(w http.ResponseWriter, r *http.Request, msg interface{}) {
	RespondWithCode(http.StatusOK, w, r, msg)
}

// RespondWithCode specified
func RespondWithCode(statusCode int, w http.ResponseWriter, r *http.Request, msg interface{}) {
	duration := ExecutionDurationSeconds(r)
	w.Header().Set(HeaderExecutionDurationS, duration)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}
