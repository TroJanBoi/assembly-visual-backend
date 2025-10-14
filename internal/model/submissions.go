package model

import "gorm.io/gorm"

type Submissions struct {
	gorm.Model
	PlaygroundID        uint                   `gorm:"not null"`
	UserID              uint                   `gorm:"not null"`
	AssignmentID        uint                   `gorm:"not null"`
	SubmissionAttemptNo uint                   `gorm:"not null"`
	ExecutionID         uint                   `gorm:"not null"`
	Status              string                 `gorm:"not null"`
	ResultSummary       map[string]interface{} `gorm:"type:jsonb"`
	Grade               int                    `gorm:"default:0"`
}

func (Submissions) TableName() string {
	return "submissions"
}
