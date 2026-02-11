package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TestCase struct {
	gorm.Model
	TestSuiteID int            `gorm:"not null"`
	Name        string         `gorm:"null"`
	Init        datatypes.JSON `gorm:"type:jsonb;not null"`
	Assert      datatypes.JSON `gorm:"type:jsonb;not null"`
	TestSuite   TestSuite      `gorm:"foreignKey:TestSuiteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (TestCase) TableName() string {
	return "test_case"
}
