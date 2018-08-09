package controllers

import (
	"net/http"
	"github.com/mozey/logutil"
	"log"
	"github.com/mozey/gateway/internal/middleware"
	"fmt"
)

func Foo(w http.ResponseWriter, r *http.Request) {
	logutil.Debug("Foo")
	fooParam := r.URL.Query().Get("foo")
	if fooParam == "" {
		logutil.Debug("Missing foo")
		msg := middleware.ResponseMsg{Message: "Missing foo"}
		middleware.RespondWithCode(http.StatusBadRequest, w, r, msg)
		return
	}
	if fooParam == "panic" {
		// Pass in foo=panic to see the RecoveryHandler in action
		log.Panic("oops!")
	}
	middleware.Respond(
		w, r, middleware.ResponseMsg{Message: fmt.Sprintf("foo: %v", fooParam)})
}

