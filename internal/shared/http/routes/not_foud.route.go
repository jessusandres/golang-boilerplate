package routes

import (
	"fmt"
	apierrors "lookerdevelopers/boilerplate/internal/shared/errors/api"
	"net/http"

	"lookerdevelopers/boilerplate/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	routePath := c.Request.URL.Path
	routeMethod := c.Request.Method

	message := fmt.Sprintf("Route [%s] %s not found", routeMethod, routePath)
	uuid := ""

	state, ok := utils.ExtractAppState(c)

	if ok {
		uuid = state.Uuid
	}

	apiErr := apierrors.ApiError{
		Message: message,
		UUID:    uuid,
	}

	apiErr.ToResponse(c, http.StatusNotFound)
}
