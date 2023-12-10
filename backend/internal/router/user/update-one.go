package user

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	request "github.com/quik/backend/internal/request/user"
	"log"
	"net/http"
	"strconv"
)

func UpdateOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Не верный id")
	}

	req := new(request.UserRequest)
	err = c.Bind(req)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Validate(req)
	if err != nil {
		return err
	}

	database, dbClose := db.Init()
	defer dbClose()

	var user model.User
	result := database.First(&user, id)
	if result.RowsAffected == 0 {
		return c.String(http.StatusInternalServerError, "Ошибка Сервера")
	}

	user.Name = req.Name
	user.Surname = req.Surname
	user.ImgLink = req.ImgLink

	database.Save(&user)

	return c.String(http.StatusOK, "Успешно")
}
