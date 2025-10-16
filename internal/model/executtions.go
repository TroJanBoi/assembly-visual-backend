package model

import (
	"time"

	"gorm.io/gorm"
)

type Executions struct {
	gorm.Model
	ExecutionsUUID string                 `gorm:"not null;uniqueIndex"`
	AssignmentID   int                    `gorm:"not null"`
	PlaygroundID   int                    `gorm:"not null"`
	StartAt        time.Time              `gorm:"null"`
	FinishAt       time.Time              `gorm:"null"`
	DurationMs     int64                  `gorm:"default:0"` // Duration in milliseconds
	StepCount      int                    `gorm:"default:0"`
	Status         int                    `gorm:"not null"` // 0: pending, 1: running, 2: completed, 3: failed
	ErrorCode      string                 `gorm:"null"`
	FinalState     map[string]interface{} `gorm:"type:jsonb;null"`
	FullLogPath    string                 `gorm:"null"`
}

func (Executions) TableName() string {
	return "executions"
}
