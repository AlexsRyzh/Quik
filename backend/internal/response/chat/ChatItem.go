package response

import "time"

type LastMessage struct {
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}

type ChatItem struct {
	ID          int         `json:"id"`
	IdUserTo    int         `json:"idUserTo"`
	Name        string      `json:"name"`
	ImgLink     string      `json:"imgLink"`
	Surname     string      `json:"surname"`
	LastMessage LastMessage `json:"lastMessage"`
}
