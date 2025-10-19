package commandhandlers

import (
	"context"
	commanddto "lookerdevelopers/boilerplate/internal/modules/incident/commands/dto"
	commandimpl "lookerdevelopers/boilerplate/internal/modules/incident/commands/impl"
	"lookerdevelopers/boilerplate/internal/modules/incident/infrastructure/repository"
)

type UpdateIncidentHandler struct {
	incidentRepository repository.IncidentRepository
}

func NewUpdateIncidentHandler(incidentRepository repository.IncidentRepository) *UpdateIncidentHandler {
	return &UpdateIncidentHandler{
		incidentRepository: incidentRepository,
	}
}

func (h *UpdateIncidentHandler) Handle(ctx context.Context, payload commandimpl.UpdateIncidentCommand) (*commanddto.CUIncidentResult, error) {
	return nil, nil
}
