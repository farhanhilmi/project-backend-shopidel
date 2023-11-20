package dtousecase

import (
	"mime/multipart"
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
	AccountID            int
	ProductID            int
	ProductOrderDetailID int
	Feedback             string
	Rating               int
	Image                multipart.File
	ImageHeader          *multipart.FileHeader
}

type AddProductReviewResponse struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	ProductID int       `json:"product_id"`
	OrderID   int       `json:"order_id"`
	Feedback  string    `json:"feedback"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductSellerOrderHistoryRequest struct {
	AccountID int
	Status    string
	SortBy    string
	Sort      string
	Limit     int
	Page      int
	StartDate string
	EndDate   string
}

type OrdersResponse struct {
	OrderID      int             `json:"order_id"`
	ShopName     string          `json:"shop_name"`
	Status       string          `json:"status"`
	Products     []OrderProduct  `json:"products"`
	TotalPayment decimal.Decimal `json:"total_payment"`
	CreateAt     string          `json:"created_at"`
}

type SellerOrdersResponse struct {
	OrderID      int                  `json:"order_id"`
	BuyerName    string               `json:"buyer_name"`
	Status       string               `json:"status"`
	Products     []SellerOrderProduct `json:"products"`
	TotalPayment decimal.Decimal      `json:"total_payment"`
	CreateAt     string               `json:"created_at"`
}
