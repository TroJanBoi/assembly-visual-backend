package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Submissions struct {
	gorm.Model
	PlaygroundID        int            `gorm:"not null"`
	UserID              int            `gorm:"not null"`
	AssignmentID        int            `gorm:"not null"`
	SubmissionAttemptNo int            `gorm:"not null"`
	ExecutionID         int            `gorm:"not null"`
	Status              string         `gorm:"not null"`
	ResultSummary       datatypes.JSON `gorm:"type:jsonb"`
	Grade               int            `gorm:"default:0"`
}

func (Submissions) TableName() string {
	return "submissions"
}
