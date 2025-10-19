package services

import (
	"context"
	"fmt"
	commanddto "lookerdevelopers/boilerplate/internal/modules/incident/commands/dto"
	commands "lookerdevelopers/boilerplate/internal/modules/incident/commands/impl"
	httpdtoreq "lookerdevelopers/boilerplate/internal/modules/incident/http/dto/req"
	httpdtores "lookerdevelopers/boilerplate/internal/modules/incident/http/dto/res"
	incidentsinterfaces "lookerdevelopers/boilerplate/internal/modules/incident/interfaces"
	"lookerdevelopers/boilerplate/internal/modules/incident/mappers"
	"lookerdevelopers/boilerplate/internal/modules/incident/queries/dto"
	queries "lookerdevelopers/boilerplate/internal/modules/incident/queries/impl"
	"lookerdevelopers/boilerplate/internal/shared/cqrs"
	cqrsinterfaces "lookerdevelopers/boilerplate/internal/shared/cqrs/interfaces"
	"lookerdevelopers/boilerplate/internal/shared/utils"
)

type IncidentsService struct {
	commandBus cqrsinterfaces.ICommandBus
	queryBus   cqrsinterfaces.IQueryBus
	mapper     *mappers.IncidentMapper
}

func NewIncidentsService(
	commandBus cqrsinterfaces.ICommandBus,
	queryBus cqrsinterfaces.IQueryBus,
) incidentsinterfaces.IIncidentsService {
	return &IncidentsService{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

func (svc *IncidentsService) CreateIncident(ctx context.Context, payload *httpdtoreq.HTTPCreateIncidentDTO) (httpdtores.IncidentDTO, error) {
	var incidentDto httpdtores.IncidentDTO

	command := commands.CreateIncidentCommand{
		Title:        payload.Title,
		Description:  payload.Description,
		IncidentType: payload.IncidentType,
		Location:     payload.Location,
		EventDate:    payload.EventDate,
	}

	result, err := cqrs.ExecuteCommand[commanddto.CUIncidentResult](ctx, svc.commandBus, command)

	if err != nil {
		return incidentDto, err
	}

	incidentDto = httpdtores.IncidentDTO{
		ID:           result.ID,
		Title:        result.Title,
		Description:  result.Description,
		IncidentType: result.IncidentType,
		Location:     result.Location,
		Image:        result.Image,
		EventDate:    result.EventDate,
		CreatedAt:    result.CreatedAt,
	}

	return incidentDto, nil
}

func (svc *IncidentsService) GetIncidents(ctx context.Context, payload httpdtoreq.HttpFindIncidentsDto) (httpdtores.IncidentResponseDTO, error) {
	var incidentsResponseDTO httpdtores.IncidentResponseDTO

	query := queries.FindIncidentsQuery{
		Description: payload.Description,
		Limit:       payload.Limit,
		Offset:      payload.Offset,
	}

	queryResult, err := cqrs.ExecuteQuery[dto.FindIncidentsResult](ctx, svc.queryBus, query)

	if err != nil {
		return incidentsResponseDTO, err
	}

	fmt.Printf("Incidents from query: %v", queryResult)

	incidentsDTO := utils.MapSlice(queryResult.Incidents, func(incident dto.SingleIncidentResult) httpdtores.IncidentDTO {
		return httpdtores.IncidentDTO{
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

	incidentsResponseDTO = httpdtores.IncidentResponseDTO{
		IncidentsDTO: incidentsDTO,
		Total:        queryResult.Total,
	}

	return incidentsResponseDTO, nil
}
