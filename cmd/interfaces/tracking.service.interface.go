package interfaces

import (
	"context"
	"lookerdevelopers/boilerplate/cmd/dto"
)

type ITrackingService interface {
	Patch(ctx context.Context, payload *dto.TrackingPatchDto) (int, error)
}
