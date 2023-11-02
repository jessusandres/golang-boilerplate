package controllers

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/dto"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/interfaces"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/middlewares"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/utils"
	"github.com/gin-gonic/gin"
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
