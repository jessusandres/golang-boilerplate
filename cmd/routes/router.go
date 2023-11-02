package routes

import (
	"github.com/gin-gonic/gin"
	"lookerdevelopers/boilerplate/cmd/apperrors"
	"lookerdevelopers/boilerplate/cmd/config"
	"net/http"
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

	TrackingRouter(apiRouter)
}
