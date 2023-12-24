package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quik/backend/internal/db"
	"github.com/quik/backend/internal/router"
	"log"
	"net/http"
	"os"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.CORS())

	database, closeDB := db.Init()
	db.Migration(database)

	router.Init(e)

	router.InitSecurity(e)

	closeDB()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
