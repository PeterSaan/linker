package seeder

import (
	"linker/internal/models"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
    profiles := []*models.Profile{
        { Name: "HelloWorld", Description: "Pleb", Type: "Recruiter", UserID: 1 },
    }

    users := []*models.User{
        { Email: "hello@example.com", Password: "1234" },
    }

	if err := db.Create(users).Error; err != nil {
		log.Println("Error inserting users")
		return err
	}

	if err := db.Create(profiles).Error; err != nil {
		log.Println("Error inserting users")
		return err
	}

	return nil
}
