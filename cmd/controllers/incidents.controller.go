package controllers

import (
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/interfaces"
	"lookerdevelopers/boilerplate/cmd/middlewares"
	"lookerdevelopers/boilerplate/cmd/types"
	"lookerdevelopers/boilerplate/cmd/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type IncidentsController struct {
	IncidentsService interfaces.IIncidentsService
}

func NewIncidentsController(svc interfaces.IIncidentsService) *IncidentsController {
	return &IncidentsController{
		IncidentsService: svc,
	}
}

func (ctrl *IncidentsController) Patch(ctx *gin.Context) {
	payload, success := middlewares.GetValidatedPayload[dto.IncidentPatchDto](ctx)

	if !success {
		return
	}

	affectedRows, err := ctrl.IncidentsService.Patch(ctx, &payload)

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
