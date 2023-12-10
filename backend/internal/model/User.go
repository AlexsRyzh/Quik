package model

type User struct {
	ID                uint   `gorm:"primaryKey"`
	Tag               string `gorm:"column:tag"`
	Name              string `gorm:"column:name"`
	Surname           string `gorm:"column:surname"`
	Login             string `gorm:"column:login"`
	PswHash           string `gorm:"column:psw_hash"`
	AmountPosts       int64  `gorm:"column:amount_posts"`
	AmountSubscribers int64  `gorm:"column:amount_subscribers"`
	AmountSubscribe   int64  `gorm:"column:amount_subscribe"`
	ImgLink           string `gorm:"column:img_link"`
	Online            bool   `gorm:"column:online"`
	RefreshToken      string `gorm:"column:refresh_token"`

	Posts []Post `gorm:"foreignKey:IDUser"`
}

func (User) TableName() string {
	return "users"
}
