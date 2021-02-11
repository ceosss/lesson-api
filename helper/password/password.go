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

// DecodePassword ...
func DecodePassword(hashedPassword string, password string) bool {

	byteHashedPassword := []byte(hashedPassword)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return false
	}

	return true

}
