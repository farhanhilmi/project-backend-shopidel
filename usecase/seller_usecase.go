package usecase

import (
	"context"
	"log"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type SellerUsecase interface {
	GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error)
	GetBestSelling(ctx context.Context, req dtousecase.GetSellerProductsRequest) (dtousecase.GetSellerProductsResponse, error)
	GetCategories(ctx context.Context, req dtousecase.GetSellerCategoriesRequest) (dtousecase.GetSellerCategoriesResponse, error)
	GetCategoryProducts(ctx context.Context, req dtousecase.GetSellerCategoryProductRequest) (dtousecase.GetSellerCategoryProductResponse, error)
	AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error)
}

type sellerUsecase struct {
	accountRepository repository.AccountRepository
}

type SellerUsecaseConfig struct {
	AccountRepository repository.AccountRepository
}

func NewSellerUsecase(config SellerUsecaseConfig) SellerUsecase {
	au := &sellerUsecase{}

	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
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
			Name:       data.Name,
			Price:      data.Price,
			PictureUrl: data.PictureUrl,
			Stars:      data.Stars,
			TotalSold:  data.TotalSold,
			CreatedAt:  data.CreatedAt,
			Category:   data.Category,
		}

		res.SellerProducts = append(res.SellerProducts, p)
	}

	return res, nil
}

func (u *sellerUsecase) GetCategories(ctx context.Context, req dtousecase.GetSellerCategoriesRequest) (dtousecase.GetSellerCategoriesResponse, error) {
	res := dtousecase.GetSellerCategoriesResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	rRes, err := u.accountRepository.FindSellerCategories(ctx, dtorepository.FindSellerCategoriesRequest{SellerId: seller.Id})
	if err != nil {
		return res, err
	}

	for _, data := range rRes {
		res.Categories = append(res.Categories, dtousecase.SellerCategory{CategoryId: data.CategoryId, CategoryName: data.CategoryName})
	}

	return res, nil
}

func (u *sellerUsecase) GetCategoryProducts(ctx context.Context, req dtousecase.GetSellerCategoryProductRequest) (dtousecase.GetSellerCategoryProductResponse, error) {
	res := dtousecase.GetSellerCategoryProductResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	products, err := u.accountRepository.FindSellerCategoryProduct(ctx, dtorepository.FindSellerCategoryProductRequest{ShopName: req.ShopName, CategoryId: req.CategoryId})
	if err != nil {
		return res, err
	}

	res.SellerProducts = products

	return res, nil
}

func (u *sellerUsecase) AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error) {
	res := dtousecase.AddNewProductResponse{}

	log.Println("REQ", req)

	for i, header := range req.Images {
		_, err := req.Images[i].Open()
		if err != nil {
			log.Println("ERR", err)
			return res, err
		}
		log.Println("FILENAME", header.Filename)
	}

	return res, nil
}
