package chat

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/chat"
	response2 "github.com/quik/backend/internal/response/message"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetChatInfoByID(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id чата")
	}

	myID := token.GetUserIDFromToken(&c)

	database, dbClose := db.Init()
	defer dbClose()

	chatInfo := response.ChatInfo{}

	err = database.Transaction(func(tx *gorm.DB) error {

		var chat model.Chat

		result := tx.Where("id = ?", chatID).First(&chat)
		if result.RowsAffected == 0 {
			return errors.New("Чат не найден")
		}

		if chat.IDUser1 != myID && chat.IDUser2 != myID {
			return errors.New("Отказано в доступе")
		}

		var userTo = 0
		if chat.IDUser1 != myID {
			userTo = int(chat.IDUser1)
		} else {
			userTo = int(chat.IDUser2)
		}

		var userToModel model.User
		tx.Where("id = ?", userTo).First(&userToModel)

		chatInfo.UserTo.ID = int(userToModel.ID)
		chatInfo.UserTo.Name = userToModel.Name
		chatInfo.UserTo.Surname = userToModel.Surname
		chatInfo.UserTo.ImgLink = userToModel.ImgLink

		var messages []model.Message
		tx.Where("id_chat = ?", chatID).Order("created_at asc").Find(&messages)

		for _, message := range messages {

			img := model.Image{}

			tx.Where("id_connect = ? and type_connect = ?", message.ID, _const.MESSAGE_TYPE_CONNECT).First(&img)

			chatInfo.Messages = append(chatInfo.Messages, response2.MessageResponse{
				ID:        int(message.ID),
				Text:      message.Text,
				CreatedAt: message.CreatedAt,
				ImgLink:   img.FileName,
				UserID:    message.IDUser,
			})
		}

		return nil
	})

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, chatInfo)
}
