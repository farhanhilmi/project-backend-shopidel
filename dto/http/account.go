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

type EditAccountRequest struct {
	FullName       string    `json:"full_name" binding:"required"`
	Username       string    `json:"username" binding:"alphanum"`
	Email          string    `json:"email" binding:"email"`
	PhoneNumber    string    `json:"phone_number" binding:"e164"`
	Gender         string    `json:"gender" binding:"lowercase"`
	Birthdate      time.Time `json:"birthdate"`
	ProfilePicture string    `json:"profile_picture"`
}

type GetAccountRequest struct {
	UserId int `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type AddressRequest struct {
	UserId int `json:"id"`
}

type CreateAddressRequest struct {
	AccountID            int
	Province             string
	District             string
	RajaOngkirDistrictId string
	SubDistrict          string
	Kelurahan            string
	ZipCode              string
	Detail               string
	IsBuyerDefault       bool
	IsSellerDefault      bool
}

type GetAccountResponse struct {
	ID                      int             `json:"id"`
	FullName                string          `json:"full_name,omitempty"`
	Username                string          `json:"username,omitempty"`
	Email                   string          `json:"email,omitempty"`
	PhoneNumber             string          `json:"phone_number,omitempty"`
	ShopName                string          `json:"shop_name,omitempty"`
	Gender                  string          `json:"gender,omitempty"`
	Birthdate               time.Time       `json:"birthdate,omitempty"`
	ProfilePicture          string          `json:"profile_picture"`
	WalletNumber            string          `json:"wallet_number,omitempty"`
	WalletPin               string          `json:"wallet_pin,omitempty"`
	Balance                 decimal.Decimal `json:"balance,omitempty"`
	ForgetPasswordToken     string          `json:"forget_password_token,omitempty"`
	ForgetPasswordExpiredAt time.Time       `json:"forget_password_expired_at,omitempty"`
}

type EditAccountResponse struct {
	ID             int       `json:"id"`
	FullName       string    `json:"full_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Gender         string    `json:"gender"`
	Birthdate      time.Time `json:"birthdate"`
	ProfilePicture string    `json:"profile_picture"`
}

type CheckPasswordResponse struct {
	IsCorrect bool `json:"isCorrect"`
}

type AddProductToCartRequest struct {
	ProductId int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type AddProductToCartResponse struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type AddressResponse struct {
	ID              int    `json:"id"`
	FullAddress     string `json:"full_address"`
	Detail          string `json:"detail"`
	ZipCode         string `json:"zip_code"`
	Kelurahan       string `json:"kelurahan"`
	SubDistrict     string `json:"sub_district"`
	District        string `json:"district"`
	Province        string `json:"province"`
	IsBuyerDefault  bool   `json:"is_buyer_default"`
	IsSellerDefault bool   `json:"is_seller_default"`
}

type RegisterSellerRequest struct {
	ShopName      string `json:"shop_name"`
	AddressId     int    `json:"address_id"`
	ListCourierId []int  `json:"list_courier_id"`
}

type RegisterSellerResponse struct {
	ShopName string
}

type DeleteCartProductRequest struct {
	ListProductID []int `json:"list_product_id" binding:"required"`
}

type RegisterAddressRequest struct {
	ProvinceId  int    `json:"province_id" binding:"required"`
	DistrictId  int    `json:"district_id" binding:"required"`
	SubDistrict string `json:"sub_district" binding:"required"`
	Kelurahan   string `json:"kelurahan" binding:"required"`
	ZipCode     string `json:"zip_code" binding:"required"`
	Detail      string `json:"detail" binding:"required"`
}

type RegisterAddressResponse struct {
	ProvinceId  int    `json:"province_id" binding:"required"`
	DistrictId  int    `json:"district_id" binding:"required"`
	SubDistrict string `json:"sub_district"`
	Kelurahan   string `json:"kelurahan"`
	ZipCode     string `json:"zip_code"`
	Detail      string `json:"detail"`
}

type UpdateAddressRequest struct {
	ProvinceId  int    `json:"province_id" binding:"required"`
	DistrictId  int    `json:"district_id" binding:"required"`
	SubDistrict string `json:"sub_district" binding:"required"`
	Kelurahan   string `json:"kelurahan" binding:"required"`
	ZipCode     string `json:"zip_code" binding:"required"`
	Detail      string `json:"detail" binding:"required"`
}

type UpdateAddressResponse struct {
	ProvinceId  int    `json:"province_id" binding:"required"`
	DistrictId  int    `json:"district_id" binding:"required"`
	SubDistrict string `json:"sub_district"`
	Kelurahan   string `json:"kelurahan"`
	ZipCode     string `json:"zip_code"`
	Detail      string `json:"detail"`
}
