package handler

import (
	"errors"
	"net/http"
	"strconv"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
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

func (h *ProductHandler) GetTopCategories(c *gin.Context) {

	topCategories := []dtorepository.TopCategoriesResponse {
		{
			CategoryID: 504,
			Name: "Fashion Anak & Bayi",
			PictureURL: "https://down-id.img.susercontent.com/file/9251edd6d6dd98855ff5a99497835d9c_tn",
		},
		{
			CategoryID: 1,
			Name: "Audio, Kamera & Elektronik Lainnya",
			PictureURL: "https://down-id.img.susercontent.com/file/dcd61dcb7c1448a132f49f938b0cb553_tn",
		},
		{
			CategoryID: 694,
			Name: "Fashion Pria",
			PictureURL: "https://down-id.img.susercontent.com/file/04dba508f1ad19629518defb94999ef9_tn",
		},
		{
			CategoryID: 1659,
			Name: "Makanan & Minuman",
			PictureURL: "https://down-id.img.susercontent.com/file/7873b8c3824367239efb02d18eeab4f5_tn",
		},
		{
			CategoryID: 222,
			Name: "Komputer & Internet",
			PictureURL: "https://down-id.img.susercontent.com/file/id-50009109-0bd6a9ebd0f2ae9b7e8b9ce7d89897d6_tn",
		},
		{
			CategoryID: 1228,
			Name: "Kecantikan",
			PictureURL: "https://down-id.img.susercontent.com/file/2715b985ae706a4c39a486f83da93c4b_tn",
		},
		{
			CategoryID: 2153,
			Name: "Sepatu Lari Pria",
			PictureURL: "https://down-id.img.susercontent.com/file/3c8ff51aab1692a80c5883972a679168_tn",
		},
		{
			CategoryID: 1561,
			Name: "Mainan & Hobi",
			PictureURL: "https://down-id.img.susercontent.com/file/42394b78fac1169d67c6291973a3b132_tn",
		},
		{
			CategoryID: 1335,
			Name: "Kesehatan",
			PictureURL: "https://down-id.img.susercontent.com/file/eb7d583e4b72085e71cd21a70ce47d7a_tn",
		},
		{
			CategoryID: 733,
			Name: "Jam Tangan Pria",
			PictureURL: "https://down-id.img.susercontent.com/file/2bdf8cf99543342d4ebd8e1bdb576f80_tn",
		},
		{
			CategoryID: 696,
			Name: "Aksesoris Kacamata Pria",
			PictureURL: "https://down-id.img.susercontent.com/file/1f18bdfe73df39c66e7326b0a3e08e87_tn",
		},
		{
			CategoryID: 782,
			Name: "Tas Pria",
			PictureURL: "https://down-id.img.susercontent.com/file/47ed832eed0feb62fd28f08c9229440e_tn",
		},
		{
			CategoryID: 1085,
			Name: "Handphone & Tablet",
			PictureURL: "https://down-id.img.susercontent.com/file/5230277eefafad8611aaf703d3e99568_tn",
		},
		{
			CategoryID: 328,
			Name: "Alat Masak Khusus",
			PictureURL: "https://down-id.img.susercontent.com/file/c1494110e0383780cdea73ed890e0299_tn",
		},
		{
			CategoryID: 3087,
			Name: "Gaun & Pakaian Wanita",
			PictureURL: "https://down-id.img.susercontent.com/file/6d63cca7351ba54a2e21c6be1721fa3a_tn",
		},
		{
			CategoryID: 629,
			Name: "Fashion Muslim",
			PictureURL: "https://down-id.img.susercontent.com/file/b98756cdb31eabe3d7664599e24ccc29_tn",
		},
		{
			CategoryID: 1158,
			Name: "Ibu & Bayi",
			PictureURL: "https://down-id.img.susercontent.com/file/4d1673a14c26c8361a76258d78446324_tn",
		},
		{
			CategoryID: 917,
			Name: "Sepatu Wanita",
			PictureURL: "https://down-id.img.susercontent.com/file/id-50009109-a947822064b7a8077b15596c85bd9303_tn",
		},
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully listed top categories", Data: topCategories})
}

func (h *ProductHandler) GetBanners(c *gin.Context) {
	banners := []string {
		"https://images.tokopedia.net/img/cache/1208/NsjrJu/2023/11/14/55e4422c-a3a7-4248-9c63-aade2a214ba6.jpg.webp?ect=4g",
		"https://images.tokopedia.net/img/cache/1208/NsjrJu/2023/11/14/329a7a18-cd48-45b7-b70b-49a1413bedbc.jpg.webp?ect=4g",
		"https://images.tokopedia.net/img/cache/1208/NsjrJu/2023/11/14/26c62498-33d4-4db6-bad2-bac7710d2746.jpg.webp?ect=4g",
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully listed banners", Data: banners})
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
