package incidentsinterfaces

import (
	"context"
	httpdtoreq "lookerdevelopers/boilerplate/internal/modules/incident/http/dto/req"
	httpdtores "lookerdevelopers/boilerplate/internal/modules/incident/http/dto/res"
)

type IIncidentsService interface {
	CreateIncident(ctx context.Context, payload *httpdtoreq.HTTPCreateIncidentDTO) (httpdtores.IncidentDTO, error)
	GetIncidents(ctx context.Context, payload httpdtoreq.HttpFindIncidentsDto) (httpdtores.IncidentResponseDTO, error)
}
