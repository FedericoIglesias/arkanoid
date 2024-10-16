package config

import (
	"finance-api/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
DSN := os.Getenv("DSN")
	dsn := DSN
	// "host=localhost user=postgres password=root dbname=finance port=5432 sslmode=disable TimeZone=UTC"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Filed to connect to database")
	}

	DB.AutoMigrate(&models.MoneyType{}, &models.User{}, &models.Transaction{}, &models.Category{}, &models.Flow{})
}
