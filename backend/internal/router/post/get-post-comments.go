package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	response "github.com/quik/backend/internal/response/post"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetPostComments(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Не верный id")
	}

	database, dbClose := db.Init()
	defer dbClose()

	var post model.Post
	result := database.Where("id = ?", postID).First(&post)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Не верный id поста")
	}

	var comments []model.Comment

	var res []response.PostCommentResponse

	err = database.Transaction(func(tx *gorm.DB) error {
		tx.Where("id_post = ?", postID).Find(&comments)

		for _, v := range comments {

			var user model.User

			tx.Where("id = ?", v.IDUser).First(&user)

			res = append(res, response.PostCommentResponse{
				Name:        user.Name,
				Surname:     user.Surname,
				Message:     v.Message,
				UserImgLink: user.ImgLink,
			})
		}

		return nil
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка сервера")
	}

	return c.JSON(http.StatusOK, res)
}
