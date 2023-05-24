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

// roleValidator is used to validate the user role and only accept the specified role.
func roleValidator(role string) error {
	supportedRoles := []string{"Admin", "Author", "Reader"} // Roles that are supported
	for _, supportedRole := range supportedRoles {          // Check and validate the role is supported or not
		if role == supportedRole { // If the role is supported, return nil
			return nil
		}
	}
	return fmt.Errorf("unsupported role: %s", role) // Role is not supported, return error
}
