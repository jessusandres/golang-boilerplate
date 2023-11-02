package main

import (
	"log"
	"os"
	"strconv"

	_ "cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/config"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port := strconv.Itoa(config.EnvConfig.Port)

	hostname := "0.0.0.0"

	if localOnly := os.Getenv("LOCAL_ONLY"); localOnly == "true" {
		hostname = "127.0.0.1"
	}

	log.Printf("🌐 Running locally with host: %s and port %s:", hostname, port)

	if err := funcframework.StartHostPort(hostname, port); err != nil {
		log.Fatalf("Error runing in local environment: %v\n", err)
	}
}
