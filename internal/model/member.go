package model

import "time"

type Member struct {
	UserID    int       `gorm:"primaryKey; not null"`
	ClassID   int       `gorm:"primaryKey; not null"`
	Role      string    `gorm:"not null; default:'member'"` // roles: member, teacher, ta
	JoinAt    time.Time `gorm:"not null; default:current_timestamp"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Classroom Classroom `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Member) TableName() string {
	return "member"
}
