package seeder

import (
	"fmt"
	"linker/internal/database"
	"linker/internal/models"

	"github.com/go-faker/faker/v4"
)

func ProfileSeeder() error {
	db := database.GetDB()
	var users []models.User

	profiles := []models.Profile{}
    db.Model(&models.User{}).Preload("Profile").Find(&users)

	fmt.Println("ProfileSeeder:: Creating user profiles")

	for _, u := range users {
		if u.Profile.UserID != u.ID {
            avatar := "https://picsum.photos/seed/" + faker.Password() + "/200/200"
            profile := models.Profile{Name: faker.Name(), Description: faker.Paragraph(), UserID: u.ID, Avatar: avatar}

            profiles = append(profiles, profile)
		}
	}

	if err := db.Create(profiles).Error; err != nil {
		return err
	}

	fmt.Println("ProfileSeeder:: DONE")

	return nil
}
