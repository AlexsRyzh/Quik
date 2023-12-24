package friend

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"net/http"
	"strconv"
)

func IsFriend(c echo.Context) error {
	user1ID := token.GetUserIDFromToken(&c)
	user2ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || user1ID == uint(user2ID) {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, closeDB := db.Init()
	defer closeDB()

	var friend model.FriendsType
	result := database.Where("id_user_1 = ? and id_user_2 = ?", user1ID, user2ID).First(&friend)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, map[string]bool{
			"is-friend": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"is-friend": true,
	})
}
