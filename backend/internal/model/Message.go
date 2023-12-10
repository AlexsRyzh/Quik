package model

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at"`

	IDChat uint `gorm:"column:id_chat"`
	IDUser uint `gorm:"column:id_user"`

	Chat Chat `gorm:"foreignKey:IDChat"`
	User User `gorm:"foreignKey:IDUser"`
}

func (Message) TableName() string {
	return "messages"
}
