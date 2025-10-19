package impl

import (
	"context"
	"lookerdevelopers/boilerplate/internal/modules/incident/adapters/models"
	"lookerdevelopers/boilerplate/internal/modules/incident/domain"
	"lookerdevelopers/boilerplate/internal/modules/incident/infrastructure/repository"
	"lookerdevelopers/boilerplate/internal/modules/incident/mappers"
	"lookerdevelopers/boilerplate/internal/shared/slog"

	"gorm.io/gorm"
)

type GormIncidentImpl struct {
	db     *gorm.DB
	mapper *mappers.IncidentMapper
}

func NewGormIncidentImpl(db *gorm.DB, mapper *mappers.IncidentMapper) repository.IncidentRepository {
	return &GormIncidentImpl{
		db:     db,
		mapper: mapper,
	}
}

func (impl *GormIncidentImpl) CreateIncident(ctx context.Context, payload domain.Incident) (domain.Incident, error) {
	slog.Logger.Infof("CreateIncident in GORM")

	var incidentDomain domain.Incident

	newRecord := models.Incident{
		Title:        payload.Title,
		Description:  payload.Description,
		IncidentType: payload.IncidentType,
		Location:     payload.Location,
		Image:        payload.Image,
		EventDate:    payload.EventDate,
	}

	result := impl.db.WithContext(ctx).Create(&newRecord)

	if result.Error != nil {
		return incidentDomain, result.Error
	}

	return impl.mapper.ToDomain(newRecord), nil
}

func (impl *GormIncidentImpl) RetrieveIncidents(ctx context.Context, payload repository.IncidentFilters) ([]domain.Incident, int, error) {
	slog.Logger.Infof("RetrieveIncidents in GORM")

	var incidents []domain.Incident
	var incidentRows []models.Incident
	var total int64

	query := impl.db.WithContext(ctx)

	if payload.Description != "" {
		query = query.Where("description LIKE ?", "%"+payload.Description+"%")
	}

	if err := query.Model(&models.Incident{}).Count(&total).Error; err != nil {
		return incidents, int(total), err
	}

	result := query.Limit(payload.Limit).Offset(payload.Offset).Find(&incidentRows)

	if result.Error != nil {
		return incidents, int(total), result.Error
	}

	for _, row := range incidentRows {
		incident := impl.mapper.ToDomain(row)
		incidents = append(incidents, incident)
	}

	return incidents, int(total), nil
}

func (impl *GormIncidentImpl) RetrieveIncident(ctx context.Context, id int) (domain.Incident, error) {
	var incident domain.Incident

	return incident, nil
}

func (impl *GormIncidentImpl) DeleteIncident(ctx context.Context, id int) error {
	return nil
}

func (impl *GormIncidentImpl) UpdateIncident(ctx context.Context, payload domain.Incident) (domain.Incident, error) {
	var updatedIncident domain.Incident

	return updatedIncident, nil
}
