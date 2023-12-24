package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

type JwtCustomClaims struct {
	IDUser uint `json:"id_user"`
	jwt.RegisteredClaims
}

func generateToken(userID uint, expiredTime time.Time) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GenerateAccessToken(userID uint) (string, error) {
	return generateToken(userID, time.Now().Add(time.Hour))
}

func GenerateRefreshToken(userID uint) (string, error) {
	return generateToken(userID, time.Now().Add(time.Hour*24*30))
}

func GetUserIDFromToken(c *echo.Context) uint {
	jwtClaim := (*c).Get("token").(*jwt.Token).Claims
	claims := jwtClaim.(*JwtCustomClaims)
	return claims.IDUser
}

func GetUserIDFromTokenString(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return -1, errors.New("No valid token or expired")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID := claims["id_user"].(float64)
		return int(userID), nil
	}
	return -1, errors.New("No valid claims")
}
