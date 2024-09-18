package models

type Flow struct {
	ID          uint16        `gorm:"primaryKey" json:"id"`
	Type        string        `json:"type"`
	Transaction []Transaction `gorm:"foreignKey:FlowID"`
}
