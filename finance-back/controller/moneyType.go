package controller

import (
	"finance-api/models"
	"finance-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMoneyType(c *gin.Context) {
	var moneyType models.MoneyType

	if err := c.ShouldBindJSON(&moneyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := service.CreateMoneyType(&moneyType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": moneyType})
}

func GetAllMoneyType(c *gin.Context) {
	listMoneyType, err := service.GetAllMoneyType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": listMoneyType})
}
