package models

type User struct {
	ID          uint          `gorm:"primarykey;not null, autoincrement" json:"id"`
	Name        string        `gorm:"not null" json:"name"`
	Email       string        `gorm:"not null" json:"email"`
	Password    string        `gorm:"not null" json:"password"`
	Token       string        `json:"token"`
	Transaction []Transaction `gorm:"foreignKey:UserID"`
}
