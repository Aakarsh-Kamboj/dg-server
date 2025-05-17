package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusNotUploaded    = "Not Uploaded"
	StatusDraft          = "Draft"
	StatusNeedsAttention = "Needs Attention"
	StatusUploaded       = "Uploaded"
)

type EvidenceTask struct {
	ID             uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EvidenceName   string       `gorm:"size:255;not null" json:"evidence_name"`
	Status         string       `gorm:"size:20" json:"status"`
	Assignee       string       `gorm:"size:255" json:"assignee"`
	DepartmentID   uuid.UUID    `gorm:"type:uuid;index" json:"department_id"`
	Department     Department   `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnDelete:SET NULL" json:"-"`
	DueDate        time.Time    `json:"due_date"`
	UploadedDate   time.Time    `json:"uploaded_date"`
	FrameworkID    uuid.UUID    `gorm:"type:uuid;not null;index" json:"framework_id"`
	OrganizationID uuid.UUID    `gorm:"type:uuid;not null;index" json:"organization_id"`
	Framework      Framework    `gorm:"foreignKey:FrameworkID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"-"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
