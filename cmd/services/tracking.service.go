package services

import (
	"context"
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/interfaces"
	llog "lookerdevelopers/boilerplate/cmd/logger"
	"lookerdevelopers/boilerplate/cmd/tasks"

	"gorm.io/gorm"
)

type IncidentsService struct {
	DB *gorm.DB
}

func NewIncidentsService(DB *gorm.DB) interfaces.IIncidentsService {
	return &IncidentsService{
		DB: DB,
	}
}

func (svc *IncidentsService) Patch(_ context.Context, payload *dto.IncidentPatchDto) (dto.IncidentDto, error) {
	result := tasks.SaveIncident(svc.DB, payload)

	llog.Logger.Infof("Affected rows: %d", result.RowsAffected)
	llog.Logger.Infof("Error: %s", result.Error)
	llog.Logger.Infof("Incident: %+v", result.Incident)

	return result.Incident, result.Error
}
