package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	UserID        int            `gorm:"not null"`
	AssignmentID  int            `gorm:"not null"`
	PlaygroundID  int            `gorm:"not null"`
	AttemptNumber int            `gorm:"not null;default:0"`
	ItemSnapshot  datatypes.JSON `gorm:"type:jsonb;null"`
	ClientResult  datatypes.JSON `gorm:"type:jsonb;null"`
	ServerResult  datatypes.JSON `gorm:"type:jsonb;null"`
	Score         float64        `gorm:"null"`
	Status        string         `gorm:"not null;default:verified"` // e.g., "submitted", "verified", "failed"
	IsVerified    bool           `gorm:"default:false"`
	DurationMS    int            `gorm:"not null"`
	User          User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Assignment    Assignment     `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Playground    Playground     `gorm:"foreignKey:PlaygroundID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Submission) TableName() string {
	return "submission"
}
