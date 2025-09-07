package interfaces

import (
	"context"
	"lookerdevelopers/boilerplate/cmd/dto"
)

type IIncidentsService interface {
	Patch(ctx context.Context, payload *dto.IncidentPatchDto) (dto.IncidentDto, error)
}
