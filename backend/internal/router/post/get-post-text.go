package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"net/http"
	"strconv"
)

func GetPostText(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var posts model.Post

	result := database.First(&posts, id)

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Текст не найден")
	}

	return c.String(http.StatusOK, posts.Text)
}
