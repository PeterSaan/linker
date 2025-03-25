package migration

import (
	"linker/internal/models"
	"log"
	"reflect"

	"gorm.io/gorm"
)

var ModelList = []interface{}{
	&models.Chat{},
	&models.Chatroom{},
	&models.JobTitle{},
	&models.Media{},
	&models.Profile{},
	&models.User{},
	"chatroom_users",
	"profile_media",
	"user_jobtitles",
}

func Migrate(db *gorm.DB) error {
	for _, model := range ModelList {
		if reflect.TypeOf(model).Kind() == reflect.String {
            continue
		}
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
		log.Printf("%T migrated\n", model)
	}

	return nil
}
