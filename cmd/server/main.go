package main

import (
	"log"
	queries "lookerdevelopers/boilerplate/internal/modules/incident/queries/impl"
	"os"
	"strconv"

	"lookerdevelopers/boilerplate/internal/config"
	"lookerdevelopers/boilerplate/internal/modules/incident/adapters/impl"
	"lookerdevelopers/boilerplate/internal/modules/incident/commands/handlers"
	"lookerdevelopers/boilerplate/internal/modules/incident/commands/impl"
	"lookerdevelopers/boilerplate/internal/modules/incident/http/controllers"
	"lookerdevelopers/boilerplate/internal/modules/incident/mappers"
	"lookerdevelopers/boilerplate/internal/modules/incident/queries/handlers"
	"lookerdevelopers/boilerplate/internal/modules/incident/services"
	"lookerdevelopers/boilerplate/internal/shared/cqrs"
	"lookerdevelopers/boilerplate/internal/shared/http/middlewares"
	"lookerdevelopers/boilerplate/internal/shared/http/routes"
	"lookerdevelopers/boilerplate/internal/shared/http/types"
	"lookerdevelopers/boilerplate/internal/shared/instances"
	llog "lookerdevelopers/boilerplate/internal/shared/slog"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("🤖 Loading function configuration...")

	logEnvironmentInfo()
}

func main() {
	port := strconv.Itoa(config.EnvConfig.Port)

	hostname := "0.0.0.0"

	if localOnly := os.Getenv("LOCAL_ONLY"); localOnly == "true" {
		hostname = "127.0.0.1"
	}

	log.Printf("🌐 Running locally with host: %s and port %s:", hostname, port)

	router := buildGin()

	err := router.Run()

	if err != nil {
		panic(err)
	}
}

func buildGin() *gin.Engine {
	log.Printf("🚀 Initializating gin server")

	router := gin.Default()

	// Middlewares
	router.Use(middlewares.BuildState())
	router.Use(middlewares.HandleErr())

	// 404 handler
	router.NoRoute(routes.NotFound)

	// Routes dependencies
	dependencies := setupDependencies()

	// Routes
	routes.BuildRouterWithDependencies(router, dependencies)

	return router
}

func setupDependencies() *types.RouterDependencies {
	log.Println("📦 Setting up dependencies...")

	cqrsSetup := cqrs.NewCQRSSetup()

	setupIncidentsHandlers(cqrsSetup)

	log.Println("✅ CQRS handlers registered")

	// Initialize services and controllers
	incidentService := services.NewIncidentsService(cqrsSetup.CommandBus, cqrsSetup.QueryBus)
	incidentController := controllers.NewIncidentsController(incidentService)

	log.Println("✅ All dependencies setup completed")

	return &types.RouterDependencies{
		IncidentController: incidentController,
	}
}

func setupIncidentsHandlers(setup *cqrs.Setup) {

	incidentsMapper := mappers.IncidentMapper{}
	incidentsRepository := impl.NewGormIncidentImpl(instances.DB, &incidentsMapper)
	createHandler := commandhandlers.NewCreateIncidentHandler(incidentsRepository)
	updateHandler := commandhandlers.NewCreateIncidentHandler(incidentsRepository)
	findHandler := queryhandlers.NewFindIncidentsHandler(incidentsRepository)

	err := setup.CommandBus.Register(commands.CreateIncidentCommand{}.CommandName(), createHandler)
	handleRegisterErr(err)

	err = setup.CommandBus.Register(commands.UpdateIncidentCommand{}.CommandName(), updateHandler)
	handleRegisterErr(err)

	err = setup.QueryBus.Register(queries.FindIncidentsQuery{}.QueryName(), findHandler)

}

func handleRegisterErr(err error) {
	if err != nil {
		llog.Logger.Fatalf("Error registering commands handler: %s", err)
	}
}

func logEnvironmentInfo() {
	log.Println("Using environment:")
	log.Printf("Environment: \t %s", config.EnvConfig.Environment)
	log.Printf("DBHost: \t\t %s", config.EnvConfig.DBHost)
	log.Printf("DBName: \t\t %s", config.EnvConfig.DBName)
	log.Printf("DBMinConnections: \t %d", config.EnvConfig.DBMinConnections)
	log.Printf("DBMaxConnections: \t %d", config.EnvConfig.DBMaxConnections)
	log.Printf("DBSSL: \t\t %s", config.EnvConfig.DBSSLMode)
}
