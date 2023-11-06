package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductOrderRequest struct {
	ID        int
	AccountID int
	CourierID int
	SellerID  int
	Notes     string
	Status    string
}

type ProductOrderResponse struct {
	ID            int
	CourierID     int
	AccountID     int
	DeliveryFee   decimal.Decimal
	Province      string
	District      string
	SubDistrict   string
	Kelurahan     string
	ZipCode       string
	AddressDetail string
	Status        string
	Notes         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type ProductVariantOrder struct {
	ID       int `json:"id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type CheckoutOrderRequest struct {
	ProductVariant       []ProductVariantOrder
	DestinationAddressID int
	VoucherID            int
	UserID               int
	CourierID            int
	SellerID             int
	Notes                string
}

type CheckoutOrderResponse struct {
	ID                   int
	ProductVariant       []ProductVariantOrder
	DestinationAddressID int
	VoucherID            int
	CourierID            int
	UserID               int
}
