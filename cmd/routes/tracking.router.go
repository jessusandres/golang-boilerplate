package routes

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/controllers"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/dto"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/instances"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/middlewares"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/services"
	"github.com/gin-gonic/gin"
)

func TrackingRouter(router gin.IRouter) {

	trackingRouter := router.Group("/tracking")

	trackingService := services.NewTrackingService(instances.DB)
	trackingController := controllers.NewTrackingController(trackingService)

	trackingRouter.PATCH("", middlewares.ValidateJSON[dto.TrackingPatchDto](), trackingController.Patch)
}
