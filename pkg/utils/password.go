package utils

import "golang.org/x/crypto/bcrypt"

const passwordCost = 14

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	return string(hash), err
}

func CompareHashPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
