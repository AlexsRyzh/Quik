package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Message   string    `gorm:"column:message"`
	CreatedAt time.Time `gorm:"column:created_at"`

	IDUser uint `gorm:"column:id_user"`
	IDPost uint `gorm:"column:id_post"`

	User User `gorm:"foreignKey:IDUser"`
	Post Post `gorm:"foreignKey:IDPost"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {

	post := Post{}
	tx.Where("id = ?", c.IDPost).First(&post)

	post.AmountComments = post.AmountComments + 1
	tx.Save(&post)

	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {

	post := Post{}
	tx.Where("id = ?", c.IDPost).First(&post)

	post.AmountComments = post.AmountComments - 1
	tx.Save(&post)

	return
}
