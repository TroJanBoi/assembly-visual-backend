package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TestCase struct {
	gorm.Model
	TestSuiteID int            `gorm:"not null"`
	Name        string         `gorm:"not null"`
	Init        datatypes.JSON `gorm:"type:jsonb;not null"`
	Assert      datatypes.JSON `gorm:"type:jsonb;not null"`
}

func (TestCase) TableName() string {
	return "test_case"
}
