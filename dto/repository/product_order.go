package dtorepository

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

type ProductOrderRequest struct {
	ID                 int
	Province           string
	District           string
	SubDistrict        string
	Kelurahan          string
	ZipCode            string
	AddressDetail      string
	AccountID          int
	SellerID           int
	CourierName        string
	Status             string
	Notes              string
	SellerWalletNumber string
	BuyerWalletNumber  string
	DeliveryFee        decimal.Decimal
	TotalAmount        decimal.Decimal
	TotalSellerAmount  decimal.Decimal
	ProductVariants    []ProductOrderDetailRequest
}

type ReceiveOrderRequest struct {
	ID                 int
	AccountID          int
	SellerID           int
	CourierID          int
	Status             string
	Notes              string
	TotalAmount        decimal.Decimal
	Products           []model.ProductCombinationVariant
	SellerWalletNumber string
}

type ProductOrderResponse struct {
	ID            int
	CourierName   string
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

type CourierData struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type SellerCourier struct {
	SellerID int
}

type SellerCourierResponse struct {
	CourierID int
}

type ProductOrderHistoryRequest struct {
	AccountID int
	Status    string
	SortBy    string
	Sort      string
	Limit     int
	Page      int
	StartDate string
	EndDate   string
}

type OrderDetailRequest struct {
	AccountID int
	OrderID   int
}

type AddProductReviewRequest struct {
	AccountID            int
	ProductID            int
	ProductOrderDetailID int
	Feedback             string
	ImageURL             string
	Rating               int
}

type AddProductReviewResponse struct {
	ID        int
	AccountID int
	ProductID int
	OrderID   int
	Feedback  string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ProductReviewRequest struct {
	AccountID int
	ProductID int
	OrderID   int
}

type ProductReviewResponse struct {
	ID        int
	AccountID int
	ProductID int
	OrderID   int
}
