package handler

import (
	"net/http"
	"strconv"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) ListProduct(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "30"))
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
	s := c.DefaultQuery("s", "")

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
	uReq := dtousecase.ProductListParam {	
		SortBy:    sortBy,
		Sort:      sort,
		Limit:     limit,
		Page:      page,
		StartDate: startDate,
		EndDate:   endDate,
		AccountID: c.GetInt("userId"),
		Search:    s,
	}

	uRes, pagination, err := h.productUsecase.GetProducts(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONWithPagination{Data: uRes, Pagination: *pagination})
}

func (h *ProductHandler) GetProductDetail(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.GetProductDetailRequest{
		ProductId: productId,
	}

	uRes, err := h.productUsecase.GetProductDetail(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes})
}

func (h *ProductHandler) AddToFavorite(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.FavoriteProduct{
		ProductID: productId,
		AccountID: c.GetInt("userId"),
	}

	_, err = h.productUsecase.AddToFavorite(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully add product to favorite"})
}

func (h *ProductHandler) GetListFavorite(c *gin.Context) {
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
	s := c.DefaultQuery("s", "")

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

	uReq := dtousecase.ProductFavoritesParams{
		SortBy:    sortBy,
		Sort:      sort,
		Limit:     limit,
		Page:      page,
		StartDate: startDate,
		EndDate:   endDate,
		AccountID: c.GetInt("userId"),
		Search:    s,
	}

	products, pagination, err := h.productUsecase.GetProductFavorites(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONWithPagination{Data: products, Pagination: *pagination})
}
