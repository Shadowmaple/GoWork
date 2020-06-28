package main

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZXN0IiwiZXhwIjoxNTkyNjUzMzI0LCJzdWIiOiIxMjMifQ.QAYq3BgTIgBHkDYqdYkn5RA3sNZJ_03AzkNbE_uYtJiuwlwEiEF1xnpUZZbpR9lrzvrE2YMKxPDT9wWyEyrmyQ"

	s, err := ResolveToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func ResolveToken(Token string) (string, error) {
	token, err := jwt.Parse(Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("oauth"), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		if sub, ok := claims["sub"]; !ok {
			return "", errors.New("Token not include `cap` field.")
		} else {
			return sub.(string), nil
		}
	}
	return "", errors.New("Unknown error.")
}
