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

// ApiErrorResponse represents the structure of an API error response containing an error payload.
type ApiErrorResponse struct {
	Error ApiError `json:"error"`
}

func (e *ApiError) ToResponse(c *gin.Context, statusCode int) {
	c.AbortWithStatusJSON(statusCode, ApiErrorResponse{
		Error: ApiError{
			Message: e.Message,
			UUID:    e.UUID,
			Details: e.Details,
		},
	})
}
