package apperrors

import (
	"net/http"

	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/types"
)

func NewBadRequestError(message string) *types.HTTPError {
	return &types.HTTPError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) *types.HTTPError {
	return &types.HTTPError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *types.HTTPError {
	return &types.HTTPError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewServiceUnavailableError(message string) *types.HTTPError {
	return &types.HTTPError{
		Code:    http.StatusServiceUnavailable,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *types.HTTPError {
	return &types.HTTPError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewUnprocessableEntityError(message string) *types.HTTPError {
	return &types.HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
