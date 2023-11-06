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

	res := dtohttp.ProductOrderReceiveResponse{
		ID:     uReq.ID,
		Notes:  uRes.Notes,
		Status: uRes.Status,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully cancel order", Data: res})
}

func (h *ProductOrderHandler) ProcessedOrderBySeller(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.ProductOrderRequest{
		ID:       id,
		SellerID: c.GetInt("userId"),
	}

	uRes, err := h.productOrderUsecase.ProcessedOrder(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.ProductOrderReceiveResponse{
		ID:     id,
		Status: uRes.Status,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully processed order", Data: res})
}

func (h *ProductOrderHandler) CheckoutOrder(c *gin.Context) {
	var req dtohttp.CheckoutOrderRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.CheckoutOrderRequest{
		ProductVariant:       req.ProductVariant,
		VoucherID:            req.VoucherID,
		DestinationAddressID: req.DestinationAddressID,
		UserID:               c.GetInt("userId"),
		CourierID:            req.CourierID,
		Notes:                req.Notes,
		SellerID:             req.SellerID,
	}

	_, err = h.productOrderUsecase.CheckoutOrder(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully checkout order"})
}
