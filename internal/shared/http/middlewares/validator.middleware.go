package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"lookerdevelopers/boilerplate/internal/interfaces"
	apierrors "lookerdevelopers/boilerplate/internal/shared/errors/api"
	apperrors "lookerdevelopers/boilerplate/internal/shared/errors/app"
	"lookerdevelopers/boilerplate/internal/shared/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateQuery[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {

		valid := isValidJsonRequest(c)

		if !valid {
			return
		}

		payload, isValid := validateQueryPayload[T](c)

		if !isValid {
			return
		}

		c.Set("payload", payload)

		c.Next()
	}
}

func validateQueryPayload[T any](c *gin.Context) (T, bool) {
	var payload T

	err := c.ShouldBindQuery(&payload)

	if err != nil {
		badRequestErr := apperrors.NewBadRequestError(err.Error())

		apiError := apierrors.ApiError{
			Message: badRequestErr.Message,
		}

		apiError.ToResponse(c, badRequestErr.Code)

		return payload, false
	}

	return payload, true
}

func ValidateJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {

		valid := isValidJsonRequest(c)

		if !valid {
			return
		}

		payload, validJSON := validateJsonPayload[T](c)

		if !validJSON {
			return
		}

		c.Set("payload", payload)

		c.Next()
	}
}

func isValidJsonRequest(c *gin.Context) bool {
	contentType := c.GetHeader("Content-Type")

	if contentType != "application/json" {
		badRequestErr := apperrors.NewBadRequestError("Content-Type must be application/json")

		apiError := apierrors.ApiError{
			Message: badRequestErr.Message,
		}

		apiError.ToResponse(c, badRequestErr.Code)
		return false
	}

	if c.Request.Body == nil {
		badRequestErr := apperrors.NewBadRequestError("Request body is empty")
		apiError := apierrors.ApiError{
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
			apiError := apierrors.ApiError{
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

			apiError := apierrors.ApiError{
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
			apiError := apierrors.ApiError{
				Message: unprocessableEntityErr.Message,
				Details: errDetails,
			}

			apiError.ToResponse(c, unprocessableEntityErr.Code)

			return payload, false
		}

		unprocessableEntityErr := apperrors.NewUnprocessableEntityError(err.Error())
		apiError := apierrors.ApiError{
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
	log.Printf("Tag: %s - Field: %s - Value: %s\n", fe.Tag(), fe.Field(), fe.Param())

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
	case "base64":
		return fmt.Sprintf("Should be a valid %s", fe.Tag())
	}

	return "Unknown error"
}

// GetValidatedPayload returns the validated payload from the context or an empty payload if it doesn't exist.
//
// Returns:
//   - The first return value is the validated payload.
//   - The second return value indicates if the payload exists or not.
//
// The error is filled in the context if the assertion fails.
func GetValidatedPayload[T any](c *gin.Context) (T, bool) {
	var zero T

	value, exists := c.Get("payload")

	if !exists {
		return zero, false
	}

	payload, ok := value.(T)

	if !ok {
		internalErr := apperrors.NewInternalServerError("Failed to get validated payload")

		apiError := apierrors.ApiError{
			Message: internalErr.Message,
		}

		apiError.ToResponse(c, internalErr.Code)
	}

	return payload, ok
}
