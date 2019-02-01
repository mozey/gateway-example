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
func JSON(code int, w http.ResponseWriter, r *http.Request, resp interface{}) {
	ctx := r.Context()

	// Marshal indented response JSON,
	// uses type switch to handle different resp types
	var b []byte
	var err error
	indent := "    "
	switch v := resp.(type) {
	case string:
		b, err = json.MarshalIndent(Response{Message: v}, "", indent)
	case error:
		b, err = json.MarshalIndent(Response{Message: v.Error()}, "", indent)
	default:
		b, err = json.MarshalIndent(resp, "", indent)
	}
	if err != nil {
		log.Panic().Err(err)
	}

	// Write headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code) // Must be called after w.Header().Set?

	// Note that request_id and user_name is set on ctx by auth middleware

	// Log request here instead of in middleware,
	// otherwise status code is not logged.
	log.Ctx(ctx).Info().Int("code", code).
		Str("method", r.Method).
		Str("request_uri", string(r.RequestURI)).
		Msg(http.StatusText(code))

	fmt.Fprint(w, string(b))
}

func HTML(code int, w http.ResponseWriter, r *http.Request, b []byte) {
	fmt.Fprint(w, string(b))
}
