package main

import (
	"github.com/mozey/gateway-wrapper/httprouter"
	"github.com/mozey/gateway/internal/app"
	"github.com/mozey/gateway/internal/config"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.New()

	h, cleanup := app.CreateRouter(conf)
	defer cleanup()

	log.Fatal().Err(gwrapper.Start(h.Handler))
}
