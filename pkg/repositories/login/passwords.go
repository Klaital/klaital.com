package login_repository

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(raw string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
