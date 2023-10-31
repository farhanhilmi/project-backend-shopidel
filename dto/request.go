package dto

type ActivateWalletRequest struct {
	PIN string `json:"wallet_pin" binding:"max=6,min=6,required"`
}
