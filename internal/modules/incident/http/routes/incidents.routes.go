package incidentsroutes

import (
	httpdtoreq "lookerdevelopers/boilerplate/internal/modules/incident/http/dto/req"
	incidentsinterfaces "lookerdevelopers/boilerplate/internal/modules/incident/interfaces"
	"lookerdevelopers/boilerplate/internal/shared/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRouter, controller incidentsinterfaces.IIncidentsController) {

	incidentsRouter := router.Group("/incidents")
	{
		// TODO: Add middleware to validate query params
		incidentsRouter.GET("", middlewares.ValidateQuery[httpdtoreq.HttpFindIncidentsDto](), controller.Get)
		incidentsRouter.POST("", middlewares.ValidateJSON[httpdtoreq.HTTPCreateIncidentDTO](), controller.Patch)
	}
}
