package friend

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"net/http"
	"strconv"
)

func AddFriend(c echo.Context) error {
	user1ID := token.GetUserIDFromToken(&c)
	user2ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || user1ID == uint(user2ID) {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, closeDB := db.Init()
	defer closeDB()

	var users []model.User
	result := database.Where("id = ? or id = ?", user1ID, user2ID).Find(&users)
	if result.RowsAffected != 2 {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	var friendType1 model.FriendsType
	result1 := database.Where("id_user_1 = ? and id_user_2 = ?", user1ID, user2ID).First(&friendType1)

	var friendType2 model.FriendsType
	result2 := database.Where("id_user_1 = ? and id_user_2 = ?", user2ID, user1ID).First(&friendType2)

	if result1.RowsAffected == 0 && result2.RowsAffected == 0 {
		friendType := model.FriendsType{
			Type:    _const.FRIENDS_TYPE_SUBCRIBER,
			User1ID: user1ID,
			User2ID: uint(user2ID),
		}

		database.Create(&friendType)
	} else if result1.RowsAffected == 0 && result2.RowsAffected == 1 {
		friendType2.Type = _const.FRIENDS_TYPE_FRIEND
		database.Save(&friendType2)

		friendType := model.FriendsType{
			Type:    _const.FRIENDS_TYPE_FRIEND,
			User1ID: user1ID,
			User2ID: uint(user2ID),
		}

		database.Create(&friendType)
	} else {
		return c.String(http.StatusBadRequest, "Не верный запрос")
	}

	return c.String(http.StatusOK, "Успешно")
}
