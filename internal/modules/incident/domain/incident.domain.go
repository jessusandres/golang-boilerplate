package domain

import "time"

type Incident struct {
	ID           int
	Title        string
	Description  string
	IncidentType string
	Location     string
	Image        string
	EventDate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
