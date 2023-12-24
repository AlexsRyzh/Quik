package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/post"
	"net/http"
	"strconv"
)

func GetPostIDsByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var posts []model.Post

	database.Where("id_user = ?", userID).Order("created_at desc").Find(&posts)

	var ids []uint
	var userIDs []uint

	for _, post := range posts {

		ids = append(ids, post.ID)
		userIDs = append(userIDs, post.IDUser)
	}

	res := response.PostIDsResponse{
		IDs:    ids,
		UserID: userIDs,
	}

	return c.JSON(http.StatusOK, res)
}
