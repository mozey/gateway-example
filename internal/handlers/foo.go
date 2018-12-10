package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mozey/gateway/internal/config"
	"log"
	"net/http"
)

// Foo route handler
func (h *Handler) Foo(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	foo := params.ByName("foo")
	if foo == "" {
		w.WriteHeader(http.StatusUnauthorized)
		RespondJSON(w, r, Response{
			Message: "missing foo in the query string",
		})
		return
	}
	if foo == "panic" {
		//time.Sleep(1 * time.Second)
		// Pass in foo=panic to see the middleware.RecoveryHandler in action
		log.Panic("oops!")
	}
	if foo == "config" {
		conf := config.New()
		resp := Response{
			Message: fmt.Sprintf("conf.Debug %v", conf.Debug())}
		RespondJSON(w, r, resp)
		return
	}

	// Auth middleware sets "user" on the echo context
	resp := Response{
		Message: fmt.Sprintf("foo: %v user: %v", foo,
			r.Context().Value("user"))}
	RespondJSON(w, r, resp)
	return
}
