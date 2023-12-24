package response

type UserResponse struct {
	ID                uint   `json:"id"`
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	AmountPosts       int64  `json:"amount_posts"`
	AmountSubscribers int64  `json:"amount_subscribers"`
	AmountSubscribe   int64  `json:"amount_subscribe"`
	ImgLink           string `json:"img_link"`
	Online            bool   `json:"online"`
	IsFriend          bool   `json:"is_friend"`
}
