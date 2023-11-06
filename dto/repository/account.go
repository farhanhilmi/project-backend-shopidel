package dtorepository

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

type EditAccountRequest struct {
	UserId         int
	FullName       string
	Username       string
	UsedEmail      string
	Email          string
	PhoneNumber    string
	Gender         string
	Birthdate      time.Time
	ProfilePicture string
}

type GetAccountRequest struct {
	UserId      int
	Email       string
	Username    string
	PhoneNumber string
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
	Password                string
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

type GetAccountCartItemsRequest struct {
	AccountId int
}

type GetAccountCartItemsResponse struct {
	CartItems []CartItem
}

type CartItem struct {
	ShopId       int
	ShopName     string
	ProductId    int
	ProductUrl   string
	ProductName  string
	ProductPrice decimal.Decimal
	Quantity     int
}
