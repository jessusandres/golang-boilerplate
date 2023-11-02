package services

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/dto"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/interfaces"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/tasks"
	"context"
	"gorm.io/gorm"
)

type TrackingService struct {
	DB *gorm.DB
}

func NewTrackingService(DB *gorm.DB) interfaces.ITrackingService {
	return &TrackingService{
		DB: DB,
	}
}

func (svc *TrackingService) Patch(_ context.Context, payload *dto.TrackingPatchDto) (int, error) {
	return tasks.SavePSPayload(svc.DB, payload)
}
