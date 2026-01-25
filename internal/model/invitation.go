package model

import (
	"time"

	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model
	ClassID            int       `gorm:"not null"`
	InvitedEmail       string    `gorm:"not null"`
	InvitedUserID      int       `gorm:"not null"`
	GoogleInvitationID string    `gorm:"null;"`
	Status             string    `gorm:"not null;default:'pending'"` // e.g., "pending", "accepted", "declined"
	token              string    `gorm:"null;"`
	expired            time.Time `gorm:"null"`
	Classroom          Classroom `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User               User      `gorm:"foreignKey:InvitedUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Invitation) TableName() string {
	return "invitation"
}
