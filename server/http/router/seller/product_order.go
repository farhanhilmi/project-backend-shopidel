package router_seller

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewProductOrderRouter(h *handler.ProductOrderHandler, gin *gin.Engine) *gin.Engine {
	group := gin.Group("sellers/orders")
	group.Use(middleware.AuthenticateJWT())
	group.PUT("/:orderId/cancel", h.CanceledOrderBySeller)
	group.PUT("/:orderId/processed", h.ProcessedOrderBySeller)

	return gin
}
