package request

type PostRequest struct {
	Text string `gorm:"column:text"`
}
