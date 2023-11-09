package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var SaleWalletTransactionHistories = []model.SaleWalletTransactionHistories{
	{
		ProductOrderID: 1,
		AccountID:      1,
		Type:           constant.SaleMoneyIncomeType,
		Amount:         decimal.NewFromInt(90000),
		From:           "4200913923913",
	},
}
