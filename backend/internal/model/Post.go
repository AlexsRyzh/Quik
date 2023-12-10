package model

import "time"

type Post struct {
	ID             uint      `gorm:"primaryKey"`
	Text           string    `gorm:"column:text"`
	AmountLikes    int64     `gorm:"column:amount_likes"`
	AmountComments int64     `gorm:"column:amount_comments"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdateAt       time.Time `gorm:"column:updated_at"`
	
	IDUser uint `gorm:"column:id_user"`

	User User `gorm:"foreignKey:IDUser"`
}

func (Post) TableName() string {
	return "posts"
}
