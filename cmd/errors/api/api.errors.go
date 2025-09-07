package apperrors

import (
	"github.com/gin-gonic/gin"
)

// ApiError represents the structure for error payloads in API responses.
// It includes a message, a unique identifier (UUID), and optionally additional error details used for payload errors.
type ApiError struct {
	Message string `json:"message"`
	Details any    `json:"payload,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

func (e *ApiError) ToResponse(c *gin.Context, statusCode int) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": ApiError{
			Message: e.Message,
			UUID:    e.UUID,
			Details: e.Details,
		},
	})
}
