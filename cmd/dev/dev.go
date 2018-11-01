package main

import (
	"fmt"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/routes"
)

func main() {
	conf := config.New()
	listen := fmt.Sprintf("localhost:%v", conf.Port())
	e, cleanup := routes.CreateMux(conf)
	defer cleanup()
	e.Logger.Fatal(e.Start(listen))
}
