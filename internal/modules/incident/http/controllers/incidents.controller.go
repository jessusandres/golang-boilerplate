package controllers

import (
	"fmt"
	httpdtoreq "lookerdevelopers/boilerplate/internal/modules/incident/http/dto/req"
	incidentsinterfaces "lookerdevelopers/boilerplate/internal/modules/incident/interfaces"
	"lookerdevelopers/boilerplate/internal/shared/http/middlewares"
	"lookerdevelopers/boilerplate/internal/shared/types"
	"lookerdevelopers/boilerplate/internal/shared/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type IncidentsController struct {
	IncidentsService incidentsinterfaces.IIncidentsService
}

func NewIncidentsController(svc incidentsinterfaces.IIncidentsService) incidentsinterfaces.IIncidentsController {
	return &IncidentsController{
		IncidentsService: svc,
	}
}

func (ctrl *IncidentsController) Patch(ctx *gin.Context) {
	payload, success := middlewares.GetValidatedPayload[httpdtoreq.HTTPCreateIncidentDTO](ctx)

	if !success {
		return
	}

	affectedRows, err := ctrl.IncidentsService.CreateIncident(ctx, &payload)

	if aborted := utils.GinAbortError(ctx, err); aborted == true {
		return
	}

	apiResponse := types.ApiResult{
		Data: gin.H{
			"incident": affectedRows,
		},
	}

	apiResponse.Response(ctx, http.StatusOK)
}

func (ctrl *IncidentsController) Get(ctx *gin.Context) {

	payload, success := middlewares.GetValidatedPayload[httpdtoreq.HttpFindIncidentsDto](ctx)

	if !success {
		return
	}

	fmt.Printf("Query: %v\n", payload)

	incidents, err := ctrl.IncidentsService.GetIncidents(ctx, payload)

	if aborted := utils.GinAbortError(ctx, err); aborted == true {
		return
	}

	apiResponse := types.ApiResult{
		Data: incidents,
	}

	apiResponse.Response(ctx, http.StatusOK)
}
