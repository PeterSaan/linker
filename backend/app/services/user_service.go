package services

import (
	"errors"
	"linker/internal/database"
	"linker/internal/models"
	"linker/internal/structs"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func LinkedInGetUser(userData structs.UserDataLinkedIn) (*models.User, error) {
	var user *models.User
	database := database.GetDB()

	log.Printf("User data: %+v", userData)

	if err := database.Where("email ILIKE ?", userData.Email).First(&user).Error; err != nil {
		log.Printf("User not found in database, creating new record \n")
		if err := LinkedInRegisterUser(userData); err != nil {
			return nil, err
		}
	}

	log.Printf("User: %+v\n", user)

	return user, nil
}

func LinkedInRegisterUser(userData structs.UserDataLinkedIn) error {
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

func PasswordRegisterUser(formData structs.PasswordRegiserData) (models.User, error) {
	database := database.GetDB()

	hash, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 12)
	if err != nil {
		log.Printf("Failed to hash password: %v \n", err)
		return models.User{}, err
	}

	user := models.User{Email: formData.Email, Password: string(hash)}

	if err := database.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func PasswordGetUser(formData structs.PasswordLoginData) (models.User, error) {
	database := database.GetDB()
	var user models.User

	if err := database.Where("email ILIKE ?", formData.Email).First(&user).Error; err != nil {
		return models.User{}, errors.New("User with this email address not found")
	}

	if user.Password == "" {
		return models.User{}, errors.New("This user has not registered with a password, try another login method")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password)); err != nil {
		return models.User{}, errors.New("Password does not match")
	}

	return user, nil
}
