package repository

import (
	"context"
	"lookerdevelopers/boilerplate/internal/modules/incident/domain"
)

type IncidentFilters struct {
	Description string
	Limit       int
	Offset      int
}

type IncidentRepository interface {
	CreateIncident(ctx context.Context, payload domain.Incident) (domain.Incident, error)
	RetrieveIncidents(ctx context.Context, payload IncidentFilters) ([]domain.Incident, int, error)
	RetrieveIncident(ctx context.Context, id int) (domain.Incident, error)
	DeleteIncident(ctx context.Context, id int) error
	UpdateIncident(ctx context.Context, payload domain.Incident) (domain.Incident, error)
}
