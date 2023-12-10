package comment

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	request "github.com/quik/backend/internal/request/comments"
	"log"
	"net/http"
	"strconv"
)

func CreateOne(c echo.Context) error {

	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	userID := token.GetUserIDFromToken(&c)

	req := new(request.CommentRequest)
	err = c.Bind(req)
	if err != nil {
		log.Fatal(err)
	}

	database, dbClose := db.Init()
	defer dbClose()

	var post model.Post
	result := database.Where("id = ?", postID).First(&post)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Не верный id поста")
	}

	comment := model.Comment{
		IDUser:  userID,
		IDPost:  uint(postID),
		Message: req.Message,
	}

	database.Create(&comment)

	return c.String(http.StatusOK, "Успешно")
}
