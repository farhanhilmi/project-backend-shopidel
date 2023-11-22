package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type GetShopAvailablePromotionsRequest struct {
	ShopId int
}

type GetShopAvailablePromotionsResponse struct {
	ShopPromotions []ShopPromotion
}

type ShopPromotion struct {
	ID                 int             `json:"id"`
	Name               string          `json:"name"`
	MinPurchaseAmount  decimal.Decimal `json:"min_purchase_amount"`
	MaxPurchaseAmount  decimal.Decimal `json:"max_purchase_amount"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage"`
	SelectedProductsId []int           `json:"selected_products,omitempty"`
	Quota              int             `json:"quota,omitempty"`
	TotalUsed          int             `json:"total_used"`
	StartDate          time.Time       `json:"start_date,omitempty"`
	EndDate            time.Time       `json:"end_date,omitempty"`
}

type GetMarketplacePromotionsResponse struct {
	MarketplacePromotions []MarketplacePromotion
}

type MarketplacePromotion struct {
	ID                 int             `json:"id"`
	Name               string          `json:"name"`
	MinPurchaseAmount  decimal.Decimal `json:"min_purchase_amount"`
	MaxPurchaseAmount  decimal.Decimal `json:"max_purchase_amount"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage"`
}

type GetShopPromotionsRequest struct {
	ShopId int
	Page   int
}

type GetShopPromotionsResponse struct {
	ShopPromotions []ShopPromotion
	CurrentPage    int
	TotalPages     int
	TotalItems     int
	Limit          int
}

type CreateShopPromotionRequest struct {
	ShopId             int
	Name               string
	Quota              int
	StartDate          time.Time
	EndDate            time.Time
	MinPurchaseAmount  decimal.Decimal
	MaxPurchaseAmount  decimal.Decimal
	DiscountPercentage decimal.Decimal
	SelectedProductsId []int
}

type CreateShopPromotionResponse struct {
	Id                 int             `json:"id"`
	ShopId             int             `json:"shop_id"`
	Name               string          `json:"name"`
	Quota              int             `json:"quota"`
	StartDate          time.Time       `json:"start_date"`
	EndDate            time.Time       `json:"end_date"`
	MinPurchaseAmount  decimal.Decimal `json:"min_purchase_amount"`
	MaxPurchaseAmount  decimal.Decimal `json:"max_purchase_amount"`
	DiscountPercentage decimal.Decimal `json:"disount_percentage"`
	SelectedProductsId []int           `json:"selected_products_id"`
	CreatedAt          time.Time       `json:"created_at"`
}

type GetShopPromotionDetailResponse struct {
	ID                 int                            `json:"id"`
	Name               string                         `json:"name"`
	MinPurchaseAmount  decimal.Decimal                `json:"min_purchase_amount"`
	MaxPurchaseAmount  decimal.Decimal                `json:"max_purchase_amount"`
	DiscountPercentage decimal.Decimal                `json:"discount_percentage"`
	Quota              int                            `json:"quota,omitempty"`
	TotalUsed          int                            `json:"total_used,omitempty"`
	StartDate          time.Time                      `json:"start_date,omitempty"`
	EndDate            time.Time                      `json:"end_date,omitempty"`
	SelectedProducts   []ShopPromotionSelectedProduct `json:"selected_products"`
	CreatedAt          time.Time                      `json:"created_at"`
}

type ShopPromotionSelectedProduct struct {
	ProductId   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateShopPromotionRequest struct {
	Id                 int
	ShopId             int
	Name               string
	Quota              int
	StartDate          time.Time
	EndDate            time.Time
	MinPurchaseAmount  decimal.Decimal
	MaxPurchaseAmount  decimal.Decimal
	DiscountPercentage decimal.Decimal
	SelectedProductsId []int
}

type UpdateShopPromotionResponse struct {
	Id                 int             `json:"id"`
	ShopId             int             `json:"shop_id"`
	Name               string          `json:"name"`
	Quota              int             `json:"quota"`
	StartDate          time.Time       `json:"start_date"`
	EndDate            time.Time       `json:"end_date"`
	MinPurchaseAmount  decimal.Decimal `json:"min_purchase_amount"`
	MaxPurchaseAmount  decimal.Decimal `json:"max_purchase_amount"`
	DiscountPercentage decimal.Decimal `json:"disount_percentage"`
	SelectedProductsId []int           `json:"selected_products_id"`
	CreatedAt          time.Time       `json:"created_at"`
}
