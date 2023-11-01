package dtohttp

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateAccountResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type CheckPasswordRequest struct {
	Password string `json:"password"`
}

type GetAccountRequest struct {
	UserId int `json:"id"`
}

type GetAccountResponse struct {
	ID                      int             `json:"id,omitempty"`
	FullName                string          `json:"full_name,omitempty"`
	Username                string          `json:"username,omitempty"`
	Email                   string          `json:"email,omitempty"`
	PhoneNumber             string          `json:"phone_number,omitempty"`
	ShopName                string          `json:"shop_name,omitempty"`
	Gender                  string          `json:"gender,omitempty"`
	Birthdate               time.Time       `json:"birthdate,omitempty"`
	ProfilePicture          string          `json:"profile_picture,omitempty"`
	WalletNumber            string          `json:"wallet_number,omitempty"`
	WalletPin               string          `json:"wallet_pin,omitempty"`
	Balance                 decimal.Decimal `json:"balance,omitempty"`
	ForgetPasswordToken     string          `json:"forget_password_token,omitempty"`
	ForgetPasswordExpiredAt time.Time       `json:"forget_password_expired_at,omitempty"`
}

type CheckPasswordResponse struct {
	IsCorrect bool `json:"isCorrect"`
}
