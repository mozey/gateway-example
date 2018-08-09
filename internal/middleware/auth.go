package middleware

import (
	"net/http"
	"context"
	"fmt"
)

const (
	ContextUserID = "session"
)

type Auth struct {}

func (a *Auth) Authorise(key string) (userID string, err error) {
	if key == "123" {
		return "joe", nil
	}
	return "", fmt.Errorf("invalid key: %v", key)
}

func WithAuth(a Auth, inner http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		userID, err := a.Authorise(key)
		if err != nil {
			RespondWithCode(http.StatusUnauthorized,
				w, r, ResponseMsg{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserID, userID)
		inner.ServeHTTP(w, r.WithContext(ctx))
	})
}

