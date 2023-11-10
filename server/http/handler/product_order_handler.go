package handler

import (
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
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
		ID:                 id,
		SellerID:           c.GetInt("userId"),
		Notes:              req.Notes,
		SellerWalletNumber: c.GetString("walletNumber"),
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
		Weight:               req.Weight,
	}

	_, err = h.productOrderUsecase.CheckoutOrder(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully checkout order"})
}

func (h *ProductOrderHandler) CheckDeveliryFee(c *gin.Context) {
	var req dtohttp.CheckDeliveryFeeRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.CheckDeliveryFeeRequest{
		SellerID:    req.SellerID,
		ID:          req.CourierID,
		Destination: req.Destination,
		Weight:      req.Weight,
	}

	response, err := h.productOrderUsecase.CheckDeliveryFee(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *ProductOrderHandler) GetCouriers(c *gin.Context) {
	id := c.Param("sellerId")
	sellerId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	uReq := dtousecase.SellerCourier{
		SellerID: sellerId,
	}

	response, err := h.productOrderUsecase.GetCouriers(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *ProductOrderHandler) GetOrderHistories(c *gin.Context) {
	status := c.DefaultQuery("status", constant.StatusOrderAll)

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.Error(err)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.Error(err)
		return
	}
	sortBy := c.DefaultQuery("sortBy", "date")
	sort := c.DefaultQuery("sort", "desc")
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")

	if valid := util.IsDateValid(startDate); !valid && startDate != "" {
		c.Error(util.ErrInvalidDateFormat)
		return
	}
	if valid := util.IsDateValid(endDate); !valid && endDate != "" {
		c.Error(util.ErrInvalidDateFormat)
		return
	}

	if !slices.Contains([]string{"date", "price"}, sortBy) {
		c.Error(util.ErrProductFavoriteSortBy)
		return
	}

	switch sortBy {
	case "date":
		sortBy = "created_at"
	}

	uReq := dtousecase.ProductOrderHistoryRequest{
		AccountID: c.GetInt("userId"),
		Status:    status,
		SortBy:    sortBy,
		Sort:      sort,
		Limit:     limit,
		Page:      page,
		StartDate: startDate,
		EndDate:   endDate,
	}

	response, pagination, err := h.productOrderUsecase.GetOrderHistories(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONWithPagination{Data: response, Pagination: pagination})
}
