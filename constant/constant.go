package constant

import "github.com/shopspring/decimal"

var (
	TopupAmountMin = decimal.NewFromInt(50000)
	TopupAmountMax = decimal.NewFromInt(10000000)
)
