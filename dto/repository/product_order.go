package dtorepository

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductOrderRequest struct {
	ID        int
	AccountID int
	SellerID  int
	CourierID int
	Status    string
	Notes     string
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
