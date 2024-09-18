package service

import (
	"finance-api/config"
	"finance-api/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	return config.DB.Create(transaction).Error
}

func GetTransaction(userId *int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := config.DB.Where("user_id = ?", userId).Find(&transactions).Error
	return transactions, err
}
