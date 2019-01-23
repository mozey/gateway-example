package main

import (
	"github.com/apex/gateway"
	"github.com/mozey/gateway/internal/app"
	"github.com/mozey/gateway/internal/config"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.New()

	h, cleanup := app.CreateRouter(conf)
	defer cleanup()

	log.Fatal().Err(gateway.ListenAndServe(conf.Port(), h.Handler))
}
