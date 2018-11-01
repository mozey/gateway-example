package handlers_test

import (
	"github.com/labstack/gommon/log"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/pkg/handler"
	stdLog "log"
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
	stdLog.SetFlags(stdLog.Ldate | stdLog.Ltime | stdLog.LUTC | stdLog.Lshortfile)
	log.SetLevel(log.DEBUG)
	log.SetHeader(hutil.LogFormat)
}
