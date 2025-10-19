package types

import (
	"github.com/gin-gonic/gin"
)

// AppState represents the state of an application with a unique identifier.
type AppState struct {
	Uuid string
}

// ApiResult represents a standardized format for API responses containing a data payload.
type ApiResult struct {
	Data any `json:"data"`
}

func (ar *ApiResult) Response(c *gin.Context, statusCode int) {
	c.JSON(statusCode, ar)
}
