package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("1234567890123456789012345678901234567890123456789012345678901234")

func GenerateToken(username *string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"username": username,
	})

	tokenStr, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}


func ValidateToken(token string) (*jwt.Token,error) {

	str, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		return nil,err
	}

	if !str.Valid {
		return nil,fmt.Errorf("Token invalid")
	}

	return str, nil
}