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

type LoginRequest struct {
	Email    string
	Password string
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

type LoginResponse struct {
	AccessToken string
}

type GetCartRequest struct {
	UserId int
}

type GetCartResponse struct {
	CartShops []CartShop `json:"cart_shops"`
}

type CartShop struct {
	ShopId    int        `json:"shop_id"`
	ShopName  string     `json:"shop_name"`
	CartItems []CartItem `json:"cart_items"`
}

type CartItem struct {
	ProductId         int             `json:"product_id"`
	ProductImageUrl   string          `json:"product_image_url"`
	ProductName       string          `json:"product_name"`
	ProductUnitPrice  decimal.Decimal `json:"product_unit_price"`
	ProductQuantity   int             `json:"product_quantity"`
	ProductTotalPrice decimal.Decimal `json:"product_total_price"`
}
