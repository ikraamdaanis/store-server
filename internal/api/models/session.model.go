package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uint      `gorm:"primaryKey"`
	AccountID uuid.UUID `gorm:"type:uuid;not null"`
	Token     string    `gorm:"type:varchar(255);not null"`
	UserAgent string    `gorm:"type:varchar(255);not null"`
	IP        string    `gorm:"type:varchar(255);not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
