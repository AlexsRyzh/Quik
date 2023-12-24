package model

import "gorm.io/gorm"

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

func (f *FriendsType) AfterCreate(tx *gorm.DB) (err error) {

	friend := FriendsType{}
	result := tx.Where("id_user_1 = ? and id_user_2 = ?", f.User2ID, f.User1ID).First(&friend)

	user1 := User{}
	tx.Where("id = ?", f.User1ID).First(&user1)

	user2 := User{}
	tx.Where("id = ?", f.User2ID).First(&user2)
	if result.RowsAffected == 0 {
		user1.AmountSubscribe = user1.AmountSubscribe + 1
		user2.AmountSubscribers = user2.AmountSubscribers + 1
	} else {
		user1.AmountSubscribers = user1.AmountSubscribers - 1
		user2.AmountSubscribe = user2.AmountSubscribe - 1
	}

	tx.Save(&user1)
	tx.Save(&user2)

	return
}

func (f *FriendsType) AfterDelete(tx *gorm.DB) (err error) {

	friend := FriendsType{}
	result := tx.Where("id_user_1 = ? and id_user_2 = ?", f.User2ID, f.User1ID).First(&friend)

	user1 := User{}
	tx.Where("id = ?", f.User1ID).First(&user1)

	user2 := User{}
	tx.Where("id = ?", f.User2ID).First(&user2)
	if result.RowsAffected == 0 {
		user1.AmountSubscribe = user1.AmountSubscribe - 1
		user2.AmountSubscribers = user2.AmountSubscribers - 1
	} else {
		user1.AmountSubscribers = user1.AmountSubscribers + 1
		user2.AmountSubscribe = user2.AmountSubscribe + 1
	}

	tx.Save(&user1)
	tx.Save(&user2)

	return
}
