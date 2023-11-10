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

type AddressRequest struct {
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

type AddProductToCartRequest struct {
	AccountId                   int
	ProductVariantCombinationId int
	Quantity                    int
}

type AddProductToCartResponse struct {
	AccountId                   int
	ProductVariantCombinationId int
	Quantity                    int
}

type AddressResponse struct {
	ID              int
	FullAddress     string
	Detail          string
	ZipCode         string
	Kelurahan       string
	SubDistrict     string
	DistrictId      int
	District        string
	ProvinceId      int
	Province        string
	IsBuyerDefault  bool
	IsSellerDefault bool
}

type RegisterSellerRequest struct {
	UserId        int
	ShopName      string
	AddressId     int
	ListCourierId []int
}

type RegisterSellerResponse struct {
	ShopName string
}

type DeleteCartProductRequest struct {
	ListProductID []int
}

type DeleteCartProductResponse struct {
	ListProductID []int
}

type DeleteAddressRequest struct {
	AddressId int
}

type DeleteAddressResponse struct {
	AddressId int
}
type SellerDataRequest struct {
	SellerId int
}

type SellerDataResponse struct {
	Id                  int
	Name                string
	ProfilePicture      string
	District            string
	StartOperatingHours string
	EndOperatingHours   string
	TimeZone            string
}

type FindSellerProductsRequest struct {
	SellerId int
}

type FindSellerProductsResponse struct {
	Products []SellerProduct
}

type SellerProduct struct {
	Name           string
	Price          decimal.Decimal
	PictureUrl     string
	Stars          decimal.Decimal
	TotalSold      int
	CreatedAt      string
	CategoryLevel1 string
	CategoryLevel2 string
	CategoryLevel3 string
}
