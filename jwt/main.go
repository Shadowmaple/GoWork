package main

import (
	"fmt"
	"time"
)

func main() {
	day := 30
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

	data, err := ResolveToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
