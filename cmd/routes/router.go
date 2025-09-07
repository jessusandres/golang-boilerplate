package routes

import (
	"lookerdevelopers/boilerplate/cmd/config"
	apperrors "lookerdevelopers/boilerplate/cmd/errors/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRouter(router gin.IRouter) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/fail", func(c *gin.Context) {
		c.Error(apperrors.NewBadRequestError("Unexpected handler."))
	})

	apiRouter := router.Group(config.EnvConfig.ApiPrefix)

	IncidentsRouter(apiRouter)
}
