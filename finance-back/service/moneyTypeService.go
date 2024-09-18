package service

import (
	"finance-api/config"
	"finance-api/models"
)

func CreateMoneyType(moneyType *models.MoneyType) error {
	return config.DB.Create(moneyType).Error
}

func GetAllMoneyType() ([]models.MoneyType, error) {
	var listMoneyType []models.MoneyType

	err := config.DB.Find(&listMoneyType).Error
	return listMoneyType, err
}
