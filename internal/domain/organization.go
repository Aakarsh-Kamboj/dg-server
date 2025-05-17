package domain

import (
	"time"

	"github.com/google/uuid"
)

// Organization holds metadata for tenants of type “organization.”
type Organization struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TenantID     uuid.UUID   `gorm:"type:uuid;not null;index"           json:"tenant_id"`
	Name         string      `gorm:"not null"                                  json:"name"`
	Address      string      `json:"address"`
	GSTIN        string      `json:"gstin" gorm:"uniqueIndex"`
	ContactEmail string      `json:"contact_email"`
	Frameworks   []Framework `gorm:"foreignKey:OrganizationID" json:"frameworks,omitempty"`
	CreatedAt    time.Time   `gorm:"autoCreateTime"                            json:"created_at"`
	UpdatedAt    time.Time   `gorm:"autoUpdateTime"                            json:"updated_at"`

	// back‐ref to Tenant
	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}
