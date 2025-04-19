package helper

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
