package model

import (
	"time"
)

type GoogleAccount struct {
	GoogleUserID string    `gorm:"primaryKey;not null"`
	UserID       int       `gorm:"not null"`
	AccessToken  string    `gorm:"null"`
	RefreshToken string    `gorm:"null"`
	ExpiredAt    time.Time `gorm:"null"`
	CreatedAt    time.Time `gorm:"not null"`
	User         User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (GoogleAccount) TableName() string {
	return "google_account"
}
