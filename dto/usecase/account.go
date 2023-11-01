package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
	UserId   int
	Username string
	FullName string
	Email    string
	Password string
}

type CreateAccountResponse struct {
	Username string
	FullName string
	Email    string
}

type EditAccountRequest struct {
	UserId         int
	FullName       string
	Username       string
	Email          string
	PhoneNumber    string
	Gender         string
	Birthdate      time.Time
	ProfilePicture string
}

type AccountResponse struct {
	ID                      int
	FullName                string
	Username                string
	Email                   string
	PhoneNumber             string
	Password                string
	ShopName                string
	Gender                  string
	Birthdate               string
	ProfilePicture          string
	WalletNumber            string
	WalletPin               string
	Balance                 decimal.Decimal
	ForgetPasswordToken     string
	ForgetPasswordExpiredAt time.Time
}

type AccountRequest struct {
	ID                      int
	FullName                string
	Username                string
	Email                   string
	PhoneNumber             string
	Password                string
	ShopName                string
	Gender                  string
	Birthdate               string
	ProfilePicture          string
	WalletNumber            string
	WalletPin               string
	Balance                 decimal.Decimal
	ForgetPasswordToken     string
	ForgetPasswordExpiredAt time.Time
}

type CheckPasswordResponse struct {
	IsCorrect bool
}
type GetAccountRequest struct {
	UserId int
}

type GetAccountResponse struct {
	ID                      int
	FullName                string
	Username                string
	Email                   string
	PhoneNumber             string
	ShopName                string
	Gender                  string
	Birthdate               time.Time
	ProfilePicture          string
	WalletNumber            string
	WalletPin               string
	Balance                 decimal.Decimal
	ForgetPasswordToken     string
	ForgetPasswordExpiredAt time.Time
}


type EditAccountResponse struct {
	ID             int    
	FullName       string    
	Username       string    
	Email          string    
	PhoneNumber    string    
	Gender         string    
	Birthdate      time.Time 
	ProfilePicture string    
}