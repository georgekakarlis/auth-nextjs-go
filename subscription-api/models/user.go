package models

import (
	"time"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID"`
}