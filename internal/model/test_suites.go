package model

import "time"

type TestSuites struct {
	ID           uint      `gorm:"primaryKey"`
	AssignmentID uint      `gorm:"not null"`
	Name         string    `gorm:"not null"`
	CreateAt     time.Time `gorm:"autoCreateTime"`
}

func (TestSuites) TableName() string {
	return "test_suites"
}
