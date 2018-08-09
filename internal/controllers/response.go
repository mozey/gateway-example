package controllers

import (
	"encoding/json"
	"time"
	"net/http"
	"log"
	"fmt"
	"github.com/mozey/gateway/internal/middleware"
)

// Response contains a simple message response.
type Response struct {
	Message string `json:"message"`
}

func Respond(w http.ResponseWriter, r interface{}) {
	startHeader := w.Header().Get(middleware.HeaderExecutionStart)
	w.Header().Del(middleware.HeaderExecutionStart)
	start, err := time.Parse(
		middleware.HeaderExecutionStartFormat, startHeader)
	if err == nil {
		diff := time.Since(start)
		w.Header().Set(middleware.HeaderExecutionDurationS,
			fmt.Sprintf("%v", diff.Seconds()))
	} else {
		log.Print(err)
	}
	json.NewEncoder(w).Encode(r)
}
