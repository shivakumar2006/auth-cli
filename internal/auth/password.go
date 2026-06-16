package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	// the default cost of bcrypt is 10
	BcryptCost = bcrypt.DefaultCost
)

var (
	ErrInvalidPassword = errors.New("invalid username or password")
)

// it validated the password needs to contains uppercase, lowercase letters and digits
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var (
		hasUpper bool
		hasLower bool
		hasDigit bool
	)

	for _, ch := range password {
		switch {
		case ch >= 'A' && ch <= 'Z':
			hasUpper = true
		case ch >= 'a' && ch <= 'z':
			hasLower = true
		case ch >= '0' && ch <= '9':
			hasDigit = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit {
		return errors.New(
			"password must contain uppercase, lowercase and digit",
		)
	}

	return nil
}

// HashPassword function hashes a plain text password bedore storing it in the database
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

// this function compares a bcrypt hash with a plain text password
func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return ErrInvalidPassword
	}
	return nil
}
