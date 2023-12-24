package model

import (
	"gorm.io/gorm"
	"time"
)

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

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {

	user := User{}
	tx.Where("id = ?", p.IDUser).First(&user)

	user.AmountPosts = user.AmountPosts + 1
	tx.Save(&user)

	return
}
