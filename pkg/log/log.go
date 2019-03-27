package logutil

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

func SetupLogger(dev bool) {
	// Prod
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = "created"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorFieldName = "message"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.With().Caller().Logger()

	if dev {
		// Dev
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(ConsoleWriter{
			Out:           os.Stderr,
			NoColor:       false,
			TimeFormat:    "2006-01-02 15:04:05",
			MarshalIndent: true,
		})
	}
}
