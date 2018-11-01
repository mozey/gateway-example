package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// IndexResponse is the index handler
type IndexResponse struct {
	Message string
	Version string
}

// Index can be used to check if the server is available
func (h *Handler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, IndexResponse{
		Message: "It works!!",
		Version: h.Config.Version(),
	})
}
