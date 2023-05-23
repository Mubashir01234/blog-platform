package controller

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword generates a bcrypt hash of the given password.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash compares a password with a bcrypt hash and returns true if they match.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func roleValidator(role string) error {
	supportedRoles := []string{"Admin", "Author", "Reader"}
	for _, supportedRole := range supportedRoles {
		if role == supportedRole {
			return nil
		}
	}
	return fmt.Errorf("unsupported role: %s", role) // Role is not supported, return error
}
