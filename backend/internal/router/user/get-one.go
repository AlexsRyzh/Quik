package user

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/user"
	"net/http"
	"strconv"
)

func GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var user model.User
	result := database.First(&user, id)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Не найден")
	}

	return c.JSON(http.StatusOK, &response.UserResponse{
		ID:                user.ID,
		Tag:               user.Tag,
		Name:              user.Name,
		Surname:           user.Surname,
		AmountPosts:       user.AmountPosts,
		AmountSubscribers: user.AmountSubscribers,
		AmountSubscribe:   user.AmountSubscribe,
		ImgLink:           user.ImgLink,
		Online:            user.Online,
	})
}
