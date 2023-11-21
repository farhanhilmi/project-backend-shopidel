package dtorepository

import (
	"github.com/shopspring/decimal"
)

type UpdateWalletPINRequest struct {
	UserID       int
	WalletNewPIN string
}

type UpdateWalletPINResponse struct {
	UserID       int
	WalletNewPIN string
}

type WalletResponse struct {
	UserID       int
	WalletNumber string
	Balance      decimal.Decimal
}

type TopUpWalletRequest struct {
	UserID  int
	Amount  decimal.Decimal
	Type    string
	From    string
	OrderID int
}

type MyWalletTransactionHistoriesRequest struct {
	AccountID      int
	Type           string
	Amount         decimal.Decimal
	ProductOrderID int
	From           string
	To             string
}

type MyWalletTransactionHistoriesResponse struct {
	ID             int
	AccountID      int
	Type           string
	Amount         decimal.Decimal
	ProductOrderID int
}

type SaleWalletTransactionHistoriesRequest struct {
	AccountID      int
	Type           string
	Amount         decimal.Decimal
	ProductOrderID int
	From           string
	To             string
}

type SaleWalletTransactionHistoriesResponse struct {
	ID             int
	AccountID      int
	Type           string
	Amount         decimal.Decimal
	ProductOrderID int
}

type MyWalletRequest struct {
	UserID          int
	WalletNumber    string
	Balance         decimal.Decimal
	TransactionType string
	ProductOrderID  int
}

type MyWalletResponse struct {
	UserID       int
	WalletNumber string
	Balance      decimal.Decimal
}

type WalletHistoriesParams struct {
	AccountID       int
	SortBy          string
	Sort            string
	Limit           int
	Page            int
	StartDate       string
	EndDate         string
	TransactionType string
}
