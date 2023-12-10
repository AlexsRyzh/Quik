package model

import "gorm.io/gorm"

type Like struct {
	ID uint `gorm:"primaryKey"`

	IDUser uint `gorm:"column:id_user"`
	IDPost uint `gorm:"column:id_post"`

	User User `gorm:"foreignKey:IDUser"`
	Post Post `gorm:"foreignKey:IDPost"`
}

func (Like) TableName() string {
	return "likes"
}

func (l *Like) AfterCreate(tx *gorm.DB) (err error) {

	post := Post{}
	tx.Where("id = ?", l.IDPost).First(&post)

	post.AmountLikes = post.AmountLikes + 1
	tx.Save(&post)

	return
}

func (l *Like) AfterDelete(tx *gorm.DB) (err error) {

	post := Post{}
	tx.Where("id = ?", l.IDPost).First(&post)

	post.AmountLikes = post.AmountLikes - 1
	tx.Save(&post)

	return
}
