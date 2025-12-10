package model

type Member struct {
	UserID  int `gorm:"not null"`
	ClassID int `gorm:"not null"`
}

func (Member) TableName() string {
	return "member"
}
