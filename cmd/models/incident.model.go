package models

import (
	"time"
)

type Incident struct {
	ID           int       `gorm:"column:id;primary_key"`
	Title        string    `gorm:"column:title"`
	Description  string    `gorm:"column:description"`
	IncidentType string    `gorm:"column:incident_type"`
	Location     string    `gorm:"column:location"`
	Image        string    `gorm:"column:image"`
	EventDate    time.Time `gorm:"column:event_date"`

	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

func (m *Incident) TableName() string {
	return "incidents"
}
