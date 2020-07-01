package main

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"

func main() {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhODUwZGE2NC0zMTBlLTQxNmYtYTZjMy1hZDlhN2FkN2ViMjUiLCJleHAiOjE1OTM1OTc1ODIsInN1YiI6IjI1NCJ9.Bnqp9hTWLudtjsGP2eZmWSWiuwKBsdkrAuGxKJxdgyFOgKuPjwRFKWP3gXR2_aFb-ezNB0o9sxYnuzwnDm8A7Q"

	s, err := ResolveToken(token)
	if err != nil {
		if IsErrTokenExpired(err) {
			fmt.Println("token is expired.")
			return
		}
		panic(err)
	}
	fmt.Println(s)
}

func ResolveToken(Token string) (string, error) {
	token, err := jwt.Parse(Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		sub, ok := claims["sub"]
		if !ok {
			return "", errors.New("Token not include `cap` field.")
		}
		return sub.(string), nil
	}
	return "", errors.New("Unknown error.")
}

func IsErrTokenExpired(err error) bool {
	return err.Error() == "Token is expired"
}
