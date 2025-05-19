package domain

import (
	"time"

	"github.com/google/uuid"
)

type Control struct {
	ID              uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ControlCode     string        `gorm:"size:50;not null;uniqueIndex" json:"control_code"`
	ControlName     string        `gorm:"size:255;not null" json:"control_name"`
	ControlDomain   string        `gorm:"size:255" json:"control_domain"`
	Status          ControlStatus `gorm:"size:20" json:"status"`
	Assignee        string        `gorm:"size:255" json:"assignee"`
	Description     string        `gorm:"type:text" json:"description,omitempty"`
	ControlQuestion string        `gorm:"type:text" json:"control_question,omitempty"`
	Frameworks      []Framework   `gorm:"many2many:framework_controls;" json:"frameworks"`                              // ✅ added for many-to-many
	OrganizationID  uuid.UUID     `gorm:"type:uuid;not null;index" json:"organization_id"`                              // changed from string to uuid.UUID
	Organization    Organization  `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"-"` // added constraint
	CreatedAt       time.Time     `gorm:"autoCreateTime" json:"created_at"`                                             // added timestamps
	UpdatedAt       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
