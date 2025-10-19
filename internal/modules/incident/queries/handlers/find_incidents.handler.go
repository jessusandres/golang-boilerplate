package queryhandlers

import (
	"context"
	"fmt"
	"lookerdevelopers/boilerplate/internal/modules/incident/domain"
	"lookerdevelopers/boilerplate/internal/modules/incident/infrastructure/repository"
	"lookerdevelopers/boilerplate/internal/modules/incident/queries/dto"
	commands "lookerdevelopers/boilerplate/internal/modules/incident/queries/impl"
	cqrsinterfaces "lookerdevelopers/boilerplate/internal/shared/cqrs/interfaces"
	"lookerdevelopers/boilerplate/internal/shared/slog"
	"lookerdevelopers/boilerplate/internal/shared/utils"
)

type FindIncidentsHandler struct {
	incidentRepository repository.IncidentRepository
}

func NewFindIncidentsHandler(incidentRepository repository.IncidentRepository) cqrsinterfaces.IQueryHandler[commands.FindIncidentsQuery, dto.FindIncidentsResult] {
	return &FindIncidentsHandler{
		incidentRepository: incidentRepository,
	}
}

func (h *FindIncidentsHandler) Handle(ctx context.Context, payload commands.FindIncidentsQuery) (dto.FindIncidentsResult, error) {
	slog.Logger.Debugf("FindIncidentsHandler - payload: %v", payload)
	var result dto.FindIncidentsResult

	fmt.Printf("FindIncidentsHandler - payload: %v", payload)

	filters := repository.IncidentFilters{
		Description: payload.Description,
		Limit:       payload.Limit,
		Offset:      payload.Offset,
	}

	incidentDomains, total, err := h.incidentRepository.RetrieveIncidents(ctx, filters)

	if err != nil {
		return result, err
	}

	incidents := utils.MapSlice(incidentDomains, func(incident domain.Incident) dto.SingleIncidentResult {
		return dto.SingleIncidentResult{
			ID:           incident.ID,
			Title:        incident.Title,
			Description:  incident.Description,
			IncidentType: incident.IncidentType,
			Location:     incident.Location,
			Image:        incident.Image,
			EventDate:    incident.EventDate,
			CreatedAt:    incident.CreatedAt,
		}
	})

	result = dto.FindIncidentsResult{
		Incidents: incidents,
		Total:     total,
	}

	return result, nil
}
