package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewProductRouter(h *handler.ProductHandler, gin *gin.Engine) *gin.Engine {
	product := gin.Group("/products")

	product.GET("", h.ListProduct)
	product.GET("/:productId", h.GetProductDetail)
	product.POST("/:productId/favorites/add-favorite", middleware.AuthenticateJWT(), h.AddToFavorite)
	product.GET("/favorites", middleware.AuthenticateJWT(), h.GetListFavorite)

	return gin
}
