package main

import (
	"log"
	"net/http"
	"github.com/mozey/gateway/internal/routes"
	"github.com/mozey/logutil"
	"github.com/gorilla/handlers"
)

func main() {
	h := routes.NewRouter()
	logutil.Debug("Using net/http")

	log.Fatal(http.ListenAndServe(":8080",
		handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true))(h)))
}
