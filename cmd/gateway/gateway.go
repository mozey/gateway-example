package main

import (
	"github.com/mozey/gateway-wrapper/httprouter"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/routes"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.New()

	h, cleanup := routes.CreateRouter(conf)
	defer cleanup()

	log.Fatal().Err(gwrapper.Start(h.Handler))
}
