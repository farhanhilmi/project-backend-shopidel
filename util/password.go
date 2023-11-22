package util

import (
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
	return strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
}

func CheckPasswordIdentical(password, target string) bool {
	pattern := []byte(strings.ToLower(password))
	text := []byte(strings.ToLower(target))
	passwordLen := len(password)
	targetLen := len(target)
	var i, j, p, t int
	hash, totalExtendedASCII, prime_number := 1, 256, 101

	for i = 0; i < passwordLen-1; i++ {
		hash = (hash * totalExtendedASCII) % prime_number
	}

	for i = 0; i < passwordLen; i++ {
		p = (totalExtendedASCII*p + int(pattern[i])) % prime_number
		t = (totalExtendedASCII*t + int(text[i])) % prime_number
	}

	for i = 0; i <= targetLen-passwordLen; i++ {
		if p == t {
			for j = 0; j < passwordLen; j++ {
				if target[i+j] != target[j] {
					break
				}
				j += 1
			}

			if j == passwordLen {
				return true
			}
		}

		if i < targetLen-passwordLen {
			t = (totalExtendedASCII*
				(t-int(target[i])*hash) +
				int(target[i+passwordLen])) % prime_number

			if t < 0 {
				t = (t + prime_number)
			}
		}
	}

	return false
}
