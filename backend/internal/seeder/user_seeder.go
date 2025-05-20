package seeder

import (
	"fmt"
	"linker/internal/database"
	"linker/internal/models"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() error {
    db := database.GetDB()

    users := []models.User{}

    fmt.Println("UserSeeder:: Creating users")

    for i := 0; i < 100; i++ {
        hash, err := bcrypt.GenerateFromPassword([]byte(faker.Password()), 10)
        if err != nil {
            return err
        }
		user := models.User{ Email: faker.Email(), Password: string(hash) }

        users = append(users, user)
	}

    if err := db.Create(users).Error; err != nil {
        return err
    }

    fmt.Println("UserSeeder:: DONE")
    
	return nil
}
