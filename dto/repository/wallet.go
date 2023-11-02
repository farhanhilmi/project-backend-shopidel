package dtorepository

import "github.com/shopspring/decimal"

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
}
