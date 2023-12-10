package db

import (
	"github.com/quik/backend/internal/model"
	"gorm.io/gorm"
	"log"
)

func Migration(database *gorm.DB) {
	err := database.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Like{},
		&model.Comment{},
		&model.Chat{},
		&model.Message{},
		&model.Image{},
		&model.FriendsType{},
	)
	if err != nil {
		log.Fatal("Error Migration")
	}
}
