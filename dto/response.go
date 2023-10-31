package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type ErrResponse struct {
	Error string `json:"error"`
}

type JSONResponse struct {
	Data        any    `json:"data,omitempty"`
	Message     string `json:"message,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type AccountResponse struct {
	ID                      int             `json:"id"`
	FullName                string          `json:"full_name,omitempty"`
	Username                string          `json:"username,omitempty"`
	Email                   string          `json:"email,omitempty"`
	PhoneNumber             string          `json:"phone_number,omitempty"`
	ShopName                string          `json:"shop_name,omitempty"`
	Gender                  string          `json:"gender,omitempty"`
	Birthdate               string          `json:"birthdate,omitempty"`
	ProfilePicture          string          `json:"profile_picture,omitempty"`
	WalletNumber            string          `json:"wallet_number,omitempty"`
	WalletPin               string          `json:"wallet_pin,omitempty"`
	Balance                 decimal.Decimal `json:"balance,omitempty"`
	ForgetPasswordToken     string          `json:"forget_password_token,omitempty"`
	ForgetPasswordExpiredAt time.Time       `json:"forget_password_expired_at,omitempty"`
}