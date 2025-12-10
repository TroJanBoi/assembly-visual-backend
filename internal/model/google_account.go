package model

import (
	"time"
)

type GoogleAccount struct {
	UserID       int       `gorm:"primaryKey"`
	GoogleUserID string    `gorm:"uniqueIndex;not null"`
	Email        string    `gorm:"not null"`
	AccessToken  string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expired      time.Time `gorm:"not null"`
}

func (GoogleAccount) TableName() string {
	return "google_account"
}
