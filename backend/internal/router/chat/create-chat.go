package chat

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateChat(c echo.Context) error {
	anotherUserID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	myID := token.GetUserIDFromToken(&c)

	if myID == uint(anotherUserID) {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	chat := model.Chat{
		IDUser1: myID,
		IDUser2: uint(anotherUserID),
	}

	err = database.Transaction(func(tx *gorm.DB) error {

		var user model.User
		result := tx.Where("id = ?", anotherUserID).First(&user)

		if result.RowsAffected == 0 {
			return errors.New("Такого пользователя не существует")
		}

		result = tx.Where("id_user_1 =? and id_user_2 = ? or id_user_2 =? and id_user_1 = ?",
			myID, anotherUserID, myID, anotherUserID).First(&chat)

		if result.RowsAffected != 0 {
			return nil
		}

		tx.Create(&chat)

		return nil
	})

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.String(http.StatusOK, strconv.Itoa(int(chat.ID)))
}
