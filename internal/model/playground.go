package model

import "gorm.io/gorm"

type Playground struct {
	gorm.Model
	AssignmentID uint                   `gorm:"not null"`
	UserID       uint                   `gorm:"not null"`
	AttemptNO    uint                   `gorm:"not null"`
	Item         map[string]interface{} `gorm:"type:jsonb"`
	Status       string                 `gorm:"not null"` // e.g., "in_progress", "completed", "failed"
}

func (Playground) TableName() string {
	return "playground"
}
