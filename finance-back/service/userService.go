package service

import (
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

func RegisterJWT(token string, email *string) error {
	err := config.DB.Model(&models.User{}).Where("email = ?", &email).Update("token", token)
	if err != nil {
		return err.Error
	}
	return nil
}
