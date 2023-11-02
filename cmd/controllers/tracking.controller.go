package controllers

import (
	"github.com/gin-gonic/gin"
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/interfaces"
	"lookerdevelopers/boilerplate/cmd/middlewares"
	"lookerdevelopers/boilerplate/cmd/utils"
)

type TrackingController struct {
	TrackingService interfaces.ITrackingService
}

func NewTrackingController(svc interfaces.ITrackingService) *TrackingController {
	return &TrackingController{
		TrackingService: svc,
	}
}

func (ctrl *TrackingController) Patch(ctx *gin.Context) {
	payload, success := middlewares.GetValidatedPayload[dto.TrackingPatchDto](ctx)

	if !success {
		return
	}

	affectedRows, err := ctrl.TrackingService.Patch(ctx, &payload)

	if hasErr := utils.HandleServiceError(ctx, err); hasErr == true {
		return
	}

	ctx.JSON(200, gin.H{
		"affected": affectedRows,
	})
}
