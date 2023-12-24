package chat

import (
	"github.com/labstack/echo/v4"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/message"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetMessages(c echo.Context) error {
	idChat, err := strconv.Atoi(c.Param("idChat"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var res []response.MessageResponse

	err = database.Transaction(func(tx *gorm.DB) error {
		var messages []model.Message

		tx.Where("id_chat = ?", idChat).Find(&messages)

		for _, message := range messages {

			img := model.Image{}

			tx.Where("id_connect = ? and type_connect = ?", message.ID, _const.MESSAGE_TYPE_CONNECT).First(&img)

			res = append(res, response.MessageResponse{
				Text:      message.Text,
				CreatedAt: message.CreatedAt,
				ImgLink:   img.FileName,
				UserID:    message.IDUser,
			})
		}

		return nil
	})

	return c.JSON(http.StatusOK, res)
}
