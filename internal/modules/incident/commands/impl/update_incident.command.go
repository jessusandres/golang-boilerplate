package commands

import "time"

type UpdateIncidentCommand struct {
	ID           int       `json:"id" validate:"required,min=1"`
	Title        string    `json:"title" validate:"required,min=3,max=100"`
	Description  string    `json:"description" validate:"required,min=10,max=1000"`
	IncidentType string    `json:"incident_type" validate:"required,oneof=emergency warning info"`
	Location     string    `json:"location" validate:"required,min=3,max=200"`
	EventDate    time.Time `json:"event_date" validate:"required"`
}

func (c UpdateIncidentCommand) CommandName() string {
	return "update_incident"
}
