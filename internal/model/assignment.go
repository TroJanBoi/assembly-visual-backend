package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	ClassID     uint                   `gorm:"not null"`
	Title       string                 `gorm:"not null"`
	Description string                 `gorm:"null"`
	DueDate     time.Time              `gorm:"null"`
	MaxAttempt  int                    `gorm:"default:1"`
	AllowLate   bool                   `gorm:"default:false"`
	Setting     map[string]interface{} `gorm:"type:jsonb;not null"`
	Condition   map[string]interface{} `gorm:"type:jsonb;not null"`
	Grade       int                    `gorm:"default:0"`
}

func (Assignment) TableName() string {
	return "assignment"
}
