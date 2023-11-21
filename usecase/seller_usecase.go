package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/shopspring/decimal"
)

type SellerUsecase interface {
	GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error)
	GetBestSelling(ctx context.Context, req dtousecase.GetSellerProductsRequest) (dtousecase.GetSellerProductsResponse, error)
	GetShowcases(ctx context.Context, req dtousecase.GetSellerShowcasesRequest) (dtousecase.GetSellerShowcasesResponse, error)
	GetShowcaseProducts(ctx context.Context, req dtousecase.GetSellerShowcaseProductRequest) (dtousecase.GetSellerShowcaseProductResponse, error)
	AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error)
	UploadPhoto(ctx context.Context, req dtousecase.UploadNewPhoto) (dtousecase.UploadNewPhoto, error)
	DeleteProduct(ctx context.Context, req dtousecase.RemoveProduct) (dtousecase.RemoveProduct, error)
	GetProducts(ctx context.Context, req dtousecase.ProductListParam) (*[]dtorepository.ProductListSellerResponse, *dtogeneral.PaginationData, error)
	GetProductByID(ctx context.Context, req dtousecase.GetProductDetailRequest) (*dtousecase.GetProductSellerResponse, error)
	WithdrawSalesBalance(ctx context.Context, req dtousecase.WithdrawBalance) (dtousecase.WithdrawBalance, error)
}

type sellerUsecase struct {
	accountRepository repository.AccountRepository
	productRepository repository.ProductRepository
	orderRepository   repository.ProductOrdersRepository
}

type SellerUsecaseConfig struct {
	AccountRepository repository.AccountRepository
	ProductRepository repository.ProductRepository
	OrderRepository   repository.ProductOrdersRepository
}

func NewSellerUsecase(config SellerUsecaseConfig) SellerUsecase {
	au := &sellerUsecase{}

	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
	}
	if config.ProductRepository != nil {
		au.productRepository = config.ProductRepository
	}
	if config.OrderRepository != nil {
		au.orderRepository = config.OrderRepository
	}

	return au
}

func (u *sellerUsecase) GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error) {
	res := dtousecase.GetSellerProfileResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	res.SellerDistrict = seller.District
	res.SellerName = seller.Name
	res.SellerOperatingHour = dtousecase.SellerOperatingHour{
		Start: seller.StartOperatingHours,
		End:   seller.EndOperatingHours,
	}
	res.SellerPictureUrl = seller.ProfilePicture
	res.ShopNameSlug = seller.ShopNameSlug
	res.SellerStars = "4.8"

	return res, nil
}

func (u *sellerUsecase) GetBestSelling(ctx context.Context, req dtousecase.GetSellerProductsRequest) (dtousecase.GetSellerProductsResponse, error) {
	res := dtousecase.GetSellerProductsResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	rRes, err := u.accountRepository.FindSellerBestSelling(ctx, dtorepository.FindSellerBestSellingRequest{SellerId: seller.Id})
	if err != nil {
		return res, err
	}

	for _, data := range rRes.Products {
		p := dtousecase.SellerProduct{
			Name:            data.Name,
			Price:           data.Price,
			PictureUrl:      data.PictureUrl,
			Stars:           data.Stars,
			TotalSold:       data.TotalSold,
			CreatedAt:       data.CreatedAt,
			Category:        data.Category,
			ShopName:        data.ShopName,
			ProductNameSlug: data.ProductNameSlug,
			ShopNameSlug:    data.ShopNameSlug,
		}

		res.SellerProducts = append(res.SellerProducts, p)
	}

	return res, nil
}

func (u *sellerUsecase) GetShowcases(ctx context.Context, req dtousecase.GetSellerShowcasesRequest) (dtousecase.GetSellerShowcasesResponse, error) {
	res := dtousecase.GetSellerShowcasesResponse{}
	res.Showcases = []dtousecase.SellerShowcase{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	rRes, err := u.accountRepository.FindSellerShowcases(ctx, dtorepository.FindSellerShowcasesRequest{SellerId: seller.Id})
	if err != nil {
		return res, err
	}

	for _, data := range rRes {
		res.Showcases = append(res.Showcases, dtousecase.SellerShowcase{ShowcaseId: data.ShowcaseId, ShowcaseName: data.ShowcaseName})
	}

	return res, nil
}

func (u *sellerUsecase) GetShowcaseProducts(ctx context.Context, req dtousecase.GetSellerShowcaseProductRequest) (dtousecase.GetSellerShowcaseProductResponse, error) {
	res := dtousecase.GetSellerShowcaseProductResponse{}
	res.SellerProducts = []dtousecase.SellerProduct{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	rRes, err := u.accountRepository.FindSellerShowcaseProduct(ctx, dtorepository.FindSellerShowcaseProductRequest{ShopName: req.ShopName, ShowcaseId: req.ShowcaseId, Page: req.Page, Limit: req.Limit})

	if err != nil {
		return res, err
	}

	res.SellerProducts = rRes.SellerProducts
	res.Limit = rRes.Limit
	res.CurrentPage = rRes.CurrentPage
	res.TotalItem = rRes.TotalItem
	res.TotalPage = rRes.TotalPage

	return res, nil
}

func (u *sellerUsecase) AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error) {
	res := dtousecase.AddNewProductResponse{}

	productVariants := []dtousecase.ProductVariants{}

	if len(req.Variants) == 1 {
		switch {
		case req.Variants[0].Variant1.Name == "":
			productVariants = []dtousecase.ProductVariants{
				{
					Name: constant.DefaultReservedKeyword,
				},
			}
			req.Variants[0].Variant1.Name = constant.DefaultReservedKeyword
			req.Variants[0].Variant1.Value = constant.DefaultReservedKeyword
		case req.Variants[0].Variant2.Name == "":
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
			}
		default:
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
				{
					Name: req.Variants[0].Variant2.Name,
				},
			}
		}
	}

	if len(req.Variants) > 1 {
		if req.Variants[0].Variant2.Name == "" {
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
			}
		} else {
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
				{
					Name: req.Variants[0].Variant2.Name,
				},
			}
		}
	}

	imageLinks := []string{}

	for _, img := range req.Images {
		currentTime := time.Now().UnixNano()
		fileExtension := path.Ext(img.Filename)
		originalFilename := img.Filename[:len(img.Filename)-len(fileExtension)]
		newFilename := fmt.Sprintf("%s_%d", originalFilename, currentTime)

		file, err := img.Open()
		if err != nil {
			return res, err
		}

		imageUrl, err := util.UploadToCloudinary(file, newFilename)
		if err != nil {
			return res, err
		}

		imageLinks = append(imageLinks, imageUrl)
	}

	product, err := u.productRepository.AddNewProduct(ctx, dtorepository.AddNewProductRequest{
		SellerID:          req.SellerID,
		ProductName:       req.ProductName,
		Description:       req.Description,
		CategoryID:        req.CategoryID,
		HazardousMaterial: req.HazardousMaterial,
		IsNew:             req.IsNew,
		InternalSKU:       req.InternalSKU,
		Weight:            req.Weight,
		IsActive:          req.IsActive,
		Size:              req.Size,
		Variants:          req.Variants,
		ProductVariants:   productVariants,
		Images:            imageLinks,
		VideoURL:          req.VideoURL,
	})
	if err != nil {
		return res, err
	}

	res.ProductName = product.Name

	return res, nil
}

func (u *sellerUsecase) UploadPhoto(ctx context.Context, req dtousecase.UploadNewPhoto) (dtousecase.UploadNewPhoto, error) {
	res := dtousecase.UploadNewPhoto{}

	err := os.MkdirAll("./imageuploads", os.ModePerm)
	if err != nil {
		return res, err
	}

	dst, err := os.Create(fmt.Sprintf("./imageuploads/%s.jpeg", req.ImageID))
	if err != nil {
		return res, err
	}

	defer dst.Close()

	fileImg, err := util.ConvertToJPEG(req.Image)
	if err != nil {
		return res, err
	}

	_, err = io.Copy(dst, fileImg)
	if err != nil {
		return res, err
	}

	res.ImageID = req.ImageID
	return res, nil
}

func (u *sellerUsecase) DeleteProduct(ctx context.Context, req dtousecase.RemoveProduct) (dtousecase.RemoveProduct, error) {
	res := dtousecase.RemoveProduct{}

	product, err := u.productRepository.FindByIDAndSeller(ctx, dtorepository.ProductRequest{ProductID: req.ID, AccountId: req.SellerID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return res, util.ErrProductNotFound
	}
	if err != nil {
		return res, err
	}

	_, err = u.productRepository.RemoveProductByID(ctx, dtorepository.RemoveProduct{ID: req.ID, SellerID: req.SellerID})
	if err != nil {
		return res, err
	}

	res.Name = product.Name

	return res, nil
}

func (u *sellerUsecase) GetProducts(ctx context.Context, req dtousecase.ProductListParam) (*[]dtorepository.ProductListSellerResponse, *dtogeneral.PaginationData, error) {
	uReq := dtorepository.ProductListParam{
		CategoryId: req.CategoryId,
		SellerID:   req.SellerID,
		SortBy:     req.SortBy,
		Search:     req.Search,
		Sort:       req.Sort,
		MinRating:  req.MinRating,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		District:   req.District,
		Limit:      req.Limit,
		Page:       req.Page,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}

	listProduct, totalItems, err := u.productRepository.FindSellerProducts(ctx, uReq)
	if err != nil {
		return &listProduct, nil, err
	}

	pagination := dtogeneral.PaginationData{
		TotalItem:   int(totalItems),
		TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}

	return &listProduct, &pagination, nil
}

func (u *sellerUsecase) GetProductByID(ctx context.Context, req dtousecase.GetProductDetailRequest) (*dtousecase.GetProductSellerResponse, error) {
	res := &dtousecase.GetProductSellerResponse{}

	rReq := dtorepository.ProductRequest{
		ProductID: req.ProductId,
		AccountId: req.AccountId,
	}
	product, err := u.productRepository.FindByIDAndSeller(ctx, rReq)
	if errors.Is(err, util.ErrNoRecordFound) {
		return res, util.ErrProductNotFound
	}
	if err != nil {
		return res, err
	}

	rReq2 := dtorepository.FindProductVariantRequest{
		ProductId: product.ID,
	}
	rRes2, err := u.productRepository.FindProductVariant(ctx, rReq2)
	if err != nil {
		return res, err
	}

	variants, err := u.convertProductVariants(ctx, product.Name, rRes2)
	if err != nil {
		return res, err
	}

	options, err := u.convertVariantOptions(ctx, rRes2)
	if err != nil {
		return res, err
	}

	productImages, err := u.productRepository.FindProductImages(ctx, dtorepository.ProductRequest{ProductID: product.ID})
	if err != nil {
		return res, err
	}

	images := []string{}

	for _, img := range productImages {
		images = append(images, img.URL)
	}

	category, err := u.accountRepository.FindCategoryByID(ctx, product.CategoryID)
	if err != nil {
		return res, err
	}

	productCategory := dtousecase.ProductCategory{
		Id:   category.CategoryLevel1Id,
		Name: category.CategoryLevel1Name,
		Children: dtousecase.ProductCategoryChildren2{
			Id:   category.CategoryLevel2Id,
			Name: category.CategoryLevel2Name,
			Children: dtousecase.ProductCategoryChildren3{
				Id:   category.CategoryLevel3Id,
				Name: category.CategoryLevel3Name,
			},
		},
	}

	res.Id = product.ID
	res.ProductName = product.Name
	res.Description = product.Description
	res.Variants = variants
	res.VariantOptions = options
	res.HazardousMaterial = product.HazardousMaterial
	res.IsActive = product.IsActive
	res.IsNew = product.IsNew
	res.Size = product.Size
	res.Weight = product.Weight
	res.InternalSKU = product.InternalSKU
	res.VideoURL = product.VideoURL
	res.Images = images
	res.Category = productCategory

	return res, nil
}

func (u *sellerUsecase) convertProductVariants(ctx context.Context, productName string, req dtorepository.FindProductVariantResponse) ([]dtousecase.ProductVariantSeller, error) {
	res := []dtousecase.ProductVariantSeller{}

	for _, data := range req.Variants {
		pv := dtousecase.ProductVariantSeller{}
		pv.VariantId = data.VariantId
		pv.Price = data.Price
		pv.Stock = data.Stock
		pv.ImageURL = data.ImageURL

		pv.Selections = append(pv.Selections, dtousecase.ProductSelection{SelectionName: data.SelectionName1, SelectionVariantName: data.VariantName1})
		if data.SelectionId2 != 0 {
			pv.Selections = append(pv.Selections, dtousecase.ProductSelection{SelectionName: data.SelectionName2, SelectionVariantName: data.VariantName2})
		}

		res = append(res, pv)
	}

	return res, nil
}

func (u *sellerUsecase) convertVariantOptions(ctx context.Context, req dtorepository.FindProductVariantResponse) ([]dtousecase.VariantOption, error) {
	res := []dtousecase.VariantOption{}
	m := map[string]map[string]string{}

	for _, data := range req.Variants {
		if data.SelectionName1 != "default_reserved_keyword" {
			if m[data.VariantName1] != nil {
				if m[data.VariantName1][data.SelectionName1] == "" {
					m[data.VariantName1][data.SelectionName1] = data.SelectionName1
				}
			} else {
				m[data.VariantName1] = map[string]string{
					data.SelectionName1: data.SelectionName1,
				}
			}

			if data.SelectionId2 != 0 {
				if m[data.VariantName2] != nil {
					if m[data.VariantName2][data.SelectionName2] == "" {
						m[data.VariantName2][data.SelectionName2] = data.SelectionName2
					}
				} else {
					m[data.VariantName2] = map[string]string{
						data.SelectionName2: data.SelectionName2,
					}
				}
			}
		}
	}

	for key, value := range m {
		vos := []string{}

		for key2 := range value {
			vos = append(vos, key2)
		}

		res = append(res, dtousecase.VariantOption{
			VariantOptionName: key,
			Childs:            vos,
		})
	}

	return res, nil
}

func (u *sellerUsecase) WithdrawSalesBalance(ctx context.Context, req dtousecase.WithdrawBalance) (dtousecase.WithdrawBalance, error) {
	res := dtousecase.WithdrawBalance{}

	log.Println("SellerID", req.SellerID)
	orders, err := u.orderRepository.FindOrderByIDAndSellerID(ctx, dtorepository.ProductOrderRequest{ID: req.OrderID, SellerID: req.SellerID})
	if err != nil {
		return res, err
	}

	if len(orders) < 1 {
		return res, util.ErrOrderDetailNotFound
	}

	if orders[0].Status != constant.StatusOrderCompleted {
		return res, util.ErrOrderNotCompleted
	}

	var totalAmount decimal.Decimal

	for _, o := range orders {
		qty, err := decimal.NewFromString(fmt.Sprintf("%v", o.Quantity))
		if err != nil {
			return res, err
		}

		totalAmount = totalAmount.Add(o.IndividualPrice.Mul(qty))
	}

	buyerWallet, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: orders[0].AccountID})
	if err != nil {
		return res, err
	}

	_, err = u.accountRepository.TopUpWalletBalanceByID(ctx, dtorepository.TopUpWalletRequest{
		UserID:  req.SellerID,
		Type:    constant.Withdraw,
		From:    buyerWallet.WalletNumber,
		Amount:  totalAmount,
		OrderID: orders[0].ID,
	})
	if err != nil {
		return res, err
	}

	res.Balance = totalAmount
	res.OrderID = orders[0].ID

	return res, nil
}
