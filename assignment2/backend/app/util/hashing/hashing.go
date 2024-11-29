package hashing

import (
	"app/config/log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {

	r := bcrypt.DefaultCost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), r)
	return bytes, err
}

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {
		log.Error(err.Error())
		return false
	}

	return true
}
