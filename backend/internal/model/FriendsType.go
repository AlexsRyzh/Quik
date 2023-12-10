package model

type FriendsType struct {
	ID      uint   `gorm:"primaryKey"`
	Type    string `gorm:"column:type"`
	User1ID uint   `gorm:"column:id_user_1"`
	User2ID uint   `gorm:"column:id_user_2"`

	User1 User `gorm:"foreignKey:User1ID"`
	User2 User `gorm:"foreignKey:User2ID"`
}

func (FriendsType) TableName() string {
	return "friends_type"
}
