package response

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Response to be used or extended by handlers
type Response struct {
	Message string `json:"message"`
}

// JSON can be used by route handlers to respond to requests
func JSON(code int, w http.ResponseWriter, r *http.Request, i interface{}) {
	j, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		log.Panic().Err(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code) // Must be called after w.Header().Set?

	// Log request here instead of in middleware,
	// otherwise status code is not logged.
	//requestID := w.Header().Get("X-Request-ID")
	log.Info().
		Int("code", code).
		Str("method", r.Method).
		Str("request_uri", string(r.RequestURI)).
		//Str("request_id", string(requestID)).
		Msg("-")

	fmt.Fprint(w, string(j))
}
