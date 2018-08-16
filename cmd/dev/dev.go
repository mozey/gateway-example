package main

import (
	"os"
	"fmt"
	"github.com/mozey/gateway/internal/routes"
)

func main() {
	port := os.Getenv("APP_PORT")
	listen := fmt.Sprintf("localhost:%v", port)
	// Start server
	e := routes.CreateMux()
	debug := os.Getenv("APP_DEBUG")
	e.Debug = debug == "true"
	fmt.Println("Debug", e.Debug)
	e.Logger.Fatal(e.Start(listen))
}
