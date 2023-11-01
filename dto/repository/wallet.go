package dtorepository

type UpdateWalletPINRequest struct {
	UserID       int
	WalletNewPIN string
}

type UpdateWalletPINResponse struct {
	UserID       int
	WalletNewPIN string
}
