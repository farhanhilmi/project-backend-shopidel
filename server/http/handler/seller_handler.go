package handler

import (
	"net/http"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	sellerUsecase usecase.SellerUsecase
}

type SellerHandlerConfig struct {
	SellerUsecase usecase.SellerUsecase
}

func NewSellerHandler(config SellerHandlerConfig) *SellerHandler {
	ah := &SellerHandler{}

	if config.SellerUsecase != nil {
		ah.sellerUsecase = config.SellerUsecase
	}

	return ah
}

func (h *SellerHandler) GetProfile(c *gin.Context) {
	shopName := c.Param("shopName")

	uRes, err := h.sellerUsecase.GetProfile(c.Request.Context(), dtousecase.GetSellerProfileRequest{ShopName: shopName})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes})
}

func (h *SellerHandler) GetBestSelling(c *gin.Context) {
	shopName := c.Param("shopName")

	uRes, err := h.sellerUsecase.GetBestSelling(c.Request.Context(), dtousecase.GetSellerProductsRequest{ShopName: shopName})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes.SellerProducts})
}

func (h *SellerHandler) GetCategories(c *gin.Context) {
	shopName := c.Param("shopName")

	uRes, err := h.sellerUsecase.GetCategories(c.Request.Context(), dtousecase.GetSellerCategoriesRequest{ShopName: shopName})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes.Categories})
}

func (h *SellerHandler) GetCategoryProducts(c *gin.Context) {
	shopName := c.Param("shopName")
	categoryId := c.Param("categoryId")

	uRes, err := h.sellerUsecase.GetCategoryProducts(c.Request.Context(), dtousecase.GetSellerCategoryProductRequest{ShopName: shopName, CategoryId: categoryId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes.SellerProducts})
}
