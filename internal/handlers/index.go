package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Message string
}

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Message: "It works!"})
}
