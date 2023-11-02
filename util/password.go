package util

import (
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePassword(password string) bool {
	regex := `^[A-Za-z]{8,}$`

	re := regexp.MustCompile(regex)

	return re.MatchString(password) && strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
}
