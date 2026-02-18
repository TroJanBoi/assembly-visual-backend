package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Playground struct {
	gorm.Model
	AssignmentID int            `gorm:"not null"`
	UserID       int            `gorm:"not null"`
	Item         datatypes.JSON `gorm:"type:jsonb; not null"`
	Status       string         `gorm:"not null;default:'in_progress'"` // e.g., "pending", "completed", "failed"
	Assignment   Assignment     `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Playground) TableName() string {
	return "playground"
}
