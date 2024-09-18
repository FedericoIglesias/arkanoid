package models

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"categoryId"`
	UserID      uint      `json:"userId"`
	FlowID      uint16    `json:"flowId"`
	Date        time.Time `gorm:"autoCreateTime",json:"date"`
	MoneyTypeID uint      `json:"moneyTypeId"`
}
