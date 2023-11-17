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

type ForgetPasswordRequest struct {
	Email string
}

type ForgetChangePasswordRequest struct {
	Token    string
	Password string
}

type SendEmailPayload struct {
	RecipientName  string
	RecipientEmail string
	Token          string
	ExpiresAt      time.Time
}

type RefreshTokenRequest struct {
	RefreshToken string
	UserId       int
}

type AddressRequest struct {
	UserId int
}

type GetAccountResponse struct {
	ID                  int
	FullName            string
	Username            string
	Email               string
	PhoneNumber         string
	ShopName            string
	Gender              string
	Birthdate           time.Time
	ProfilePicture      string
	WalletNumber        string
	WalletPin           string
	Balance             decimal.Decimal
	ForgetPasswordToken string
	IsSeller            bool
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
	AccessToken  string
	RefreshToken string
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

type AddProductToCartRequest struct {
	UserId           int
	ProductVariantId int
	Quantity         int
}

type AddProductToCartResponse struct {
	ProductId int
	Quantity  int
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

type RegisterAddressRequest struct {
	AccountId   int
	ProvinceId  int
	DistrictId  int
	SubDistrict string
	Kelurahan   string
	ZipCode     string
	Detail      string
}

type RegisterAddressResponse struct {
	AccountId       int
	ProvinceId      int
	DistrictId      int
	SubDistrict     string
	Kelurahan       string
	ZipCode         string
	Detail          string
	IsBuyerDefault  bool
	IsSellerDefault bool
}

type UpdateAddressRequest struct {
	AddressId       int
	AccountId       int
	ProvinceId      int
	DistrictId      int
	SubDistrict     string
	Kelurahan       string
	ZipCode         string
	Detail          string
	IsBuyerDefault  bool
	IsSellerDefault bool
}

type UpdateAddressResponse struct {
	AccountId       int
	ProvinceId      int
	DistrictId      int
	SubDistrict     string
	Kelurahan       string
	ZipCode         string
	Detail          string
	IsBuyerDefault  bool
	IsSellerDefault bool
}

type GetProvincesResponse struct {
	Provinces []Province `json:"provinces"`
}

type GetDistrictsResponse struct {
	Districts []District `json:"districts"`
}

type Province struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetDistrictRequest struct {
	ProvinceId int
}

type GetDistrictResponse struct {
	Districts []District `json:"districts"`
}

type District struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteAddressRequest struct {
	AddressId int
}

type ChangePasswordRequest struct {
	AccountID   int
	OldPassword string
	NewPassword string
}

type GetCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Children []Category `json:"children,omitempty"`
}
