package response

import response "github.com/quik/backend/internal/response/message"

type UserChatInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	ImgLink string `json:"imgLink"`
}

type ChatInfo struct {
	UserTo   UserChatInfo               `json:"userTo"`
	Messages []response.MessageResponse `json:"messages"`
}
