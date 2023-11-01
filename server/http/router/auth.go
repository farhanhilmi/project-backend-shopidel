package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"github.com/gin-gonic/gin"
)

func NewAuthRouter(h *handler.AccountHandler, gin *gin.Engine) *gin.Engine {
	group := gin.Group("api/auth")
	group.POST("/register", h.CreateAccount)
	return gin
}
