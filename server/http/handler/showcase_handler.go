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

type ShowcaseHandler struct {
	showcaseUsecase usecase.ShowcaseUsecase
}

func NewShowcaseHandler(pu usecase.ShowcaseUsecase) *ShowcaseHandler {
	return &ShowcaseHandler{
		showcaseUsecase: pu,
	}
}

func (h *ShowcaseHandler) CreateShowcase(c *gin.Context) {
	uid := c.GetInt("userId")

	req := dtohttp.CreateShowcaseRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.CreateShowcaseRequest{
		ShopId:             uid,
		Name:               req.Name,
		SelectedProductsId: req.SelectedProductsId,
	}

	response, err := h.showcaseUsecase.CreateShopPromotions(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *ShowcaseHandler) DeleteShowcase(c *gin.Context) {
	uid := c.GetInt("userId")

	spid, err := strconv.Atoi(c.Param("showcaseId"))
	if err != nil {
		c.Error(err)
		return
	}

	err = h.showcaseUsecase.DeleteShopPromotions(c.Request.Context(), spid, uid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "showcase successfully deleted"})
}

func (h *ShowcaseHandler) GetShowcases(c *gin.Context) {
	uid := c.GetInt("userId")
	pageString := c.Query("page")
	page := 1
	if pageString != "" {
		res, err := strconv.Atoi(pageString)
		if err != nil {
			c.Error(err)
			return
		}

		page = res
	}

	response, err := h.showcaseUsecase.GetShowcases(c.Request.Context(), dtousecase.GetShowcasesRequest{ShopId: uid, Page: page})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONPagination{Data: response.Showcases, Pagination: dtogeneral.PaginationData{
		TotalPage:   response.TotalPages,
		TotalItem:   response.TotalItems,
		CurrentPage: response.CurrentPage,
		Limit:       response.Limit,
	}})
}

func (h *ShowcaseHandler) GetShowcaseDetail(c *gin.Context) {
	spid, err := strconv.Atoi(c.Param("showcaseId"))
	if err != nil {
		c.Error(err)
		return
	}

	response, err := h.showcaseUsecase.GetShowcaseDetail(c.Request.Context(), spid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *ShowcaseHandler) UpdateShowcase(c *gin.Context) {
	uid := c.GetInt("userId")

	spid, err := strconv.Atoi(c.Param("showcaseId"))
	if err != nil {
		c.Error(err)
		return
	}

	req := dtohttp.UpdateShowcaseRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.UpdateShowcaseRequest{
		Id:                 spid,
		ShopId:             uid,
		Name:               req.Name,
		SelectedProductsId: req.SelectedProductsId,
	}

	response, err := h.showcaseUsecase.UpdateShowcase(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}
