package response

type PostCommentResponse struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Message     string `json:"message"`
	UserImgLink string `json:"user_img_link"`
}
