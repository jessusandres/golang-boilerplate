package tasks

import (
	"lookerdevelopers/boilerplate/cmd/dto"
	llog "lookerdevelopers/boilerplate/cmd/logger"
	"lookerdevelopers/boilerplate/cmd/models"

	"gorm.io/gorm"
)

type Result struct {
	RowsAffected int
	Error        error
	Incident     dto.IncidentDto
}

func SaveIncident(db *gorm.DB, payload *dto.IncidentPatchDto) Result {
	llog.Logger.Infof("Payload to save: %+v", payload)

	var newIncident dto.IncidentDto

	newRecord := models.Incident{
		Title:        payload.Title,
		Description:  payload.Description,
		IncidentType: payload.IncidentType,
		Location:     payload.Location,
		Image:        payload.Image,
		EventDate:    payload.EventDate,
	}

	result := db.Create(&newRecord)

	if result.Error != nil {
		return Result{
			Error: result.Error,
		}
	}

	newIncident = dto.IncidentDto{
		ID:           newRecord.ID,
		Title:        newRecord.Title,
		Description:  newRecord.Description,
		IncidentType: newRecord.IncidentType,
		Location:     newRecord.Location,
		Image:        newRecord.Image,
		EventDate:    newRecord.EventDate,
		CreatedAt:    newRecord.CreatedAt,
	}

	llog.Logger.Infof("Result: %+v", result)

	return Result{
		RowsAffected: int(result.RowsAffected),
		Error:        nil,
		Incident:     newIncident,
	}
}
