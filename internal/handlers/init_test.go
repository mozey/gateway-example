package handlers_test

import (
	"github.com/mozey/gateway/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var conf *config.Config

func init() {
	var err error
	// Load conf
	conf, err = config.LoadFile("dev")
	if err != nil {
		panic(err)
	}

	// Setup logging
	log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
