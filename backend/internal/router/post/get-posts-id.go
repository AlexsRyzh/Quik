package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/post"
	"net/http"
)

func GetPostIDs(c echo.Context) error {
	database, dbClose := db.Init()
	defer dbClose()

	var posts []model.Post

	database.Order("created_at desc").Find(&posts)

	var ids []uint
	var userID []uint

	for _, post := range posts {

		ids = append(ids, post.ID)
		userID = append(userID, post.IDUser)
	}

	res := response.PostIDsResponse{
		IDs:    ids,
		UserID: userID,
	}

	return c.JSON(http.StatusOK, res)
}
