package model

import "time"

type GoogleCourseSyncLog struct {
	ID       uint                   `gorm:"primaryKey"`
	UserID   uint                   `gorm:"not null"`
	Action   string                 `gorm:"not null"` // e.g., "sync_started", "sync_completed", "sync_failed"
	Response map[string]interface{} `gorm:"type:jsonb;not null"`
	Status   string                 `gorm:"not null"` // e.g., "success", "failure"
	CreateAt time.Time              `gorm:"autoCreateTime"`
}

func (GoogleCourseSyncLog) TableName() string {
	return "google_course_sync_log"
}
