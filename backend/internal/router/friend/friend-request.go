package friend

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/friend"
	"net/http"
)

func FriendRequests(c echo.Context) error {
	user1ID := token.GetUserIDFromToken(&c)

	database, closeDB := db.Init()
	defer closeDB()

	var friends []model.FriendsType
	database.Where("id_user_2 = ? and type =? ", user1ID, _const.FRIENDS_TYPE_SUBCRIBER).Find(&friends)

	var sub []response.SubFriendResponse

	for _, friend := range friends {

		var user model.User
		database.Where("id = ?", friend.User1ID).First(&user)

		sub = append(sub, response.SubFriendResponse{
			ID:      int(friend.User1ID),
			ImgLink: user.ImgLink,
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	return c.JSON(http.StatusOK, sub)
}
