package interfaces

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/dto"
	"context"
)

type ITrackingService interface {
	Patch(ctx context.Context, payload *dto.TrackingPatchDto) (int, error)
}
