package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
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


