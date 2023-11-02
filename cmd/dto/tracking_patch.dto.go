package dto

type TrackingPatchDto struct {
	ProviderIntegrationID string   `json:"provider_integration_id" binding:"required,uuid4" validate:"required,uuid4"`
	Payload               *Payload `json:"payload" binding:"required" validate:"required"`
	CreatedAt             string   `json:"createdAt" binding:"required" validate:"required"`
	UpdatedAt             string   `json:"updatedAt" binding:"required" validate:"required"`
}

type Payload struct {
	Shipment *Shipment `json:"shipment" binding:"required" validate:"required"`
}

type Shipment struct {
	ID                   string        `json:"id" binding:"required,uuid4" validate:"required,uuid4"`
	CreatedDateTime      string        `json:"createdDateTime" binding:"required" validate:"required"`
	Identifiers          []*Identifier `json:"identifiers" binding:"required,dive" validate:"required,dive"`
	ShipmentShareLink    string        `json:"shipmentShareLink" binding:"required" validate:"required"`
	LastModifiedDateTime string        `json:"lastModifiedDateTime" binding:"required" validate:"required"`
}

type Identifier struct {
	Type  string `json:"type" binding:"required" validate:"required"`
	Value string `json:"value" binding:"required" validate:"required"`
}
