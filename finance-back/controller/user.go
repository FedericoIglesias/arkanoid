package controller

import (
	"encoding/json"
	"finance-api/models"
	"finance-api/service"
	"finance-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {
	var r UserReq
	if err := json.NewDecoder(c.Request.Body).Decode(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := service.GetUser(&r.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
	var r UserReq
	if err := json.NewDecoder(c.Request.Body).Decode(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.GetUser(&r.Email)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "User or password are wrong"})
		return
	}

	if user.Password != r.Password {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "User or password are wrong"})
		return
	}

	jwt, err := utils.GenerateToken(&r.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"jwt": jwt})
}
