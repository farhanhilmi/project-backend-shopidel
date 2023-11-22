package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewShopPromotionRouter(h *handler.PromotionOrderHandler, gin *gin.Engine) *gin.Engine {
	promotion := gin.Group("/shop-promotions")
	promotion.Use(middleware.AuthenticateJWT())

	promotion.GET("", h.GetShopPromotions)
	promotion.POST("", h.CreateShopPromotion)
	promotion.GET("/:shopPromotionId", h.GetShopPromotionDetail)
	promotion.PUT("/:shopPromotionId", h.UpdateShopPromotion)

	return gin
}
