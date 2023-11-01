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
