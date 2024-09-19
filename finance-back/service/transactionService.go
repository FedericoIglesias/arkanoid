package service

import (
	"finance-api/config"
	"finance-api/models"
)

type Result struct {
	Sum  float64
	Less float64
}

func CreateTransaction(transaction *models.Transaction) error {
	return config.DB.Create(transaction).Error
}

func GetTransaction(userId *int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := config.DB.Where("user_id = ?", userId).Find(&transactions).Error
	return transactions, err
}

func GetBalance(userId *int) (Result, error) {

	var result Result

	tx := config.DB.Raw(`select 
(select sum(t.amount) from "transactions" t  where t.flow_id = 1 and t.user_id = ?) as sum,
(select sum(t.amount) from "transactions" t  where t.flow_id = 2 and t.user_id = ?) as less`, userId, userId).Scan(&result)

	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}
