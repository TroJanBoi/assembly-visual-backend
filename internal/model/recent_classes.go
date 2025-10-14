package model

import "time"

type RecentClasses struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"not null"`
	ClassID      uint      `gorm:"not null"`
	LastViewedAt time.Time `gorm:"autoUpdateTime"`
}

func (RecentClasses) TableName() string {
	return "recent_classes"
}
