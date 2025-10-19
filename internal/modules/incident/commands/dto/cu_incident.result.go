package commanddto

import "time"

type CUIncidentResult struct {
	ID           int
	Title        string
	Description  string
	IncidentType string
	Location     string
	Image        string
	EventDate    time.Time
	CreatedAt    time.Time
}
