package dtohttp

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateShopPromotionRequest struct {
	Name              string          `json:"name" binding:"required"`
	Quota             int             `json:"quota" binding:"required"`
	StartDate         time.Time       `json:"start_date" binding:"required"`
	EndDate           time.Time       `json:"end_date" binding:"required"`
	MinPurchaseAmount decimal.Decimal `json:"min_purchase_amount" binding:"required"`
	MaxPurchaseAmount decimal.Decimal `json:"max_purchase_amount" binding:"required"`
	DiscountAmount    decimal.Decimal `json:"discount_amount" binding:"required"`
}

type UpdateShopPromotionRequest struct {
	Name              string          `json:"name" binding:"required"`
	Quota             int             `json:"quota" binding:"required"`
	StartDate         time.Time       `json:"start_date" binding:"required"`
	EndDate           time.Time       `json:"end_date" binding:"required"`
	MinPurchaseAmount decimal.Decimal `json:"min_purchase_amount" binding:"required"`
	MaxPurchaseAmount decimal.Decimal `json:"max_purchase_amount" binding:"required"`
	DiscountAmount    decimal.Decimal `json:"discount_amount" binding:"required"`
}
