package response

import "time"

type MessageResponse struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	ImgLink   string    `json:"img_link"`
	UserID    uint      `json:"user_id"`
}
