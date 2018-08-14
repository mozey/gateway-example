package main

import (
	"log"
	"github.com/apex/gateway"
	"github.com/mozey/gateway/internal/routes"
	"github.com/mozey/logutil"
	"github.com/mozey/gateway/internal/middleware"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	h := routes.NewRouter()

	logutil.Debug("mozey-gateway main")

	log.Fatal(gateway.ListenAndServe("",
		middleware.RecoveryHandler(h)))
}




