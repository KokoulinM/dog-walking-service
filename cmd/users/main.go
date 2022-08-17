package main

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logLevel = flag.String("log-level", "info", "logging level")
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func main() {
	flag.Parse()

	l, err := zerolog.ParseLevel(*logLevel)
	if err != nil {
		log.Error().Caller().Str("log-level", *logLevel).Err(err).Msg("")
		return
	}
	zerolog.SetGlobalLevel(l)
}
