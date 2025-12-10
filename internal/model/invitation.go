package model

import "gorm.io/gorm"

type Invitation struct {
	gorm.Model
	ClassID            int    `gorm:"not null"`
	UserID             int    `gorm:"not null"`
	InvitationEmail    string `gorm:"not null"`
	GoogleInvitationID string `gorm:"null;"`
	Status             string `gorm:"not null"` // e.g., "pending", "accepted", "declined"
}

func (Invitation) TableName() string {
	return "invitation"
}
