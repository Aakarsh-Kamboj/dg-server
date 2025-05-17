package domain

import "time"

type UserProfile struct {
	UserID       string    `gorm:"primaryKey"`
	FirstName    *string   `gorm:"default:null"`
	LastName     *string   `gorm:"default:null"`
	ProfileImage *string   `gorm:"default:null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
