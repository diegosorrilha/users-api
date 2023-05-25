// Package crypt collects encryption functions.
package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is a function to encrypt a password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
