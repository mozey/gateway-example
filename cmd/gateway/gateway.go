package main

import (
	"github.com/mozey/gateway/internal/routes"
	"github.com/mozey/gateway-echo"
)

func main() {
	// Start server
	e := routes.CreateMux()
	// TODO Is this the right way to use echo with apex/gateway?
	// https://forum.labstack.com/t/is-echo-v3-compatible-with-http-handlerfunc/523
	e.Logger.Fatal(gateway.Start(e))
}
