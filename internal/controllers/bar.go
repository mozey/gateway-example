package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/mozey/logutil"
)

func Bar(w http.ResponseWriter, r *http.Request) {
	logutil.Debug("Bar")
	barParam := r.URL.Query().Get("bar")
	if barParam == "" {
		logutil.Debug("Missing bar")
		msg := Response{Message: "Missing bar"}
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := Response{Message: fmt.Sprintf("bar: %v", barParam)}
	json.NewEncoder(w).Encode(msg)
}

