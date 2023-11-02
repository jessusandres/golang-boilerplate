package middlewares

import (
	"lookerdevelopers/boilerplate/cmd/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BuildState() gin.HandlerFunc {
	return func(c *gin.Context) {
		newUuid, _ := uuid.NewUUID()
		reqUuid := newUuid.String()

		appState := types.AppState{
			Uuid: reqUuid,
		}

		c.Set("state", appState)

		c.Writer.Header().Set("X-Trace-Id", reqUuid)

		c.Next()
	}
}
