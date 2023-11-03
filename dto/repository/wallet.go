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
	UserID int
	Amount decimal.Decimal
	Type   string
}

type MyWalletTransactionHistoriesRequest struct {
	AccountID      int
	Type           string
	Amount         decimal.Decimal
	ProductOrderID int
}

type MyWalletTransactionHistoriesResponse struct {
	ID             int
	AccountID      int
	Type           string
	Amount         decimal.Decimal
	ProductOrderID int
}

type MyWalletRequest struct {
	UserID       int
	WalletNumber string
	Balance      decimal.Decimal
}

type MyWalletResponse struct {
	UserID       int
	WalletNumber string
	Balance      decimal.Decimal
}
