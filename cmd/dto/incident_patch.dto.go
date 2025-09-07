package dto

import "time"

type IncidentPatchDto struct {
	Title        string    `json:"title" binding:"required" validate:"required"`
	Description  string    `json:"description" binding:"required" validate:"required"`
	IncidentType string    `json:"incidentType" binding:"required" validate:"required"`
	Location     string    `json:"location" binding:"required" validate:"required"`
	Image        string    `json:"image" binding:"base64" validate:"base64"`
	EventDate    time.Time `json:"eventDate" binding:"required" validate:"required,iso8601"`
}
