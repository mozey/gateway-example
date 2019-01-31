package handlers

import (
	"github.com/mozey/gateway/pkg/response"
	"net/http"
)

// Bar route handler
func (h *Handler) Console(w http.ResponseWriter, r *http.Request) {
	response.JSON(http.StatusOK, w, r, "todo")
}

