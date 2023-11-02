package function

import (
	"log"
	"net/http"

	"lookerdevelopers/boilerplate/cmd/config"
	"lookerdevelopers/boilerplate/cmd/middlewares"
	"lookerdevelopers/boilerplate/cmd/routes"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("🤖 Loading function configuration...")

	log.Println("Using environment:")
	log.Printf("FunctionTarget: \t %s", config.EnvConfig.FunctionTarget)
	log.Printf("Environment: \t %s", config.EnvConfig.Environment)
	log.Printf("DBHost: \t\t %s", config.EnvConfig.DBHost)
	log.Printf("DBName: \t\t %s", config.EnvConfig.DBName)
	log.Printf("DBMinConnections: \t %d", config.EnvConfig.DBMinConnections)
	log.Printf("DBMaxConnections: \t %d", config.EnvConfig.DBMaxConnections)
	log.Printf("DBSSL: \t\t %s", config.EnvConfig.DBSSLMode)

	log.Printf("🚀 Initializating gin server")

	router := buildGin()

	buildFunction(router)
}

func buildGin() *gin.Engine {
	router := gin.Default()

	router.NoRoute(routes.NotFound)

	router.Use(middlewares.BuildState())
	router.Use(middlewares.HandleErr())

	routes.BuildRouter(router)

	return router
}

func buildFunction(router *gin.Engine) {
	functionTarget := config.EnvConfig.FunctionTarget

	functions.HTTP(functionTarget, func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	})
}
