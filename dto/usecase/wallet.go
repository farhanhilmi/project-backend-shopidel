package dtousecase

import "github.com/shopspring/decimal"

type UpdateWalletPINRequest struct {
	UserID       int
	WalletPIN    string
	WalletNewPIN string
}

type UpdateWalletPINResponse struct {
	WalletNewPIN string
}

type WalletResponse struct {
	Balance      decimal.Decimal
	WalletNumber string
	IsActive     bool
}

type TopUpBalanceWalletRequest struct {
	UserID int
	Amount decimal.Decimal
}

type TopUpBalanceWalletResponse struct {
	WalletNumber string
	Balance      decimal.Decimal
}
