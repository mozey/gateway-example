package middleware

import (
	"context"
	"github.com/mozey/gateway/pkg/response"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for public paths
		path := r.URL.Path
		switch path {
		case
			"/",
			"/v1",
			"/v1/bar",
			"/v1/status":
			// Call the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Authenticate
		apiKey := r.URL.Query().Get("api_key")
		if apiKey != "123" {
			// http.StatusUnauthorized => Authentication
			// http.StatusForbidden => Authorization
			msg := "Invalid or missing api_key"
			response.JSON(http.StatusUnauthorized, w, r, response.Response{
				Message: msg,
			})
			return
		}

		user := User{Name: "joe"}

		ctx := r.Context()
		ctx = context.WithValue(ctx, User{}, user)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
