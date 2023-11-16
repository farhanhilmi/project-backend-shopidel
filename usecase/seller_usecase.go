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
	GetShowcases(ctx context.Context, req dtousecase.GetSellerShowcasesRequest) (dtousecase.GetSellerShowcasesResponse, error)
	GetShowcaseProducts(ctx context.Context, req dtousecase.GetSellerShowcaseProductRequest) (dtousecase.GetSellerShowcaseProductResponse, error)
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

	products, err := u.accountRepository.FindSellerShowcaseProduct(ctx, dtorepository.FindSellerShowcaseProductRequest{ShopName: req.ShopName, ShowcaseId: req.ShowcaseId})
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
