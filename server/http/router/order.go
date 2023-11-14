package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewProductOrderRouter(h *handler.ProductOrderHandler, gin *gin.Engine) *gin.Engine {
	order := gin.Group("orders")
	order.Use(middleware.AuthenticateJWT())
	order.POST("/checkout", h.CheckoutOrder)
	order.POST("/cost/check", h.CheckDeveliryFee)
	order.GET("/couriers/:sellerId", h.GetCouriers)
	order.GET("/histories", h.GetOrderHistories)
	order.GET("/:orderId", h.GetOrderDetail)
	order.POST("/:orderId/add-review", h.AddProductReview)

	return gin
}
