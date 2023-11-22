package router_seller

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewSellerProductRouter(h *handler.SellerHandler, gin *gin.Engine) *gin.Engine {
	group := gin.Group("sellers/products")
	group.Use(middleware.AuthenticateJWT())
	group.Use(middleware.IsRoleSeller())
	group.POST("", h.AddNewProduct)
	group.POST("/upload", h.UploadPhoto)
	group.DELETE("/:productId", h.DeleteProduct)
	group.PUT("/:productId", h.UpdateProduct)
	group.GET("/:productId", h.GetProductByID)
	group.GET("", h.ListProduct)

	return gin
}
