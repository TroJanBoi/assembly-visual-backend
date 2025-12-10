package model

import (
	"gorm.io/gorm"
)

type TestSuites struct {
	gorm.Model
	AssignmentID int    `gorm:"not null"`
	Name         string `gorm:"not null"`
}

func (TestSuites) TableName() string {
	return "test_suites"
}
