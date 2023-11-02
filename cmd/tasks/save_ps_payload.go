package tasks

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/apperrors"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/dto"
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/models"
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"time"
)

func SavePSPayload(db *gorm.DB, payload *dto.TrackingPatchDto) (int, error) {
	log.Printf("Payload: %+v", payload)

	stringifyBytes, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling payload: %s", err)
		return 0, apperrors.NewInternalServerError("Error marshalling payload")
	}

	jsonPayload := datatypes.JSON(stringifyBytes)

	newRecord := models.TrackingProvider{
		ProviderIntegrationID: payload.ProviderIntegrationID,
		Payload:               jsonPayload,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	result := db.Create(&newRecord)

	if result.Error != nil {
		return 0, result.Error
	}

	log.Printf("Result: %+v", result)

	return int(result.RowsAffected), nil
}
