package handlers_test

import (
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/logutil"
	"log"
)

var conf *config.Config

func init() {
	var err error
	// Load conf
	conf, err = config.LoadFile("dev")
	if err != nil {
		log.Panic(err)
	}
	// Avoid setting env vars in IDE
	logutil.SetDebug(conf.Debug == "true")
	// Include file name and line number in logs
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}
