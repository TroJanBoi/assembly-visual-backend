package model

type Member struct {
	UserID    int       `gorm:"primaryKey; not null"`
	ClassID   int       `gorm:"primaryKey; not null"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Classroom Classroom `gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Member) TableName() string {
	return "member"
}
