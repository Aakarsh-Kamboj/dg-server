package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User represents a login within a tenant’s scope.
type User struct {
	ID           uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email        string         `gorm:"uniqueIndex;not null"                          json:"email"`
	Phone        *string        `gorm:"uniqueIndex"                                   json:"phone,omitempty"`
	PasswordHash string         `gorm:"not null"                                       json:"-"`
	TenantID     uuid.UUID      `gorm:"index;type:uuid;not null" json:"tenant_id"`
	Active       bool           `gorm:"default:true"                                   json:"active"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"                                json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"                                json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"`

	// FK → tenants.id Relation back to tenant
	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID;constraint:OnDelete:CASCADE" json:"-"`

	// Profile (one‐to‐one)
	Profile *UserProfile `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"     json:"profile,omitempty"`
}
