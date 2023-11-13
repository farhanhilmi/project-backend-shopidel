package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ProductOrderDetails = []model.ProductOrderDetails{
	{
		ProductOrderID:                       1,
		ProductVariantSelectionCombinationID: 7,
		Quantity:                             2,
		IndividualPrice:                      decimal.NewFromInt(30000),
	},
	{
		ProductOrderID:                       1,
		ProductVariantSelectionCombinationID: 8,
		Quantity:                             1,
		IndividualPrice:                      decimal.NewFromInt(17000),
	},
	{
		ProductOrderID:                       2,
		ProductVariantSelectionCombinationID: 9,
		Quantity:                             2,
		IndividualPrice:                      decimal.NewFromInt(20000),
	},
	{
		ProductOrderID:                       3,
		ProductVariantSelectionCombinationID: 11,
		Quantity:                             1,
		IndividualPrice:                      decimal.NewFromInt(18000),
	},
	{
		ProductOrderID:                       4,
		ProductVariantSelectionCombinationID: 10,
		Quantity:                             1,
		IndividualPrice:                      decimal.NewFromInt(19000),
	},
	{
		ProductOrderID:                       4,
		ProductVariantSelectionCombinationID: 59,
		Quantity:                             2,
		IndividualPrice:                      decimal.NewFromInt(37000),
	},
}
