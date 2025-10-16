package model

import "time"

type TestCase struct {
	ID          int                    `gorm:"primaryKey"`
	TestSuiteID int                    `gorm:"not null"`
	Name        string                 `gorm:"not null"`
	Init        map[string]interface{} `gorm:"type:jsonb;not null"`
	Assert      map[string]interface{} `gorm:"type:jsonb;not null"`
	CreateAt    time.Time              `gorm:"autoCreateTime"`
}

func (TestCase) TableName() string {
	return "test_case"
}
