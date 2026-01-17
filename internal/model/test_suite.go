package model

import (
	"gorm.io/gorm"
)

type TestSuite struct {
	gorm.Model
	AssignmentID int        `gorm:"not null"`
	Name         string     `gorm:"null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (TestSuite) TableName() string {
	return "test_suite"
}
