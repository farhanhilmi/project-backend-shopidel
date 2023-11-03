package handler

import (
	"net/http"
	"strconv"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
)

type ProductOrderHandler struct {
	productOrderUsecase usecase.ProductOrderUsecase
}

func NewProductOrderHandler(productOrderUsecase usecase.ProductOrderUsecase) *ProductOrderHandler {
	return &ProductOrderHandler{
		productOrderUsecase: productOrderUsecase,
	}
}

func (h *ProductOrderHandler) CanceledOrderBySeller(c *gin.Context) {
	var req dtohttp.CanceledOrderRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	id, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.ProductOrderRequest{
		ID:       id,
		SellerID: c.GetInt("userId"),
		Notes:    req.Notes,
	}

	uRes, err := h.productOrderUsecase.CancelOrderBySeller(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.ProductOrderResponse{
		ID:     uReq.ID,
		Notes:  uRes.Notes,
		Status: uReq.Status,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Succesfully cancel order", Data: res})
}
