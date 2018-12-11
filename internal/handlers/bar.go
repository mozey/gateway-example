package handlers

import (
	"fmt"
	"github.com/mozey/gateway/pkg/response"
	"net/http"
)

// Bar route handler
func (h *Handler) Bar(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{Message: fmt.Sprintf("no api_key required")}
	response.JSON(http.StatusOK, w, r, resp)
}

