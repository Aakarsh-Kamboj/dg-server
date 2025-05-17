package domain

import (
	"github.com/google/uuid"
	"time"
)

// UserPreference represents user-specific settings
type UserPreference struct {
	UserID     uuid.UUID `gorm:"primaryKey;not null"`
	SettingKey string    `gorm:"primaryKey;not null"`
	Value      string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	// ✅ Foreign Key to `users.id`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
