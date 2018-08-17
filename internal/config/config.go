// TODO This file should be generated from a JSON config

package config

import (
	"os"
	"fmt"
)

var timestamp string
var foo string

type Config struct {
	Timestamp string
	Foo string
}

var conf *Config

// New creates an instance of Config,
// fields are set from private package vars or OS env.
// For dev the config is read from env.
// The prod build must be compiled with ldflags to set the package vars.
// OS env vars will override ldflags if set.
// Config fields correspond to the config file keys less the prefix.
// Use https://github.com/mozey/config to manage the JSON config file
func New() *Config {
	var v string
	v = os.Getenv("APP_TIMESTAMP")
	if v != "" {
		timestamp = v
	}
	v = os.Getenv("APP_FOO")
	if v != "" {
		foo = v
	}

	conf = &Config{
		Timestamp: timestamp,
		Foo: foo,
	}

	return conf
}

// Refresh returns a new Config if the Timestamp has changed.
func Refresh() *Config {
	if conf == nil {
		fmt.Println("init")
		// conf not initialised
		return New()
	}

	timestamp = os.Getenv("APP_TIMESTAMP")
	if conf.Timestamp != timestamp {
		fmt.Println("Timestamp")
		// Timestamp changed, reload config
		return New()
	}

	// No change
	fmt.Println("no change")
	return conf
}
