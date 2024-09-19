package utils

import (
	"finance-api/service"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

// 1234567890123456789012345678901234567890123456789012345678901234
var jwtKey = []byte("1234567890123456789012345678901234567890123456789012345678901234")

func GenerateToken(email *string) (string, error) {

	claims := CustomClaims{
		"bar",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			Subject:   *email,
			ID:        "2",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	err = service.RegisterJWT(tokenStr, email)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		println(err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	return token, nil
}
