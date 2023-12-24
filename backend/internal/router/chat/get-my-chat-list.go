package chat

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response2 "github.com/quik/backend/internal/response/chat"
	"gorm.io/gorm"
	"net/http"
)

func GetMyChatList(c echo.Context) error {
	userID := token.GetUserIDFromToken(&c)

	database, dbClose := db.Init()
	defer dbClose()

	var chatList []response2.ChatItem

	err := database.Transaction(func(tx *gorm.DB) error {
		var chats []model.Chat

		tx.Where("id_user_1 = ? or id_user_2 = ?", userID, userID).Find(&chats)

		for _, chat := range chats {
			var lastMessage model.Message
			var user model.User

			var userTo uint = 0
			if userID != chat.IDUser1 {
				userTo = chat.IDUser1
			} else {
				userTo = chat.IDUser2
			}

			result := tx.Where("id_chat = ?", chat.ID).Order("created_at desc").First(&lastMessage)
			if result.RowsAffected == 0 {
				continue
			}

			tx.Where("id = ?", userTo).First(&user)

			chatList = append(chatList, response2.ChatItem{
				ID:       int(chat.ID),
				IdUserTo: int(userTo),
				Name:     user.Name,
				Surname:  user.Surname,
				ImgLink:  user.ImgLink,
				LastMessage: response2.LastMessage{
					Text: lastMessage.Text,
					Date: lastMessage.CreatedAt,
				},
			})

		}

		return nil
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка сервера")
	}

	return c.JSON(http.StatusOK, chatList)
}
