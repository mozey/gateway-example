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

	port := os.Getenv("APP_PORT")
	listen := fmt.Sprintf("localhost:%v", port)

	logutil.Debugf("Listening on %v", listen)

	log.Fatal(http.ListenAndServe(listen,
		handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true))(h)))
}
