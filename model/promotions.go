package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ShopPromotion struct {
	ID                 int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	ShopId             int             `gorm:"not null"`
	Name               string          `gorm:"not null"`
	Quota              int             `gorm:"not null"`
	TotalUsed          int             `gorm:"not null;default:0"`
	StartDate          time.Time       `gorm:"not null"`
	EndDate            time.Time       `gorm:"not null"`
	MinPurchaseAmount  decimal.Decimal `gorm:"not null"`
	MaxPurchaseAmount  decimal.Decimal `gorm:"not null"`
	DiscountPercentage decimal.Decimal `gorm:"not null"`
	SelectedProducts   []ShopPromotionSelectedProduct
	CreatedAt          time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt          time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt          time.Time `gorm:"type:timestamp;default:null"`
}

type ShopPromotionSelectedProduct struct {
	ID              int `gorm:"primaryKey;not null,autoIncrement;serial"`
	ShopPromotionId int `gorm:"not null"`
	ShopPromotion   ShopPromotion
	ProductId       int `gorm:"not null"`
	Product         Products
	CreatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
}

type MarketplacePromotion struct {
	ID                 int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name               string          `gorm:"not null"`
	Quota              int             `gorm:"not null"`
	TotalUsed          int             `gorm:"not null;default:0"`
	StartDate          time.Time       `gorm:"not null"`
	EndDate            time.Time       `gorm:"not null"`
	MinPurchaseAmount  decimal.Decimal `gorm:"not null"`
	MaxPurchaseAmount  decimal.Decimal `gorm:"not null"`
	DiscountPercentage decimal.Decimal `gorm:"not null"`
	CreatedAt          time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt          time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt          time.Time       `gorm:"type:timestamp;default:null"`
}
