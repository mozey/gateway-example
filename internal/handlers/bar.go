package handlers

import (
	"fmt"
	"net/http"
)

// Bar route handler
func (h *Handler) Bar(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: fmt.Sprintf("no api_key required")}
	RespondJSON(w, r, resp)
}

