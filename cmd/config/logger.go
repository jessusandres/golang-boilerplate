package config

import (
	"go.uber.org/zap"
	"log"
)

var Logger *zap.SugaredLogger

func init() {
	log.Printf("Initializing logger")

	var logger *zap.Logger

	if EnvConfig.Environment == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync() // flushes buffer, if any

	Logger = logger.Sugar()
}
