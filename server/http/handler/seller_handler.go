package handler

import (
	"log"
	"net/http"
	"strconv"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
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

	c.JSON(http.StatusOK, dtogeneral.JSONPagination{Data: uRes.SellerProducts, Pagination: dtogeneral.PaginationData{CurrentPage: p, Limit: limit}})
}

func (h *SellerHandler) AddNewProduct(c *gin.Context) {
	var req dtohttp.AddNewProductRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("ERROR", err)
		c.Error(util.ErrInvalidInput)
		return
	}

	form, err := c.MultipartForm()
	files := form.File["images[]"]
	// for i, header := range files {
	// 	file, err := files[i].Open()
	// 	log.Println(header.Filename)
	// }

	if err != nil {
		log.Println("ERR", err)
		c.Error(err)
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
		Images:            files,
	}

	response, err := h.sellerUsecase.AddNewProduct(c.Request.Context(), productReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}
