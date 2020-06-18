package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u2 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u3, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s\n", u3)

	// convert type from uuid to string
	u4 := u3.String()
	fmt.Println(u4)
}
