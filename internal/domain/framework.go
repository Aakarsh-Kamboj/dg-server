package domain

import (
	"time"

	"github.com/google/uuid"
)

type Framework struct {
	ID                     uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FrameworkName          string       `gorm:"size:255;not null" json:"framework_name"`
	NumberOfPolicies       int          `json:"number_of_policies"`
	NumberOfEvidenceTasks  int          `json:"number_of_evidence_tasks"`
	NumberOfAutomatedTests int          `json:"number_of_automated_tests"`
	IsCustom               bool         `json:"is_custom"`
	OrganizationID         uuid.UUID    `gorm:"type:uuid;not null;index" json:"organization_id"`
	Organization           Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	Controls               []Control    `gorm:"many2many:framework_controls;" json:"controls"` // ✅ added for many-to-many
	EvidenceTaskPercentage float64      `gorm:"-" json:"evidence_task_percentage"`             // <- calculated at runtime
	CreatedAt              time.Time    `gorm:"autoCreateTime" json:"created_at"`              // added timestamps
	UpdatedAt              time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
