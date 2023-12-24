package response

type SubFriendResponse struct {
	ID      int    `json:"id"`
	ImgLink string `json:"imgLink"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
