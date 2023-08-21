package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Sessions  []Session `gorm:"foreignKey:AccountID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
