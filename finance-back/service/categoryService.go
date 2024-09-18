package service

import (
	"finance-api/config"
	"finance-api/models"
)

func CreateCategory(category *models.Category) error {
	return config.DB.Create(category).Error
}

func GetCatergory(nameCategory *string) (models.Category, error) {
	var category models.Category

	err := config.DB.Where(&models.Category{NameCategory: *nameCategory}).First(&category).Error

	return category, err
}

func GetAllCategories() ([]models.Category, error) {
	var listCategory []models.Category
	err := config.DB.Find(&listCategory).Error
	return listCategory, err
}
