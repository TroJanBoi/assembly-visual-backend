package model

import "time"

type TestSuites struct {
	ID           int       `gorm:"primaryKey"`
	AssignmentID int       `gorm:"not null"`
	Name         string    `gorm:"not null"`
	CreateAt     time.Time `gorm:"autoCreateTime"`
}

func (TestSuites) TableName() string {
	return "test_suites"
}
