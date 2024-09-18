package models

type Category struct {
	ID           uint          `gorm:"primarykey"`
	NameCategory string        `gorm:"unique;not null" json:"nameCategory"`
	Transaction  []Transaction `gorm:"foreignKey:CategoryID"`
}
