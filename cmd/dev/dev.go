package main

import (
	"fmt"
	"github.com/mozey/gateway/internal/app"
	"github.com/mozey/gateway/internal/config"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	conf := config.New()

	h, cleanup := app.CreateRouter(conf)
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
