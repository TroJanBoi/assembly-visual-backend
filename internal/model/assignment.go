package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	ClassID     int            `gorm:"not null"`
	Title       string         `gorm:"not null"`
	Description string         `gorm:"null"`
	DueDate     time.Time      `gorm:"null"`
	MaxAttempt  int            `gorm:"not null;default:0"` // 0 means unlimited
	Setting     datatypes.JSON `gorm:"type:jsonb;null"`
	Condition   datatypes.JSON `gorm:"type:jsonb;null"`
	Grade       int            `gorm:"not null;default:1"` // max grade
	Classroom   Classroom      `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Assignment) TableName() string {
	return "assignment"
}
