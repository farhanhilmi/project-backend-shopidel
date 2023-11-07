package dtogeneral

import "github.com/golang-jwt/jwt/v4"

type ClaimsJWT struct {
	UserId       int    `json:"user_id"`
	Role         string `json:"role"`
	WalletNumber string `json:"wallet_number"`
	jwt.RegisteredClaims
}
