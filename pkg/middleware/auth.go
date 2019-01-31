package middleware

import (
	"context"
	"github.com/mozey/gateway/pkg/response"
	"github.com/rs/zerolog/log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Skip auth for public paths
		path := r.URL.Path
		switch path {
		case
			"/",

			"/console",

			"/v1",
			"/v1/bar",
			"/v1/status":

			// Call the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Authenticate
		apiKey := r.URL.Query().Get("api_key")
		user := User{}
		if apiKey == "123" {
			user.Name = "joe"
		} else if apiKey == "456" {
			user.Name = "jane"
		} else {
			// http.StatusUnauthorized => Authentication
			// http.StatusForbidden => Authorization
			msg := "Invalid or missing api_key"
			response.JSON(http.StatusUnauthorized, w, r, response.Response{
				Message: msg,
			})
			return
		}

		// Set user on context
		ctx = context.WithValue(ctx, User{}, user)

		// Pass a sub-logger by context
		// https://github.com/rs/zerolog#pass-a-sub-logger-by-context
		logger := log.With().
			Str("user_name", user.Name).
			Str("request_id", w.Header().Get(HeaderXRequestID)).
			Logger()
		ctx = logger.WithContext(ctx)

		// Call the next handler
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
