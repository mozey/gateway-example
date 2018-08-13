package middleware

import (
	"net/http"
	"time"
	"log"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		// Lambda will write stdout to CloudWatch
		log.Printf(
			"%s\t%s\t%v",
			r.Method,
			r.RequestURI,
			time.Since(start).Seconds(),
		)
	})
}
