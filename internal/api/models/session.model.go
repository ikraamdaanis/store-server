package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Token      string    `gorm:"type:varchar(255);not null"`
	UserAgent  string    `gorm:"type:varchar(255);not null"`
	IP_Address string    `gorm:"type:varchar(255);not null"`
	ExpiresAt  time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
