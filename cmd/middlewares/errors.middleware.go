package middlewares

import (
	"errors"

	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"lookerdevelopers/boilerplate/cmd/apperrors"
	"lookerdevelopers/boilerplate/cmd/config"
	"lookerdevelopers/boilerplate/cmd/types"
	"lookerdevelopers/boilerplate/cmd/utils"
)

func HandleErr() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			state, _ := utils.ExtractState(context)
			reqUuid := state.Uuid

			err := context.Errors.Last()

			config.Logger.Info("Catching err for request: ", reqUuid)
			config.Logger.Error(err)

			var httpErr *types.HTTPError

			if errors.As(err.Err, &httpErr) {
				apiError := apperrors.ApiError{
					Message: httpErr.Message,
					UUID:    reqUuid,
				}

				apiError.ToResponse(context, httpErr.Code)

				return
			}

			var pgError *pgconn.PgError
			var pqError *pq.Error

			if errors.As(err, &pgError) || errors.As(err, &pqError) {
				config.Logger.Errorf("Database error for request %s, %v", reqUuid, err)

				dbErr := apperrors.NewInternalServerError("Oops, something went wrong. Please try again later.")

				apiError := apperrors.ApiError{
					Message: dbErr.Message,
					UUID:    reqUuid,
				}

				apiError.ToResponse(context, dbErr.Code)

				return
			}

			config.Logger.Info("Caution, unhandled error")
			unhandledErr := apperrors.NewInternalServerError("Oops, something went wrong when processing your request. Please try again later.")

			apiError := apperrors.ApiError{
				Message: unhandledErr.Message,
				UUID:    reqUuid,
			}

			apiError.ToResponse(context, unhandledErr.Code)
		}
	}

}
