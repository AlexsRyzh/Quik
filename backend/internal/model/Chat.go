package model

type Chat struct {
	ID uint `gorm:"primaryKey"`

	IDUser1 uint `gorm:"column:id_user_1"`
	IDUser2 uint `gorm:"column:id_user_2"`

	User1 User `gorm:"foreignKey:IDUser1"`
	User2 User `gorm:"foreignKey:IDUser2"`
}

func (Chat) TableName() string {
	return "chats"
}
