package model

import "time"

type RecentViewClass struct {
	ID           int       `gorm:"primaryKey; not null"`
	UserID       int       `gorm:"not null"`
	ClassID      int       `gorm:"not null"`
	LastViewedAt time.Time `gorm:"autoUpdateTime"`
	User         User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Classroom    Classroom `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (RecentViewClass) TableName() string {
	return "recent_view_class"
}
