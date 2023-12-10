package post

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/file"
	"github.com/quik/backend/internal/api/token"
	_const "github.com/quik/backend/internal/const"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	"gorm.io/gorm"
	"net/http"
)

func CreateOne(c echo.Context) error {

	text := c.FormValue("text")

	userID := token.GetUserIDFromToken(&c)

	database, dbClose := db.Init()
	defer dbClose()

	err := database.Transaction(func(tx *gorm.DB) error {
		post := model.Post{
			Text:   text,
			IDUser: userID,
		}

		if err := tx.Create(&post).Error; err != nil {
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
				TypeConnect: _const.POST_TYPE_CONNECT,
				IDConnect:   post.ID,
			}

			database.Create(&image)
		}

		return nil
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось создать пост")
	}

	return c.String(http.StatusOK, "Успешно")
}
