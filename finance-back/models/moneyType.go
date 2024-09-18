package models

type MoneyType struct {
	ID          uint          `gorm:"primarykey"`
	MoneyType   string        `gorm:"unique;not null" json:"moneyType"`
	Transaction []Transaction `gorm:"foreignKey:MoneyTypeID"`
}
