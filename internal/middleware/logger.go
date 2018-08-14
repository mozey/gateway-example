package middleware

import (
	"net/http"
	"log"
)

// Logger logs all the requests
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Lambda will write stdout to CloudWatch
		defer func() {
			log.Printf(
				"%s\t%s\t%v",
				r.Method,
				r.RequestURI,
				ExecutionDurationSeconds(r),
			)
		}()
		inner.ServeHTTP(w, r)
	})
}
