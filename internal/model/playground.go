package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Playground struct {
	gorm.Model
	AssignmentID int            `gorm:"not null"`
	UserID       int            `gorm:"not null"`
	AttemptNO    int            `gorm:"not null"`
	Item         datatypes.JSON `gorm:"type:jsonb"`
	Status       string         `gorm:"not null"` // e.g., "in_progress", "completed", "failed"
}

func (Playground) TableName() string {
	return "playground"
}
