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

func Login(c echo.Context) error {
	req := new(request.LoginRequest)
	err := c.Bind(req)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Validate(req)
	if err != nil {
		return err
	}

	database, dbClose := db.Init()
	defer dbClose()

	var user model.User
	result := database.Where("login = ?", req.Login).First(&user)
	if result.RowsAffected == 0 {
		return c.String(http.StatusBadRequest, "Не верный логин или пароль")
	}

	pwdHash, err := hash.HashPassword(req.Password)
	if err != nil {
		log.Fatal(err)
	}

	if !hash.CheckPasswordHash(req.Password, pwdHash) {
		return c.String(http.StatusBadRequest, "Не верный логин или пароль")
	}

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
