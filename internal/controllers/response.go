package controllers

import (
	"encoding/json"
	"time"
	"fmt"
	"net/http"
	"log"
)

// Response contains a simple message response.
type Response struct {
	Message string `json:"message"`
}

func Respond(w http.ResponseWriter, r interface{}) {
	startHeader := w.Header().Get("X-Execution-Start")
	if startHeader != "" {
		// Header will be empty when testing
		start, err := time.Parse(time.RFC3339, startHeader)
		if err == nil {
			diff := time.Since(start)
			w.Header().Set("X-Execution-Duration-s",
				fmt.Sprintf("%v", diff.Seconds()))
		} else {
			log.Print(err)
		}
	}
	json.NewEncoder(w).Encode(r)
}
