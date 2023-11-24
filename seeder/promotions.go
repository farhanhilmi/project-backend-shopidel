package seeder

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ShopPromotions = []*model.ShopPromotion{
	{
		ID:                1,
		ShopId:            4,
		Name:              "Kalung dan sandal discount",
		Quota:             100,
		StartDate:         time.Date(2020, 01, 01, 01, 01, 01, 01, time.Local),
		EndDate:           time.Date(2024, 01, 01, 01, 01, 01, 01, time.Local),
		MinPurchaseAmount: decimal.NewFromFloat(10000),
		MaxPurchaseAmount: decimal.NewFromFloat(200000),
		DiscountAmount:    decimal.NewFromFloat(5000),
	},
	{
		ID:                2,
		ShopId:            4,
		Name:              "Sandal dan ikat pinggang discount",
		Quota:             100,
		StartDate:         time.Date(2020, 01, 01, 01, 01, 01, 01, time.Local),
		EndDate:           time.Date(2024, 01, 01, 01, 01, 01, 01, time.Local),
		MinPurchaseAmount: decimal.NewFromFloat(10000),
		MaxPurchaseAmount: decimal.NewFromFloat(200000),
		DiscountAmount:    decimal.NewFromFloat(5000),
	},
}

var MarketplacePromotions = []model.MarketplacePromotion{
	{
		ID:                1,
		Name:              "Discount meriah November",
		Quota:             100,
		StartDate:         time.Date(2020, 01, 01, 01, 01, 01, 01, time.Local),
		EndDate:           time.Date(2024, 01, 01, 01, 01, 01, 01, time.Local),
		MinPurchaseAmount: decimal.NewFromFloat(10000),
		MaxPurchaseAmount: decimal.NewFromFloat(200000),
		DiscountAmount:    decimal.NewFromFloat(5000),
	},
}
