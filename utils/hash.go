package utils

import "golang.org/x/crypto/bcrypt"

func CreateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash) , err
}

func VerifyPassword(Hashpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Hashpassword), []byte(password))
	return err == nil
}