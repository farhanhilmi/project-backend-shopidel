package router_seller

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewSellerProfileRouter(h *handler.SellerHandler, gin *gin.Engine) *gin.Engine {
	seller := gin.Group("sellers")
	{
		seller.GET("/:shopName/profile", h.GetProfile)
		seller.GET("/:shopName/best-selling", h.GetBestSelling)
		seller.GET("/:shopName/showcases", h.GetShowcases)
		seller.GET("/:shopName/showcases/:showcaseId/products", h.GetShowcaseProducts)
		seller.POST("/orders/:orderId/withdraw", middleware.AuthenticateJWT(), middleware.IsRoleSeller(), h.WithdrawOrderSales)

	}

	return gin
}
