package services

import (
	"linker/internal/database"
	"linker/internal/models"
	"linker/internal/structs"
	"log"
)

func LinkedInLogin(userData structs.UserDataLinkedIn) ( *models.User, error ) {
	var user *models.User
    database := database.GetDB()

	log.Printf("User data: %+v", userData)

	if err := database.Where("email = ?", userData.Email).First(&user).Error; err != nil {
		log.Printf("User not found in database, creating new record \n")
        if err := LinkedInRegister(userData); err != nil {
            return nil, err
        }
	 }

	log.Printf("User: %+v\n", user)

    return user, nil
}

func LinkedInRegister(userData structs.UserDataLinkedIn) error {
	user := models.User{Email: userData.Email}
    database := database.GetDB()

	if err := database.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
        return err
	}

	profile := models.Profile{UserID: user.ID, Avatar: userData.Picture, Name: userData.Name}
	if err := database.Create(&profile).Error; err != nil {
		log.Printf("Failed to create users profile: %v", err)
        return err
	}

    return nil
}
