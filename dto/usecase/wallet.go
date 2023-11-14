package dtousecase

import "github.com/shopspring/decimal"

type UpdateWalletPINRequest struct {
	UserID       int
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

type ValidateWAlletPINRequest struct {
	WalletPIN string
	UserID    int
	CountFail int
}

type ValidateWAlletPINResponse struct {
	CountFail int
	IsCorrect bool
}
