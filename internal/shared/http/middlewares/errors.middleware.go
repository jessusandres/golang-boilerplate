package middlewares

import (
	"errors"
	apierrors "lookerdevelopers/boilerplate/internal/shared/errors/api"
	apperrors "lookerdevelopers/boilerplate/internal/shared/errors/app"
	llog "lookerdevelopers/boilerplate/internal/shared/slog"

	"github.com/lib/pq"

	"lookerdevelopers/boilerplate/internal/shared/types"
	"lookerdevelopers/boilerplate/internal/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func HandleErr() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			state, _ := utils.ExtractAppState(context)
			reqUuid := state.Uuid

			err := context.Errors.Last()

			llog.Logger.Info("Catching err for request: ", reqUuid)
			llog.Logger.Error(err)

			var httpErr *types.HTTPError

			if errors.As(err.Err, &httpErr) {
				apiError := apierrors.ApiError{
					Message: httpErr.Message,
					UUID:    reqUuid,
				}

				apiError.ToResponse(context, httpErr.Code)

				return
			}

			var pgError *pgconn.PgError
			var pqError *pq.Error

			if errors.As(err, &pgError) || errors.As(err, &pqError) {
				llog.Logger.Errorf("Database error for request %s, %v", reqUuid, err)

				dbErr := apperrors.NewInternalServerError("Oops, something went wrong. Please try again later.")

				apiError := apierrors.ApiError{
					Message: dbErr.Message,
					UUID:    reqUuid,
				}

				apiError.ToResponse(context, dbErr.Code)

				return
			}

			llog.Logger.Info("Caution, unhandled error")
			unhandledErr := apperrors.NewInternalServerError("Oops, something went wrong when processing your request. Please try again later.")

			apiError := apierrors.ApiError{
				Message: unhandledErr.Message,
				UUID:    reqUuid,
			}

			apiError.ToResponse(context, unhandledErr.Code)
		}
	}

}
