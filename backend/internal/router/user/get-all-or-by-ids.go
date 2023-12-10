package user

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"github.com/quik/backend/internal/request"
	response "github.com/quik/backend/internal/response/user"
	"gorm.io/gorm"
	"net/http"
)

func GetAllOrByIDs(c echo.Context) error {

	req := request.IDsRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный запрос")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var users []model.User
	var result *gorm.DB
	if len(req.IDs) == 0 {
		result = database.Find(&users)
	} else {
		result = database.Find(&users, req.IDs)
	}
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Не найдены")
	}

	var res []response.UserResponse
	for _, user := range users {
		res = append(res, response.UserResponse{
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

	return c.JSON(http.StatusOK, res)
}
