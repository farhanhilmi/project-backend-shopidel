package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewProductRouter(h *handler.ProductHandler, gin *gin.Engine) *gin.Engine {
	product := gin.Group("/products")

	product.GET("", h.ListProduct)
	product.GET("/top-categories", h.GetTopCategories)
	product.GET("/banners", h.GetBanners)
	product.GET("/:productId/reviews", h.GetProductReviews)
	product.GET("/:productId/pictures", h.GetProductPictures)
	product.GET("/:productId/recommended-products", h.GetProductDetailRecomendedProduct)
	product.GET("/:productId/total-favorites", h.GetProductTotalFavorites)
	product.GET("/:productId", middleware.IfExistAuthenticateJWT(), h.GetProductDetail)
	product.POST("/:productId/favorites/add-favorite", middleware.AuthenticateJWT(), h.AddToFavorite)
	product.GET("/detail/:shopName/:productName", middleware.IfExistAuthenticateJWT(), h.GetProductDetailV2)
	product.GET("/favorites", middleware.AuthenticateJWT(), h.GetListFavorite)

	return gin
}
