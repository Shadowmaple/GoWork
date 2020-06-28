package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	clientID, secret := generateClientIDAndSecret()
	fmt.Println(clientID, secret)
}

func generateClientIDAndSecret() (string, string) {
	clientID := generateUUID()

	Secret := generateUUID()

	return clientID, Secret
}

func generateUUID() string {
	return uuid.NewV4().String()
}
