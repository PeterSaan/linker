package migration

import (
	"linker/internal/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	modelList := []interface{}{
		&models.Chat{},
		&models.Chatroom{},
		&models.ChatroomUser{},
		&models.JobTitle{},
		&models.Media{},
		&models.Profile{},
		&models.User{},
	}

	for _, model := range modelList {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
		log.Printf("%T migrated\n", model)
	}

	return nil
}
