package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ProductOrderDetails = []model.ProductOrderDetails{
	{
		ProductOrderID:  1,
		ProductID:       4,
		Quantity:        2,
		IndividualPrice: decimal.NewFromInt(30000),
	},
}
