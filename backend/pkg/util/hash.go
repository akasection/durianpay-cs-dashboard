package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenSaltAndHash(password string) ([]byte, error) {
	saltBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println("hashing error:", err)
	}
	return saltBytes, err
}

func MatchPassword(storedHash string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPassword))
	return err == nil
}
