package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type SellerUsecase interface {
	GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error)
	GetBestSelling(ctx context.Context, req dtousecase.GetSellerProductsRequest) (dtousecase.GetSellerProductsResponse, error)
	GetShowcases(ctx context.Context, req dtousecase.GetSellerShowcasesRequest) (dtousecase.GetSellerShowcasesResponse, error)
	GetShowcaseProducts(ctx context.Context, req dtousecase.GetSellerShowcaseProductRequest) (dtousecase.GetSellerShowcaseProductResponse, error)
	AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error)
	UploadPhoto(ctx context.Context, req dtousecase.UploadNewPhoto) (dtousecase.UploadNewPhoto, error)
	DeleteProduct(ctx context.Context, req dtousecase.RemoveProduct) (dtousecase.RemoveProduct, error)
}

type sellerUsecase struct {
	accountRepository repository.AccountRepository
	productRepository repository.ProductRepository
}

type SellerUsecaseConfig struct {
	AccountRepository repository.AccountRepository
	ProductRepository repository.ProductRepository
}

func NewSellerUsecase(config SellerUsecaseConfig) SellerUsecase {
	au := &sellerUsecase{}

	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
	}
	if config.ProductRepository != nil {
		au.productRepository = config.ProductRepository
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
