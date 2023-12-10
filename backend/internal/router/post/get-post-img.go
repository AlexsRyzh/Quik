package post

import (
	"github.com/labstack/echo/v4"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"net/http"
	"strconv"
)

func GetPostImg(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var image model.Image

	result := database.Where("type_connect = ? and id_connect = ?", _const.POST_TYPE_CONNECT, id).First(&image)

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Картинка не найдена")
	}

	return c.String(http.StatusOK, image.FileName)
}
