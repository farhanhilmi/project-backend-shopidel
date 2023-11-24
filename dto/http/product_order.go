package dtohttp

import (
	"time"

	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"github.com/shopspring/decimal"
)

type CanceledOrderRequest struct {
	ProductID int    `json:"product_id"`
	AccountID int    `json:"account_id"`
	Notes     string `json:"notes" binding:"required"`
}

type CompletedOrderBuyerRequest struct {
	OrderID int `json:"order_id" binding:"required"`
}

type ProductOrderResponse struct {
	ID            int             `json:"id,omitempty"`
	CourierID     int             `json:"courier_id,omitempty"`
	AccountID     int             `json:"account_id,omitempty"`
	DeliveryFee   decimal.Decimal `json:"delivery_id,omitempty"`
	Province      string          `json:"province,omitempty"`
	District      string          `json:"district,omitempty"`
	SubDistrict   string          `json:"sub_district,omitempty"`
	Kelurahan     string          `json:"kelurahan,omitempty"`
	ZipCode       string          `json:"zip_code,omitempty"`
	AddressDetail string          `json:"address_detail,omitempty"`
	Notes         string          `json:"notes,omitempty"`
	Status        string          `json:"status,omitempty"`
	CreatedAt     time.Time       `json:"created_at,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty"`
	DeletedAt     time.Time       `json:"deleted_at,omitempty"`
}

type ProductOrderReceiveResponse struct {
	ID     int    `json:"id,omitempty"`
	Notes  string `json:"notes,omitempty"`
	Status string `json:"status,omitempty"`
}

type CheckoutOrderRequest struct {
	SellerID               int                              `json:"seller_id" binding:"required"`
	ProductVariant         []dtousecase.ProductVariantOrder `json:"product_variant" binding:"required"`
	DestinationAddressID   string                           `json:"destination_address_id" binding:"required"`
	VoucherID              int                              `json:"voucher_id"`
	Notes                  string                           `json:"notes"`
	Weight                 string                           `json:"weight" binding:"required"`
	CourierID              int                              `json:"courier_id" binding:"required"`
	ShopPromotionId        int                              `json:"shop_promotion_id"`
	MarketplacePromotionId int                              `json:"marketplace_promotion_id"`
}

type AddProductReviewRequest struct {
	AccountID int
	OrderID   int
	ProductID int    `form:"product_id" binding:"required"`
	Feedback  string `form:"feedback" binding:"required"`
	Rating    int    `form:"rating" binding:"required"`
}
