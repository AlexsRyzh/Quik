package like

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"net/http"
	"strconv"
)

func DeleteOne(c echo.Context) error {

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
	like := model.Like{}

	result = database.Where("id_user = ? and id_post = ?", userID, postID).First(&like)
	if result.RowsAffected == 0 {
		return c.String(http.StatusBadRequest, "Нет лайка")
	}

	database.Delete(&like)

	return c.String(http.StatusOK, "Успешно")
}
