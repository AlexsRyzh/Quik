package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/post"
	"net/http"
)

func GetMyPost(c echo.Context) error {
	userID := token.GetUserIDFromToken(&c)

	database, dbClose := db.Init()
	defer dbClose()

	var posts []model.Post

	database.Where("id_user = ?", userID).Find(&posts)

	var res []response.PostResponse

	for _, post := range posts {

		var images []model.Image

		database.Where("id_connect = ? and type_connect = ?", post.ID, _const.POST_TYPE_CONNECT).Find(&images)

		res = append(res, response.PostResponse{
			ID:             post.ID,
			Text:           post.Text,
			AmountLikes:    post.AmountLikes,
			AmountComments: post.AmountComments,
			CreatedAt:      post.CreatedAt,
			UpdateAt:       post.UpdateAt,
			Images:         images,
		})
	}

	return c.JSON(http.StatusOK, res)
}
