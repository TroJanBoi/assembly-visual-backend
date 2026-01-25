package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GoogleCourseSyncLog struct {
	gorm.Model
	ClassID   int            `gorm:"not null"`
	Action    string         `gorm:"not null;default:'create'"` // e.g., "create", "update", "delete"
	Response  datatypes.JSON `gorm:"type:jsonb;null"`
	Status    string         `gorm:"not null"` // e.g., "success", "failure"
	Classroom Classroom      `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (GoogleCourseSyncLog) TableName() string {
	return "google_course_sync_log"
}
