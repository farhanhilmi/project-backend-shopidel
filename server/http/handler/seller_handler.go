package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
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
	uReq := dtousecase.GetSellerProfileRequest{}

	shopId, err := strconv.Atoi(shopName)
	if err != nil {
		uReq.ShopName = shopName
	} else {
		uReq.ShopId = shopId
	}

	uRes, err := h.sellerUsecase.GetProfile(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.GetSellerProfileResponse{
		SellerName:       uRes.SellerName,
		SellerPictureUrl: uRes.SellerPictureUrl,
		SellerDistrict:   uRes.SellerDistrict,
		SellerOperatingHour: dtohttp.SellerOperatingHour{
			Start: uRes.SellerOperatingHour.Start.Format("15:04"),
			End:   uRes.SellerOperatingHour.End.Format("15:04"),
		},
		ShopNameSlug:      uRes.ShopNameSlug,
		SellerStars:       uRes.SellerStars,
		SellerDescription: uRes.SellerDescription,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
}

func (h *SellerHandler) UpdateProfile(c *gin.Context) {
	shopId := c.GetInt("userId")
	body := dtohttp.UpdateSellerProfileBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}

	openingHours, err := time.Parse("15:04", body.OpeningHours)
	if err != nil {
		c.Error(errors.New("opening hours format is 00:00"))
		return
	}

	closingHours, err := time.Parse("15:04", body.ClosingHours)
	if err != nil {
		c.Error(errors.New("closing hours format is 00:00"))
		return
	}

	uReq := dtousecase.UpdateShopProfileRequest{
		ShopId:          shopId,
		ShopName:        body.ShopName,
		ShopDescription: body.ShopDescription,
		OpeningHours:    openingHours,
		ClosingHours:    closingHours,
	}

	uRes, err := h.sellerUsecase.UpdateShopProfile(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.UpdateSellerProfileResponse{
		ShopName:        uRes.ShopName,
		ShopDescription: uRes.ShopDescription,
		OpeningHours:    uRes.OpeningHours.Format("15:04"),
		ClosingHours:    uRes.ClosingHours.Format("15:04"),
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
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

func (h *SellerHandler) GetShowcases(c *gin.Context) {
	shopName := c.Param("shopName")

	uRes, err := h.sellerUsecase.GetShowcases(c.Request.Context(), dtousecase.GetSellerShowcasesRequest{ShopName: shopName})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes.Showcases})
}

func (h *SellerHandler) GetShowcaseProducts(c *gin.Context) {
	shopName := c.Param("shopName")
	showcaseId := c.Param("showcaseId")
	page := c.Query("page")
	var err error
	p := 0

	if page != "" {
		if p, err = strconv.Atoi(page); err != nil {
			c.Error(err)
			return
		}
	}

	if p == 0 {
		p = 1
	}
	limit := 20

	uRes, err := h.sellerUsecase.GetShowcaseProducts(c.Request.Context(), dtousecase.GetSellerShowcaseProductRequest{ShopName: shopName, ShowcaseId: showcaseId, Page: p, Limit: limit})
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONPagination{
		Data: uRes.SellerProducts,
		Pagination: dtogeneral.PaginationData{
			CurrentPage: uRes.CurrentPage,
			Limit:       uRes.Limit,
			TotalPage:   uRes.TotalPage,
			TotalItem:   uRes.TotalItem,
		},
	}
	c.JSON(http.StatusOK, res)
}

func (h *SellerHandler) UploadPhoto(c *gin.Context) {
	var req dtohttp.UploadNewPhoto

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.Error(util.ErrNoImage)
		return
	}

	uReq := dtousecase.UploadNewPhoto{
		ImageID:     req.ImageID,
		Image:       file,
		ImageHeader: header,
	}

	response, err := h.sellerUsecase.UploadPhoto(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *SellerHandler) AddNewProduct(c *gin.Context) {
	var req dtohttp.AddNewProductRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	form, err := c.MultipartForm()
	files := form.File["images[]"]

	if err != nil {
		c.Error(util.ErrNoImage)
		return
	}

	productReq := dtousecase.AddNewProductRequest{
		SellerID:          c.GetInt("userId"),
		ProductName:       req.ProductName,
		Description:       req.Description,
		CategoryID:        req.CategoryID,
		HazardousMaterial: req.HazardousMaterial,
		IsNew:             req.IsNew,
		InternalSKU:       req.InternalSKU,
		Weight:            req.Weight,
		Size:              req.Size,
		IsActive:          req.IsActive,
		Variants:          req.Variants,
		VideoURL:          req.VideoURL,
		Images:            files,
	}

	newProduct, err := h.sellerUsecase.AddNewProduct(c.Request.Context(), productReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := fmt.Sprintf("Successfully create new product %v", newProduct.ProductName)
	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: res})
}

func (h *SellerHandler) UpdateProduct(c *gin.Context) {
	var req dtohttp.AddNewProductRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	productReq := dtousecase.AddNewProductRequest{
		ProductID:         productId,
		SellerID:          c.GetInt("userId"),
		ProductName:       req.ProductName,
		Description:       req.Description,
		CategoryID:        req.CategoryID,
		HazardousMaterial: req.HazardousMaterial,
		IsNew:             req.IsNew,
		InternalSKU:       req.InternalSKU,
		Weight:            req.Weight,
		Size:              req.Size,
		IsActive:          req.IsActive,
		Variants:          req.Variants,
		VideoURL:          req.VideoURL,
	}

	form, err := c.MultipartForm()
	if err != nil {
		productReq.Images = nil
	} else {
		files := form.File["images[]"]
		productReq.Images = files
	}

	product, err := h.sellerUsecase.UpdateProduct(c.Request.Context(), productReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := fmt.Sprintf("Successfully update product %v", product.ProductName)
	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: res})
}

func (h *SellerHandler) DeleteProduct(c *gin.Context) {

	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	productReq := dtousecase.RemoveProduct{
		ID:       productId,
		SellerID: c.GetInt("userId"),
	}

	product, err := h.sellerUsecase.DeleteProduct(c.Request.Context(), productReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := fmt.Sprintf("Successfully deleted product %v", product.Name)
	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: res})
}

func (h *SellerHandler) ListProduct(c *gin.Context) {
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
		SellerID:   c.GetInt("userId"),
		Search:     s,
	}

	uRes, pagination, err := h.sellerUsecase.GetProducts(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONPagination{Data: uRes, Pagination: *pagination})
}

func (h *SellerHandler) GetProductByID(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.GetProductDetailRequest{
		ProductId: productId,
		AccountId: c.GetInt("userId"),
	}

	uRes, err := h.sellerUsecase.GetProductByID(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes})
}

func (h *SellerHandler) WithdrawOrderSales(c *gin.Context) {
	id := c.Param("orderId")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	productReq := dtousecase.WithdrawBalance{
		SellerID: c.GetInt("userId"),
		OrderID:  orderId,
	}

	product, err := h.sellerUsecase.WithdrawSalesBalance(c.Request.Context(), productReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: product})
}
