package dto

import "time"

type SingleIncidentResult struct {
	ID           int
	Title        string
	Description  string
	IncidentType string
	Location     string
	Image        string
	EventDate    time.Time
	CreatedAt    time.Time
}

type IncidentResults []SingleIncidentResult

type FindIncidentsResult struct {
	Incidents IncidentResults
	Total     int
}
