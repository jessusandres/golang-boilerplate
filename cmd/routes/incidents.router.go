package routes

import (
	"lookerdevelopers/boilerplate/cmd/controllers"
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/instances"
	"lookerdevelopers/boilerplate/cmd/middlewares"
	"lookerdevelopers/boilerplate/cmd/services"

	"github.com/gin-gonic/gin"
)

func IncidentsRouter(router gin.IRouter) {

	incidentsRouter := router.Group("/incidents")

	iIncidentsService := services.NewIncidentsService(instances.DB)
	incidentsController := controllers.NewIncidentsController(iIncidentsService)

	incidentsRouter.POST("", middlewares.ValidateJSON[dto.IncidentPatchDto](), incidentsController.Patch)
}
