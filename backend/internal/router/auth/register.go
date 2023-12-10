package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/hash-pwd"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/model"
	request "github.com/quik/backend/internal/request/auth"
	response "github.com/quik/backend/internal/response/auth"
	"log"
	"net/http"
)

func Register(c echo.Context) error {

	req := new(request.RegisterRequest)
	err := c.Bind(req)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Validate(req)
	if err != nil {
		log.Fatal(1)
		return err
	}

	database, dbClose := db.Init()
	defer dbClose()

	var user model.User
	result := database.Where("login = ?", req.Login).First(&user)
	if result.RowsAffected > 0 {
		return c.String(http.StatusBadRequest, "Такой пользователь существует")
	}

	pwdHash, err := hash.HashPassword(req.Password)
	if err != nil {
		log.Fatal(err)
	}

	user = model.User{
		Tag:     req.Tag,
		Name:    req.Name,
		Surname: req.Surname,
		Login:   req.Login,
		PswHash: pwdHash,
	}

	database.Create(&user)

	accessToken, err := token.GenerateAccessToken(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка Сервера")
	}

	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка Сервера")
	}

	user.RefreshToken = refreshToken
	database.Save(&user)

	return c.JSON(http.StatusOK, &response.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	})
}
