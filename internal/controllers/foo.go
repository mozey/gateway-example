package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/mozey/logutil"
	"log"
)

func Foo(w http.ResponseWriter, r *http.Request) {
	logutil.Debug("Foo")
	fooParam := r.URL.Query().Get("foo")
	if fooParam == "" {
		logutil.Debug("Missing foo")
		msg := Response{Message: "Missing foo"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if fooParam == "panic" {
		log.Panic("oops!")
	}
	msg := Response{Message: fmt.Sprintf("foo: %v", fooParam)}
	json.NewEncoder(w).Encode(msg)
}

