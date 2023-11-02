package services

import (
	"context"
	"gorm.io/gorm"
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/interfaces"
	"lookerdevelopers/boilerplate/cmd/tasks"
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
