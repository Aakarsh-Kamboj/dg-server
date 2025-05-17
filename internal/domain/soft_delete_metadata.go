package domain

import (
	"github.com/google/uuid"
	"time"
)

// UserSoftDeleteMetadata keeps track of deleted users
type UserSoftDeleteMetadata struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	DeletedAt time.Time `gorm:"autoCreateTime"`
	Reason    string    `gorm:"not null"` // Reason for deletion (optional)
}
