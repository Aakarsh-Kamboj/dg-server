package domain

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID             uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"` // change string to uuid.UUID
	DepartmentName string       `gorm:"size:255;not null" json:"department_name"`
	OrganizationID uuid.UUID    `gorm:"type:uuid;not null;index" json:"organization_id"`                              // change string to uuid.UUID, add constraints
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"-"` // add constraint
	CreatedAt      time.Time    `gorm:"autoCreateTime" json:"created_at"`                                             // add timestamps
	UpdatedAt      time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
