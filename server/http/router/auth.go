package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewAuthRouter(h *handler.AccountHandler, gin *gin.Engine) *gin.Engine {
	auth := gin.Group("auth")
	auth.POST("/register", h.CreateAccount)
	auth.POST("/login", h.Login)
	auth.POST("/request-forget-password", h.RequestForgetPassword)
	auth.POST("/change-forget-password", h.RequestForgetChangePassword)
	seller := auth.Group("seller")
	{
		seller.Use(middleware.AuthenticateJWT())
		seller.POST("/register", h.RegisterSeller)
	}

	return gin
}
