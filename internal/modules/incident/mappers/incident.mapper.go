package mappers

import (
	"lookerdevelopers/boilerplate/internal/modules/incident/adapters/models"
	"lookerdevelopers/boilerplate/internal/modules/incident/domain"
)

type IncidentMapper struct{}

func (m *IncidentMapper) ToDomain(model models.Incident) domain.Incident {
	return domain.Incident{
		ID:           model.ID,
		Title:        model.Title,
		Description:  model.Description,
		IncidentType: model.IncidentType,
		Location:     model.Location,
		Image:        model.Image,
		EventDate:    model.EventDate,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

func (m *IncidentMapper) ToModel(incident domain.Incident) models.Incident {
	return models.Incident{
		ID:           incident.ID,
		Title:        incident.Title,
		Description:  incident.Description,
		IncidentType: incident.IncidentType,
		Location:     incident.Location,
		Image:        incident.Image,
		EventDate:    incident.EventDate,
		CreatedAt:    incident.CreatedAt,
		UpdatedAt:    incident.UpdatedAt,
	}
}
