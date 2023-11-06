package dtorepository

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

type ProductOrderRequest struct {
	ID              int
	Province        string
	District        string
	SubDistrict     string
	Kelurahan       string
	ZipCode         string
	AddressDetail   string
	AccountID       int
	SellerID        int
	CourierID       int
	Status          string
	Notes           string
	DeliveryFee     decimal.Decimal
	TotalAmount     decimal.Decimal
	ProductVariants []ProductOrderDetailRequest
}

type ReceiveOrderRequest struct {
	ID          int
	AccountID   int
	SellerID    int
	CourierID   int
	Status      string
	Notes       string
	TotalAmount decimal.Decimal
	Products    []model.ProductCombinationVariant
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

type ProductOrderSeller struct {
	ID              int
	ProductID       int
	SellerID        int
	IndividualPrice decimal.Decimal
	Quantity        int
	Status          string
}

type ProductOrderDetailRequest struct {
	ProductOrderID                       int
	ProductVariantSelectionCombinationID int
	Quantity                             int
	IndividualPrice                      decimal.Decimal
}

type ProductOrderDetailResponse struct {
	ID                                   int
	ProductOrderID                       int
	ProductVariantSelectionCombinationID int
	Quantity                             int
	IndividualPrice                      decimal.Decimal
	CreatedAt                            time.Time
	UpdatedAt                            time.Time
	DeletedAt                            time.Time
}
