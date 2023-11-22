package seeder

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ShopPromotions = []*model.ShopPromotion{
	{
		ID:                 1,
		ShopId:             4,
		Name:               "Kalung dan sandal discount",
		Quota:              100,
		StartDate:          time.Date(2020, 01, 01, 01, 01, 01, 01, time.Local),
		EndDate:            time.Date(2024, 01, 01, 01, 01, 01, 01, time.Local),
		MinPurchaseAmount:  decimal.NewFromFloat(10000),
		MaxPurchaseAmount:  decimal.NewFromFloat(200000),
		DiscountPercentage: decimal.NewFromFloat(7.5),
	},
	{
		ID:                 2,
		ShopId:             4,
		Name:               "Sandal dan ikat pinggang discount",
		Quota:              100,
		StartDate:          time.Date(2020, 01, 01, 01, 01, 01, 01, time.Local),
		EndDate:            time.Date(2024, 01, 01, 01, 01, 01, 01, time.Local),
		MinPurchaseAmount:  decimal.NewFromFloat(10000),
		MaxPurchaseAmount:  decimal.NewFromFloat(200000),
		DiscountPercentage: decimal.NewFromFloat(12.5),
	},
}

var ShopPromotionSelectedProducts = []*model.ShopPromotionSelectedProduct{
	{
		ShopPromotionId: 1,
		ProductId:       4,
	},
	{
		ShopPromotionId: 1,
		ProductId:       5,
	},
	{
		ShopPromotionId: 2,
		ProductId:       5,
	},
	{
		ShopPromotionId: 2,
		ProductId:       6,
	},
}

var MarketplacePromotions = []model.MarketplacePromotion{
	{
		ID:                 2,
		Name:               "Discount meriah November",
		Quota:              100,
		StartDate:          time.Date(2020, 01, 01, 01, 01, 01, 01, time.Local),
		EndDate:            time.Date(2024, 01, 01, 01, 01, 01, 01, time.Local),
		MinPurchaseAmount:  decimal.NewFromFloat(10000),
		MaxPurchaseAmount:  decimal.NewFromFloat(200000),
		DiscountPercentage: decimal.NewFromFloat(12.5),
	},
}
