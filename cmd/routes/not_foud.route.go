package routes

import (
	"fmt"
	"lookerdevelopers/boilerplate/cmd/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
	"lookerdevelopers/boilerplate/cmd/utils"
)

func NotFound(c *gin.Context) {
	routePath := c.Request.URL.Path
	routeMethod := c.Request.Method

	message := fmt.Sprintf("Route [%s] %s not found", routeMethod, routePath)
	uuid := ""

	state, ok := utils.ExtractState(c)

	if ok {
		uuid = state.Uuid
	}

	c.JSON(http.StatusNotFound, apperrors.ApiErrorResponse{
		Error: apperrors.ApiError{
			Message: message,
			UUID:    uuid,
		},
	})
}
