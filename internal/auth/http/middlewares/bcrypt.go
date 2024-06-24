package middlewares

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(passwors string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwors), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}
