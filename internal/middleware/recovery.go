package middleware

import (
	"net/http"
	"runtime/debug"
	"github.com/mozey/logutil"
	"encoding/json"
	"fmt"
	"strings"
)

var respondWithStack = false
func RespondWithStack(toggle bool) {
	respondWithStack = toggle
}

// RecoveryHandler is HTTP middleware that recovers from a panic,
// writes http.StatusInternalServerError and the panic err,
// optionally prints a stack trace,
// and continues to the next handler
// Inspired by
// [gorilla/handlers](https://github.com/gorilla/handlers/blob/master/recovery.go)
func RecoveryHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				msg := fmt.Sprintf("%v", err)
				if logutil.DebugEnabled() || respondWithStack {
					// Log stack to debug
					stack := debug.Stack()
					s := string(stack)
					logutil.Debug(s)
					// Include stack in response
					if respondWithStack {
						r := ResponseWithStack{Message: msg}
						s = strings.Replace(s, "\t", "    ", -1)
						r.Stack = strings.Split(s, "\n")
						json.NewEncoder(w).Encode(r)
					}
				}
				if !respondWithStack {
					// Respond with err message only
					r := ResponseMsg{Message: msg}
					json.NewEncoder(w).Encode(r)
				}
			}
		}()
		inner.ServeHTTP(w, r)
	})
}
