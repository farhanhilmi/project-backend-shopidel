package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(h *handler.AccountHandler, gin *gin.Engine) *gin.Engine {
	account := gin.Group("accounts")
	{
		account.Use(middleware.AuthenticateJWT())
		account.GET("/couriers", h.GetCouriers)

		account.GET("/carts", h.GetCart)
		account.POST("/carts", h.AddProductToCart)

		account.POST("/check-password", h.CheckISPasswordCorrect)
		account.PUT("/carts", h.UpdateCart)
		account.POST("/carts/delete", h.DeleteCartProduct)

		profile := account.Group("profile")
		{
			profile.GET("", h.GetProfile)
			profile.PUT("", h.EditProfile)
		}

		wallet := account.Group("wallets")
		{
			wallet.POST("/activate", h.ActivateMyWallet)
			wallet.PUT("/change-pin", h.ChangeWalletPIN)
			wallet.GET("", h.GetWallet)
			wallet.GET("/histories", h.GetListTransactions)
			wallet.POST("/validate-pin", h.ValidateWalletPIN)

			wallet.POST("/topup", h.TopUpBalanceWallet)
		}

		address := account.Group("address")
		{
			address.GET("", h.GetAddresses)
			address.POST("", h.RegisterAdress)
		}
	}

	address := gin.Group("/address")
	{
		address.GET("/provinces", h.GetProvinces)
		address.GET("/provinces/:provinceId/districts", h.GetDistricts)
	}

	return gin
}
