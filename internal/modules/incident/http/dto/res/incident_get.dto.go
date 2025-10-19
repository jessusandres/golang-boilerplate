package httpdtores

import "time"

type IncidentDTO struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	IncidentType string    `json:"incidentType"`
	Location     string    `json:"location"`
	Image        string    `json:"image"`
	EventDate    time.Time `json:"eventDate"`
	CreatedAt    time.Time `json:"createdAt"`
}

type IncidentsDTO []IncidentDTO

type IncidentResponseDTO struct {
	IncidentsDTO IncidentsDTO `json:"incidents"`
	Total        int          `json:"total"`
}
