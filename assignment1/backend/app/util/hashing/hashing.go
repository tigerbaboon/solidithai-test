package hashing

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {

	r := bcrypt.DefaultCost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), r)
	log.Println(err)
	return bytes, err
}

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	log.Println(err)
	return err == nil
}
