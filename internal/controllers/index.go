package controllers

import (
	"net/http"
	"github.com/mozey/logutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
	logutil.Debug("Index")
	Respond(w, Response{Message: "It works!"})
}

