package chat

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/file"
	"github.com/quik/backend/internal/api/token"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/message"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateMessage(c echo.Context) error {
	idChat, err := strconv.Atoi(c.Param("idChat"))
	userFrom := token.GetUserIDFromToken(&c)
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	text := c.FormValue("text")

	database, dbClose := db.Init()
	defer dbClose()

	var message model.Message
	chat := model.Chat{}
	var imageLink = ""

	err = database.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", idChat).First(&chat)
		if result.RowsAffected == 0 {
			return errors.New("Такого чата не существует")
		}

		message = model.Message{
			Text:   text,
			IDChat: chat.ID,
			IDUser: userFrom,
		}

		if err := tx.Create(&message).Error; err != nil {
			return err
		}
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		images := form.File["images"]

		for _, img := range images {

			filename, err := file.SaveFile(img)
			if err != nil {
				return err
			}

			image := model.Image{
				FileName:    filename,
				FileFormat:  img.Header.Get("Content-Type"),
				TypeConnect: _const.MESSAGE_TYPE_CONNECT,
				IDConnect:   message.ID,
			}

			database.Create(&image)

			imageLink = image.FileName
		}

		return nil
	})

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	var (
		conn *websocket.Conn
		ok   bool
	)

	conn, ok = Clients[chat.IDUser1]

	if ok {
		str, _ := json.Marshal(response.MessageResponse{
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
			ImgLink:   imageLink,
			UserID:    message.IDUser,
		})
		conn.WriteMessage(websocket.TextMessage, str)
	}

	conn, ok = Clients[chat.IDUser2]

	if ok {
		str, _ := json.Marshal(response.MessageResponse{
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
			ImgLink:   imageLink,
			UserID:    message.IDUser,
		})
		conn.WriteMessage(websocket.TextMessage, str)
	}

	return c.String(http.StatusOK, "OK")
}
