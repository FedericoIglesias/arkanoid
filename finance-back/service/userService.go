package service

import (
	"errors"
	"finance-api/config"
	"finance-api/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUser(email *string) (models.User, error) {
	var user models.User

	err := config.DB.Where(&models.User{Email: *email}).First(&user).Error
	return user, err
}

func RegisterJWT(token *string, email *string) error {
	err := config.DB.Model(&models.User{}).Where("email = ?", &email).Update("token", token)
	if err != nil {
		return err.Error
	}
	return nil
}

func SearchJWT(token *string) error {
	var user models.User
	err := config.DB.Where(&models.User{Token: *token}).First(&user).Error

	if err != nil {
		return errors.New("Fail to found jwt")
	}
	return nil
}
