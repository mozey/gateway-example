package handlers

import (
	"github.com/mozey/gateway/pkg/response"
	"net/http"
)

// IndexResponse is the index handler
type IndexResponse struct {
	Message string
	Version string
}

// Index can be used to check if the server is available
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	response.JSON(http.StatusOK, w, r, IndexResponse{
		Message: "It works!!",
		Version: h.Config.VersionApi(),
	})
}

