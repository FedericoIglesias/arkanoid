package controller

import (
	"finance-api/models"
	"finance-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	flagError := 0
	listError := []string{}
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if transaction.Amount <= 0 {
		listError = append(listError, "The amount can't be less or equal than 0")
		flagError = 1
	}

	if len(transaction.Description) > 50 {
		listError = append(listError, "The description can't be more than 50 characters")
		flagError = 1
	}

	if flagError == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": listError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})

}

func GetTransactions(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("idUser"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	listTransaction, err := service.GetTransaction(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": listTransaction})
}
