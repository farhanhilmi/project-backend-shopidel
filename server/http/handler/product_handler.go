package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

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

func (h *ProductHandler) UploadPhotos(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}
	log.Println("FILE", file)

	currentTime := time.Now().UnixNano()
	fileExtension := path.Ext(header.Filename)

	originalFilename := header.Filename[:len(header.Filename)-len(fileExtension)]

	newFilename := fmt.Sprintf("%s_%d", originalFilename, currentTime)

	imageUrl, err := util.UploadToCloudinary(file, newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: imageUrl})
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
	categoryId := c.DefaultQuery("categoryId", "")
	s := c.DefaultQuery("s", "")
	minRating := c.DefaultQuery("minRating", "")
	minPrice := c.DefaultQuery("minPrice", "")
	maxPrice := c.DefaultQuery("maxPrice", "")
	district := c.DefaultQuery("district", "")

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

	uReq := dtousecase.ProductListParam{
		CategoryId: categoryId,
		SortBy:     sortBy,
		Sort:       sort,
		MinRating:  util.StrToInt(minRating),
		MinPrice:   util.StrToInt(minPrice),
		MaxPrice:   util.StrToInt(maxPrice),
		District:   district,
		Limit:      limit,
		Page:       page,
		StartDate:  startDate,
		EndDate:    endDate,
		AccountID:  c.GetInt("userId"),
		Search:     s,
	}

	uRes, pagination, err := h.productUsecase.GetProducts(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONPagination{Data: uRes, Pagination: *pagination})
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

func (h *ProductHandler) GetProductDetailV2(c *gin.Context) {
	uReq := dtousecase.GetProductDetailRequestV2{
		AccountId:   c.GetInt("userId"),
		ShopName:    c.Param("shopName"),
		ProductName: c.Param("productName"),
	}

	uRes, err := h.productUsecase.GetProductDetailV2(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes})
}

func (h *ProductHandler) GetProductPictures(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.GetProductPicturesRequest{
		ProductId: productId,
	}

	uRes, err := h.productUsecase.GetProductPictures(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONResponse{
		Data: uRes.PicturesUrl,
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) GetProductReviews(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.GetProductReviewsRequest{
		ProductId: productId,
	}

	err = h.handleProductReviewsQueryParams(c, &uReq)
	if err != nil {
		c.Error(err)
		return
	}

	uRes, err := h.productUsecase.GetProductReviews(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONPagination{
		Data: uRes.Reviews,
		Pagination: dtogeneral.PaginationData{
			TotalPage:   uRes.TotalPage,
			TotalItem:   uRes.TotalItem,
			CurrentPage: uRes.CurrentPage,
			Limit:       uRes.Limit,
		},
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) GetProductDetailRecomendedProduct(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.GetProductDetailRecomendedProductRequest{
		ProductId: productId,
	}

	uRes, err := h.productUsecase.GetProductDetailRecomendedProducts(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONResponse{
		Data: uRes.AnotherProducts,
	}

	c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) handleProductReviewsQueryParams(c *gin.Context, uReq *dtousecase.GetProductReviewsRequest) error {
	uReq.OrderBy = "asc"
	if c.Query("orderBy") != "" {
		if c.Query("orderBy") != "newest" {
			uReq.OrderBy = "asc"
		} else if c.Query("orderBy") != "oldest" {
			uReq.OrderBy = "desc"
		} else {
			return errors.New("available orderBy query option is newest or oldest")
		}
	}

	uReq.Comment = false
	if c.Query("comment") != "" {
		if c.Query("comment") == "true" {
			uReq.Comment = true
		} else if c.Query("comment") == "false" {
			uReq.Comment = false
		} else {
			return errors.New("available comment query option is true or false")
		}
	}

	uReq.Image = false
	if c.Query("image") != "" {
		if c.Query("image") == "true" {
			uReq.Image = true
		} else if c.Query("image") == "false" {
			uReq.Image = false
		} else {
			return errors.New("available image query option is true or false")
		}
	}

	uReq.Page = 1
	if c.Query("page") != "" {
		pageString := c.Query("page")
		page, err := strconv.Atoi(pageString)
		if err != nil {
			return errors.New("available page query option is integer")
		}

		if page < 1 {
			return errors.New("page query must above 0")
		}

		uReq.Page = page
	}

	uReq.Stars = 0
	if c.Query("stars") != "" {
		starsString := c.Query("stars")
		stars, err := strconv.Atoi(starsString)
		if err != nil {
			return errors.New("available stars query option is integer")
		}

		if stars < 1 || stars > 5 {
			return errors.New("available stars range is 1 - 5")
		}

		uReq.Stars = stars
	}

	return nil
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

	c.JSON(http.StatusOK, dtogeneral.JSONPagination{Data: products, Pagination: *pagination})
}
