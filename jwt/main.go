package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	payload := &TokenPayload{
		ID:      233,
		Expired: time.Hour * 24 * 30,
	}
	token, err := GenerateToken(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	id, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

var (
	jwtKey = ""
)

type TokenPayload struct {
	ID      uint32        `json:"id"`
	Expired time.Duration `json:"expired"`
}

type TokenResolve struct {
	ID      uint32 `json:"id"`
	Expired int64  `json:"expired"`
}

type TokenClaims struct {
	ID uint32 `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(payload *TokenPayload) (string, error) {
	claims := &TokenClaims{
		ID: payload.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func ParseToken(token string) (uint32, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		fmt.Println(claims)
		sub, ok := claims["id"]
		if !ok {
			return 0, errors.New("Token not include `cap` field.")
		}
		return uint32(sub.(float64)), nil
	}
	return 0, errors.New("Unknown error.")
}
