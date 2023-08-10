package models

// a Subscription struct to keep track of the users subscriptions

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	UserID uint
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email string `gorm:"not null"`
	SubName string `gorm:"not null"`
	SubPrice float64 `gorm:"not null"`
	SubDate time.Time `gorm:"not null"`
	SubType string `gorm:"not null"`
	SubCategory string `gorm:"not null"`
	SubFrequency string `gorm:"not null"`
	SubNotes string `gorm:"not null"`
}
