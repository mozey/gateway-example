package controllers

import (
	"net/http"
	"github.com/mozey/logutil"
	"github.com/mozey/gateway/internal/middleware"
)

// Index route handler
func Index(w http.ResponseWriter, r *http.Request) {
	logutil.Debug("Index")
	middleware.Respond(w, r, middleware.ResponseMsg{Message: "It works!"})
}
