package model

import "time"

type RecentClasses struct {
	ID           int       `gorm:"primaryKey"`
	UserID       int       `gorm:"not null"`
	ClassID      int       `gorm:"not null"`
	LastViewedAt time.Time `gorm:"autoUpdateTime"`
}

func (RecentClasses) TableName() string {
	return "recent_classes"
}
