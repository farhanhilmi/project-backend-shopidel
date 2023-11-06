package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"github.com/gin-gonic/gin"
)

func NewProductRouter(h *handler.ProductHandler, gin *gin.Engine) *gin.Engine {
	group := gin.Group("/products")

	group.GET("/:productId", h.GetProductDetail)

	return gin
}
