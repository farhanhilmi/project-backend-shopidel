package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(h *handler.AccountHandler, gin *gin.Engine) *gin.Engine {
	group := gin.Group("accounts")
	group.Use(middleware.AuthenticateJWT())

	group.POST("/wallets/activate", h.ActivateMyWallet)
	group.PUT("/wallets/change-pin", h.ChangeWalletPIN)
	group.GET("/wallets", h.GetWallet)
	group.POST("/wallets/topup", h.TopUpBalanceWallet)
	group.GET("/carts", h.GetCart)

	group.POST("/check-password", h.CheckISPasswordCorrect)

	group.GET("", h.GetProfile)
	group.PUT("", h.EditProfile)
	group.GET("/address", h.GetAddresses)
	return gin
}
