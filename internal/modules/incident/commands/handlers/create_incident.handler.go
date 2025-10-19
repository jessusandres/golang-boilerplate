package commandhandlers

import (
	"context"
	commanddto "lookerdevelopers/boilerplate/internal/modules/incident/commands/dto"
	commands "lookerdevelopers/boilerplate/internal/modules/incident/commands/impl"
	"lookerdevelopers/boilerplate/internal/modules/incident/domain"
	"lookerdevelopers/boilerplate/internal/modules/incident/infrastructure/repository"
	cqrsinterfaces "lookerdevelopers/boilerplate/internal/shared/cqrs/interfaces"
	"lookerdevelopers/boilerplate/internal/shared/slog"
)

type CreateIncidentHandler struct {
	incidentRepository repository.IncidentRepository
}

func NewCreateIncidentHandler(incidentRepository repository.IncidentRepository) cqrsinterfaces.ICommandHandler[commands.CreateIncidentCommand, commanddto.CUIncidentResult] {
	return &CreateIncidentHandler{
		incidentRepository: incidentRepository,
	}
}

func (h *CreateIncidentHandler) Handle(ctx context.Context, payload commands.CreateIncidentCommand) (commanddto.CUIncidentResult, error) {
	slog.Logger.Debug("Executing CreateIncidentHandler for commands:", payload.CommandName())

	var result commanddto.CUIncidentResult

	domainIncident := domain.Incident{
		Title:        payload.Title,
		Description:  payload.Description,
		IncidentType: payload.IncidentType,
		Location:     payload.Location,
		Image:        payload.Image,
		EventDate:    payload.EventDate,
	}

	createdIncident, err := h.incidentRepository.CreateIncident(ctx, domainIncident)

	if err != nil {
		return result, err
	}

	result = commanddto.CUIncidentResult{
		ID:           createdIncident.ID,
		Title:        createdIncident.Title,
		Description:  createdIncident.Description,
		IncidentType: createdIncident.IncidentType,
		Location:     createdIncident.Location,
		Image:        createdIncident.Image,
		EventDate:    createdIncident.EventDate,
		CreatedAt:    createdIncident.CreatedAt,
	}

	return result, nil
}
