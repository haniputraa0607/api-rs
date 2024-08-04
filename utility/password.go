package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash)
}

func CheckPassword(inputPassword, hashedPassword string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return e == nil
}
