package util

import (
	"math/rand"
	"encoding/base64"
	"fmt"

	"strconv"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userId int, role, walletNumber string) (string, error) {
	tokenDuration, err := strconv.Atoi(config.GetEnv("JWT_DURATION"))
	if err != nil {
		return "", err
	}

	claims := dtogeneral.ClaimsJWT{
		UserId:       userId,
		Role:         role,
		WalletNumber: walletNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(tokenDuration))},
			Issuer:    config.GetEnv("JWT_ISSUER"),
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetEnv("JWT_SECRET_KEY")))
}

func GenerateRefreshJWT(userId int) (string, error) {
	tokenDuration, err := strconv.Atoi(config.GetEnv("JWT_REFRESH_DURATION"))
	if err != nil {
		return "", err
	}

	claims := dtogeneral.ClaimsJWT{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(tokenDuration))},
			Issuer:    config.GetEnv("JWT_REFRESH_ISSUER"),
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetEnv("JWT_REFRESH_SECRET_KEY")))
}

func ValidateRefreshToken(tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &dtogeneral.ClaimsJWT{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.GetEnv("JWT_REFRESH_SECRET_KEY")), nil
	})
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &dtogeneral.ClaimsJWT{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.GetEnv("JWT_SECRET_KEY")), nil
	})
}

func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GenerateRandomOTP() (string, error) {
	baseOTP := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	OTPLength := 6
	OTP := make([]byte, OTPLength)

	for index := range OTP {
		OTP[index] = baseOTP[rand.Intn(len(baseOTP))]
	}
	
	return string(OTP), nil
}
