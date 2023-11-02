package routes

import (
	"github.com/gin-gonic/gin"
	"lookerdevelopers/boilerplate/cmd/controllers"
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/instances"
	"lookerdevelopers/boilerplate/cmd/middlewares"
	"lookerdevelopers/boilerplate/cmd/services"
)

func TrackingRouter(router gin.IRouter) {

	trackingRouter := router.Group("/tracking")

	trackingService := services.NewTrackingService(instances.DB)
	trackingController := controllers.NewTrackingController(trackingService)

	trackingRouter.PATCH("", middlewares.ValidateJSON[dto.TrackingPatchDto](), trackingController.Patch)
}
