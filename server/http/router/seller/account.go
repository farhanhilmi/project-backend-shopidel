package router_seller

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"github.com/gin-gonic/gin"
)

func NewSellerProfileRouter(h *handler.SellerHandler, gin *gin.Engine) *gin.Engine {
	seller := gin.Group("sellers/profile")
	{
		seller.GET("/:sellerId", h.GetProfile)
	}

	return gin
}
