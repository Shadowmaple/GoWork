package main

import (
	"fmt"
	"time"
)

func main() {
	day := 0
	payload := &TokenPayload{
		ID:      233,
		Role:    3,
		TeamID:  1,
		Expired: time.Hour * 24 * time.Duration(day),
	}
	token, err := GenerateToken(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NzMsInJvbGUiOjMsInRlYW1faWQiOjEsImV4cGlyZXNfYXQiOjE2MDc2OTU5MjN9.hHNxUuAELgQJegjhoxYceavacL9GX4zjBWddxviFn6Q"

	data, err := ResolveToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
