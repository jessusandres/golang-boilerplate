package commands

import "time"

type CreateIncidentCommand struct {
	Title        string    `validate:"required,min=3,max=100"`
	Description  string    `validate:"required,min=10,max=1000"`
	IncidentType string    `validate:"required,oneof=emergency warning info"`
	Location     string    `validate:"required,min=3,max=200"`
	Image        string    `validate:"base64"`
	EventDate    time.Time `validate:"required"`
}

func (c CreateIncidentCommand) CommandName() string {
	return "create_incident"
}
