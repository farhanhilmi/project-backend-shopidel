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

type PromotionOrderHandler struct {
	promotionUsecase usecase.PromotionUsecase
}

func NewPromotionHandler(pu usecase.PromotionUsecase) *PromotionOrderHandler {
	return &PromotionOrderHandler{
		promotionUsecase: pu,
	}
}

func (h *PromotionOrderHandler) CreateShopPromotion(c *gin.Context) {
	uid := c.GetInt("userId")

	req := dtohttp.CreateShopPromotionRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.CreateShopPromotionRequest{
		ShopId:             uid,
		Name:               req.Name,
		Quota:              req.Quota,
		StartDate:          req.StartDate,
		EndDate:            req.EndDate,
		MinPurchaseAmount:  req.MinPurchaseAmount,
		MaxPurchaseAmount:  req.MaxPurchaseAmount,
		DiscountPercentage: req.DiscountPercentage,
		SelectedProductsId: req.SelectedProductsId,
	}

	response, err := h.promotionUsecase.CreateShopPromotions(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *PromotionOrderHandler) GetShopPromotions(c *gin.Context) {
	uid := c.GetInt("userId")

	response, err := h.promotionUsecase.GetShopPromotions(c.Request.Context(), dtousecase.GetShopPromotionsRequest{ShopId: uid})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response.ShopPromotions})
}

func (h *PromotionOrderHandler) GetShopPromotionDetail(c *gin.Context) {
	spid, err := strconv.Atoi(c.Param("shopPromotionId"))
	if err != nil {
		c.Error(err)
		return
	}

	response, err := h.promotionUsecase.GetShopPromotionDetail(c.Request.Context(), spid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *PromotionOrderHandler) UpdateShopPromotion(c *gin.Context) {
	uid := c.GetInt("userId")

	spid, err := strconv.Atoi(c.Param("shopPromotionId"))
	if err != nil {
		c.Error(err)
		return
	}

	req := dtohttp.UpdateShopPromotionRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.UpdateShopPromotionRequest{
		Id:                 spid,
		ShopId:             uid,
		Name:               req.Name,
		Quota:              req.Quota,
		StartDate:          req.StartDate,
		EndDate:            req.EndDate,
		MinPurchaseAmount:  req.MinPurchaseAmount,
		MaxPurchaseAmount:  req.MaxPurchaseAmount,
		DiscountPercentage: req.DiscountPercentage,
		SelectedProductsId: req.SelectedProductsId,
	}

	response, err := h.promotionUsecase.UpdateShopPromotion(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}
