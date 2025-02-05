package seeder

import (
	"linker/internal/models"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	chats := []*models.Chat{
		{ Body: 1, UserID: 1, ChatroomID: 1},
		{ Body: 2, UserID: 2, ChatroomID: 2},
	}
	
	if err := db.Create(chats).Error; err != nil {
		log.Println("Error inserting chats")
		return err
	}

	return nil
}
