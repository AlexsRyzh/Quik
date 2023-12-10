package response

import (
	"github.com/quik/backend/internal/model"
	"time"
)

type PostResponse struct {
	ID             uint      `json:"id"`
	Text           string    `json:"text"`
	AmountLikes    int64     `json:"amount_likes"`
	AmountComments int64     `json:"amount_comments"`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"updated_at"`
	IsLike         bool      `json:"is_like"`

	IDUser uint `json:"id_user"`

	Images []model.Image
}
