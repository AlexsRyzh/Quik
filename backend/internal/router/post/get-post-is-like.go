package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/post"
	"net/http"
	"strconv"
)

func GetPostIsLike(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	userID := token.GetUserIDFromToken(&c)

	database, dbClose := db.Init()
	defer dbClose()

	var post model.Post
	result := database.Where("id = ?", postID).First(&post)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Не верный id поста")
	}

	var like model.Like

	result = database.Where("id_user = ? and id_post = ?", userID, postID).First(&like)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, response.PostIsLike{
			IsLike: false,
		})
	}

	return c.JSON(http.StatusOK, response.PostIsLike{
		IsLike: true,
	})
}
