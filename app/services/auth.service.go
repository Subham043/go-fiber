package services

import (
	"github.com/subham043/go-fiber/app/models"
	"github.com/subham043/go-fiber/platform/database"
)

func GetUsers() []models.User {
	return []models.User{
		{
			Name: "John Doe",
		},
	}
}

func RegisterUser(user models.User) (*models.User, error) {
	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	user := models.User{Email: email}
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
