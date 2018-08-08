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
	start, err := time.Parse(time.RFC3339, w.Header().Get("X-Execution-Start"))
	if err != nil {
		log.Panic(err)
	}
	diff := time.Since(start)
	w.Header().Set("X-Execution-Duration-s", fmt.Sprintf("%v", diff.Seconds()))
	json.NewEncoder(w).Encode(r)
}
