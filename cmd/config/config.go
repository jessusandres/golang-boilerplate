package config

import (
	"fmt"
	"log"
	"testing"

	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/types"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/utils"
	"github.com/joho/godotenv"
)

var EnvConfig types.Config

func init() {
	// Load .env by default
	err := godotenv.Load()

	if testing.Testing() {
		fmt.Println("We are testing - not loading envs")
		return
	}

	if err != nil {
		log.Printf("Error loading envs: %s", err.Error())
	}

	errs := utils.ParseEnvSchema(&EnvConfig)

	if len(errs) > 0 {
		log.Println("Error loading environment schema:")

		for _, err := range errs {
			log.Println(err)
		}

		log.Fatal("Environment variable validation failed")
	}

	log.Println("Environment variables loaded successfully")
}
