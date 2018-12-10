package main

import (
	"fmt"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/internal/routes"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	conf := config.New()
	listen := fmt.Sprintf("localhost:%v", conf.Port())
	h := routes.CreateRouter(conf)
	defer h.Handler.Cleanup()
	log.Info().Msgf("listening on %s", listen)
	log.Fatal().Err(http.ListenAndServe(listen, h.Router))
}
