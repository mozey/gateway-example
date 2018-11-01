package handlers

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
)

// Bar route handler
func (h *Handler) Bar(c echo.Context) error {
	resp := Response{Message: fmt.Sprintf("no api_key required")}
	return c.JSON(http.StatusOK, resp)
}

