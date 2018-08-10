package main

import (
	"log"
	"net/http"
	"github.com/mozey/gateway/internal/routes"
	"github.com/mozey/logutil"
	"github.com/gorilla/handlers"
	"os"
	"fmt"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	h := routes.NewRouter()
	logutil.Debug("Using net/http")

	port := os.Getenv("APP_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port),
		handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true))(h)))
}
