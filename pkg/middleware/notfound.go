package middleware

import (
	"fmt"
	"github.com/mozey/gateway/pkg/response"
	"net/http"
)

func NotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.JSON(http.StatusNotFound, w, r, response.Response{
			Message: fmt.Sprintf("unknown path %s", r.URL.Path),
		})
	})
}
