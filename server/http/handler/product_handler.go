package handler

import (
	"net/http"
	"strconv"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.ProductRequest{
		ProductID: productId,
	}

	uRes, err := h.productUsecase.GetProduct(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.ProductResponse{
		ID:          uRes.ID,
		Name:        uReq.Name,
		Description: uRes.Description,
		CategoryID:  uReq.CategoryID,
		Category: dtohttp.CategoryResponse{
			ID:   uRes.Category.ID,
			Name: uRes.Category.Name,
		},
		HazardousMaterial: uRes.HazardousMaterial,
		Weight:            uRes.Weight,
		Size:              uRes.Size,
		IsNew:             uRes.IsNew,
		InternalSKU:       uRes.InternalSKU,
		ViewCount:         uRes.ViewCount,
		IsActive:          uRes.IsActive,
		CreatedAt:         uRes.CreatedAt,
		UpdatedAt:         uRes.UpdatedAt,
		DeletedAt:         uRes.DeletedAt,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
}
