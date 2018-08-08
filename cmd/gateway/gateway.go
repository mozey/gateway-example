package main

import (
	"log"
	"github.com/apex/gateway"
	"github.com/mozey/gateway/internal/routes"
	"github.com/mozey/logutil"
	"github.com/gorilla/handlers"
)

func main() {
	h := routes.NewRouter()
	logutil.Debug("Using apex/gateway")
	log.Fatal(gateway.ListenAndServe("",
		handlers.RecoveryHandler()(h)))
}




