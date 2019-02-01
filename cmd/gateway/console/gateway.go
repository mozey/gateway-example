package main

import (
	"github.com/apex/gateway"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/console/app"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.New()

	h, cleanup := app.CreateRouter(conf)
	defer cleanup()

	log.Fatal().Err(gateway.ListenAndServe("", h.HTTPHandler))
}
