package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// EncodePassword ...
func EncodePassword(password string) (string, error) {
	bytePassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}
