package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductOrderRequest struct {
	ID                 int
	AccountID          int
	CourierID          int
	SellerID           int
	Notes              string
	Status             string
	SellerWalletNumber string
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
	DestinationAddressID string
	VoucherID            int
	UserID               int
	CourierID            int
	SellerID             int
	Notes                string
	Weight               string
}

type CheckoutOrderResponse struct {
	ID                   int
	ProductVariant       []ProductVariantOrder
	DestinationAddressID int
	VoucherID            int
	CourierID            int
	UserID               int
}

type CheckDeliveryFeeRequest struct {
	Origin      string
	ID          int    `json:"id"`
	SellerID    int    `json:"seller_id" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	Weight      string `json:"weight" binding:"required"`
	Courier     string `json:"courier" binding:"required"`
}

type CourierFeeResponse struct {
	Cost      int    `json:"cost"`
	Estimated string `json:"estimated"`
	Note      string `json:"note"`
}

type SellerCourier struct {
	SellerID int
}

type AddProductReview struct {
	AccountID int
	ProductID int
	OrderID   int
	Feedback  string
	Rating    int
}

type AddProductReviewResponse struct {
	ID        int
	AccountID int
	ProductID int
	OrderID   int
	Feedback  string
	Rating    int
	CreatedAt time.Time
}
