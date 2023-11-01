package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(h *handler.AccountHandler, gin *gin.Engine) *gin.Engine {
	group := gin.Group("api/accounts")
	group.Use(middleware.AuthenticateJWT())

	group.POST("/wallets/activate", middleware.CheckContentType(), h.ActivateMyWallet)
	group.PUT("/wallets/change-pin", middleware.CheckContentType(), h.ChangeWalletPIN)
	group.GET("/wallets", middleware.CheckContentType(), h.GetWallet)

	group.POST("/check-password", middleware.CheckContentType(), h.CheckISPasswordCorrect)

	group.GET("", h.GetProfile)
	group.PUT("", h.EditProfile)
	return gin
}
