package dtohttp

type ActivateWalletRequest struct {
	PIN string `json:"wallet_pin" binding:"max=6,min=6,required"`
}

type ChangeWalletPINRequest struct {
	UserID       int    `json:"user_id"`
	WalletPIN    string `json:"wallet_pin" binding:"max=6,min=6,required"`
	WalletNewPIN string `json:"wallet_new_pin" binding:"max=6,min=6,required"`
}

type ChangeWalletPINResponse struct {
	WalletNewPIN string `json:"wallet_new_pin"`
}
