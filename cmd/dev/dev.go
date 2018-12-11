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

	h, cleanup := routes.CreateRouter(conf)
	defer cleanup()

	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")

	listen := fmt.Sprintf("localhost:%v", conf.Port())
	log.Info().Msgf("listening on %s", listen)
	log.Fatal().Err(http.ListenAndServe(listen, h.Handler))
}
