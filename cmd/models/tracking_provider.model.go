package models

import "gorm.io/datatypes"
import "time"

type TrackingProvider struct {
	ProviderIntegrationID string         `gorm:"column:provider_integration_id;primaryKey"`
	Payload               datatypes.JSON `gorm:"column:payload;type:jsonb"`
	CreatedAt             time.Time      `gorm:"column:created_at;type:timestamp"`
	UpdatedAt             time.Time      `gorm:"column:updated_at;type:timestamp"`
}

func (m *TrackingProvider) TableName() string {
	return "provider_integration"
}
