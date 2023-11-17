package constant

import "github.com/shopspring/decimal"

var (
	TopupAmountMin         = decimal.NewFromInt(50000)
	TopupAmountMax         = decimal.NewFromInt(10000000)
	StatusOrderOnProcess   = "On Process"
	StatusCanceled         = "Canceled"
	StatusProcessedOrder   = "Processed"
	StatusOrderDelivered   = "Delivered"
	StatusOrderCompleted   = "Completed"
	StatusOrderAll         = "All"
	SellerRole             = "seller"
	SaleMoneyIncomeType    = "Income"
	SaleRefundCancelType   = "Refund"
	DefaultReservedKeyword = "default_reserved_keyword"
)
