package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ProductOrderDetails = []model.ProductOrderDetails{
	{
		ProductOrderID:                       1,
		ProductVariantSelectionCombinationID: 1,
		Quantity:                             2,
		IndividualPrice:                      decimal.NewFromInt(20000),
	},
	{
		ProductOrderID:                       1,
		ProductVariantSelectionCombinationID: 3,
		Quantity:                             1,
		IndividualPrice:                      decimal.NewFromInt(50000),
	},
}
