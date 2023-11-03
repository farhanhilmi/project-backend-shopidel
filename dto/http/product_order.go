package dtohttp

import (
	"time"

	"github.com/shopspring/decimal"
)

type CanceledOrderRequest struct {
	ProductID int    `json:"product_id"`
	AccountID int    `json:"account_id"`
	Notes     string `json:"notes" binding:"required"`
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
	Notes         string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
