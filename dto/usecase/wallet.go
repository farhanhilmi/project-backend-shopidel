package dtousecase

type UpdateWalletPINRequest struct {
	UserID       int
	WalletPIN    string
	WalletNewPIN string
}

type UpdateWalletPINResponse struct {
	WalletNewPIN string
}
