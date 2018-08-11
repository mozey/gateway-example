package controllers

import (
	"net/http"
	"fmt"
	"github.com/mozey/logutil"
	"github.com/mozey/gateway/internal/middleware"
)

// Bar route handler
func Bar(w http.ResponseWriter, r *http.Request) {
	logutil.Debug("Bar")
	user := r.Context().Value(middleware.ContextUserID)
	middleware.Respond(w, r,
		middleware.ResponseMsg{Message: fmt.Sprintf("user = %v", user)})
}
