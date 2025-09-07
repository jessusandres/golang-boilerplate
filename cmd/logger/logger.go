package llog

import (
	"log"
	"lookerdevelopers/boilerplate/cmd/config"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	log.Printf("Initializing logger")

	var logger *zap.Logger

	if config.EnvConfig.Environment == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync() // flushes buffer, if any

	Logger = logger.Sugar()
}
