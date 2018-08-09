package middleware

import (
	"net/http"
	"time"
	"fmt"
	"log"
	"encoding/json"
)

const (
	HeaderContentType = "Content-Type"
	HeaderContentTypeJson = "application/json; charset=utf-8"
	HeaderExecutionStart = "X-Execution-Start"
	HeaderExecutionStartFormat = time.RFC3339
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
		w.Header().Set(HeaderExecutionStart,
			time.Now().UTC().Format(HeaderExecutionStartFormat))
		inner.ServeHTTP(w, r)
	})
}

// Respond with StatusOK
func Respond(w http.ResponseWriter, msg interface{}) {
	RespondWithCode(http.StatusOK, w, msg)
}

// RespondWithCode specified
func RespondWithCode(statusCode int, w http.ResponseWriter, msg interface{}) {
	startHeader := w.Header().Get(HeaderExecutionStart)
	w.Header().Del(HeaderExecutionStart)
	start, err := time.Parse(
		HeaderExecutionStartFormat, startHeader)
	if err == nil {
		diff := time.Since(start)
		w.Header().Set(HeaderExecutionDurationS,
			fmt.Sprintf("%v", diff.Seconds()))
	} else {
		log.Print(err)
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}