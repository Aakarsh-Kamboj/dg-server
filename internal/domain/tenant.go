package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TenantStatus int

const (
	TenantStatusUnspecified TenantStatus = iota
	TenantStatusActive
	TenantStatusSuspended
	TenantStatusDeleted
)

type TenantType int

const (
	TenantTypeUnspecified TenantType = iota
	TenantTypeOrganization
	TenantTypeUser
)

type CompanySize int

const (
	CompanySizeUnspecified CompanySize = iota
	CompanySizeSmall
	CompanySizeMedium
	CompanySizeLarge
)

type Tenant struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TenantName      string         `gorm:"uniqueIndex;not null"                          json:"tenant_name"`
	TenantType      TenantType     `gorm:"not null;default:0"                            json:"tenant_type"`
	Address         TenantAddress  `gorm:"embedded;embeddedPrefix:address_"              json:"address"`
	Phone           *string        `json:"phone,omitempty"`
	Website         *string        `gorm:"size:100"                                      json:"website,omitempty"`
	Email           *string        `gorm:"uniqueIndex"                                   json:"email,omitempty"`
	Status          TenantStatus   `gorm:"default:0"                                     json:"status"`
	ParentID        *uuid.UUID     `gorm:"type:uuid;index"                                     json:"parent_id,omitempty"`
	CreatedByUserID uuid.UUID      `gorm:"type:uuid;"                            json:"created_by_user_id"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"                                json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"                                json:"updated_at"`
	IsDeleted       gorm.DeletedAt `json:"-"`
	Industry        *string        `json:"industry,omitempty" gorm:"size:255"`
	CompanySize     CompanySize    `gorm:"default:0"                                     json:"company_size"`
	Version         int            `gorm:"default:0"                                     json:"version"`
	Organizations   []Organization `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE" json:"organizations,omitempty"`
	Users           []User         `gorm:"foreignKey:TenantID;constraint:OnDelete:CASCADE" json:"users,omitempty"`
	Parent          *Tenant        `gorm:"foreignKey:ParentID"                            json:"parent,omitempty"`
	Children        []Tenant       `gorm:"foreignKey:ParentID"                            json:"children,omitempty"`
}
