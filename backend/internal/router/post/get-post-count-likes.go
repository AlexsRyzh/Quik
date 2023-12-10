package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"net/http"
	"strconv"
)

func GetPostCountLike(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var post model.Post
	result := database.Where("id = ?", postID).First(&post)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Не верный id поста")
	}

	return c.JSON(http.StatusOK, &map[string]int64{
		"count": post.AmountLikes,
	})
}
