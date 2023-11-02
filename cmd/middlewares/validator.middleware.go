package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"lookerdevelopers/boilerplate/cmd/apperrors"
	"lookerdevelopers/boilerplate/cmd/interfaces"
	"lookerdevelopers/boilerplate/cmd/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {

		isValid := validateRequest(c)

		if !isValid {
			return
		}

		payload, isValid := validateJsonPayload[T](c)

		if !isValid {
			return
		}

		c.Set("payload", payload)

		c.Next()
	}
}

func validateRequest(c *gin.Context) bool {
	contentType := c.GetHeader("Content-Type")

	if contentType != "application/json" {
		badRequestErr := apperrors.NewBadRequestError("Content-Type must be application/json")
		apiError := apperrors.ApiError{
			Message: badRequestErr.Message,
		}

		apiError.ToResponse(c, badRequestErr.Code)
		return false
	}

	if c.Request.Body == nil {
		badRequestErr := apperrors.NewBadRequestError("Request body is empty")
		apiError := apperrors.ApiError{
			Message: badRequestErr.Message,
		}

		apiError.ToResponse(c, badRequestErr.Code)
		return false
	}

	return true
}

func validateJsonPayload[T any](c *gin.Context) (T, bool) {
	var payload T

	if err := c.ShouldBindJSON(&payload); err != nil {

		if errors.Is(err, io.EOF) {
			badRequestErr := apperrors.NewBadRequestError("Request body is empty")
			apiError := apperrors.ApiError{
				Message: badRequestErr.Message,
			}

			apiError.ToResponse(c, badRequestErr.Code)

			return payload, false
		}

		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			badRequestErr := apperrors.NewBadRequestError(
				fmt.Sprintf(
					"Invalid type for field '%s': expected %s, got %s",
					unmarshalTypeError.Field,
					unmarshalTypeError.Type,
					unmarshalTypeError.Value,
				),
			)

			apiError := apperrors.ApiError{
				Message: badRequestErr.Message,
			}

			apiError.ToResponse(c, badRequestErr.Code)

			return payload, false
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errDetails := make([]types.ValidationError, len(ve))

			for i, fe := range ve {
				errDetails[i] = types.ValidationError{Field: fe.Field(), Message: ValidationErrorMsg(fe)}
			}

			unprocessableEntityErr := apperrors.NewUnprocessableEntityError("Errors found in request body")
			apiError := apperrors.ApiError{
				Message: unprocessableEntityErr.Message,
				Details: errDetails,
			}

			apiError.ToResponse(c, unprocessableEntityErr.Code)

			return payload, false
		}

		unprocessableEntityErr := apperrors.NewUnprocessableEntityError(err.Error())
		apiError := apperrors.ApiError{
			Message: unprocessableEntityErr.Message,
		}

		apiError.ToResponse(c, unprocessableEntityErr.Code)

		return payload, false
	}

	if schema, ok := interface{}(&payload).(interfaces.Validator); ok {
		schema.SetDefaults()
	}

	return payload, true
}

func ValidationErrorMsg(fe validator.FieldError) string {
	log.Printf("Tag: %s, Param: %s\n", fe.Tag(), fe.Param())

	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Should be a valid email"
	case "min":
		return fmt.Sprintf("Should be at least %s characters long", fe.Param())
	case "len":
		return fmt.Sprintf("Should be exactly %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Should be at most %s characters long", fe.Param())
	case "gte":
		return fmt.Sprintf("Should be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Should be less than or equal to %s", fe.Param())
	case "uuid4":
		return fmt.Sprintf("Should be a valid %s", fe.Tag())
	}

	return "Unknown error"
}

func GetValidatedPayload[T any](c *gin.Context) (T, bool) {
	var zero T

	value, exists := c.Get("payload")

	if !exists {
		return zero, false
	}

	payload, ok := value.(T)

	if !ok {
		internalErr := apperrors.NewInternalServerError("Failed to get validated payload")
		apiError := apperrors.ApiError{
			Message: internalErr.Message,
		}

		apiError.ToResponse(c, internalErr.Code)
	}

	return payload, ok
}
