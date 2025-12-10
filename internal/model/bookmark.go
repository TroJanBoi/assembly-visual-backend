package model

type BookMark struct {
	ID      int `gorm:"primaryKey"`
	UserID  int `gorm:"not null"`
	ClassID int `gorm:"not null"`
}

func (BookMark) TableName() string {
	return "bookmark"
}
