package main

import (
	"github.com/mozey/gateway-echo"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/routes"
)

func main() {
	conf := config.New()
	e, cleanup := routes.CreateMux(conf)
	defer cleanup()
	// TODO Is this the right way to use echo with apex/gateway?
	// https://forum.labstack.com/t/is-echo-v3-compatible-with-http-handlerfunc/523
	e.Logger.Fatal(gateway.Start(e))
}
