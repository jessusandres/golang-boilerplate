package routes

import (
	"lookerdevelopers/boilerplate/internal/config"
	incidentsroutes "lookerdevelopers/boilerplate/internal/modules/incident/http/routes"
	apperrors "lookerdevelopers/boilerplate/internal/shared/errors/app"
	"lookerdevelopers/boilerplate/internal/shared/http/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRouterWithDependencies(router gin.IRouter, dependencies *types.RouterDependencies) {

	setupHealthRoutes(router)

	router.GET("/fail", func(c *gin.Context) {
		c.Error(apperrors.NewBadRequestError("Unexpected handler."))
	})

	apiV1Router := router.Group(config.EnvConfig.ApiPrefix + "/v1")

	incidentsroutes.RegisterRoutes(apiV1Router, dependencies.IncidentController)
}

func setupHealthRoutes(router gin.IRouter) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin CQRS API!",
			"status":  "healthy",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "incidents-api",
		})
	})
}
