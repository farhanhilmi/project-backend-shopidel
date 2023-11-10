package usecase

import (
	"context"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
)

type SellerUsecase interface {
	GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error)
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

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{SellerId: req.SellerId})
	if err != nil {
		return res, err
	}

	rRes, err := u.accountRepository.FindSellerProducts(ctx, dtorepository.FindSellerProductsRequest{SellerId: req.SellerId})
	if err != nil {
		return res, err
	}

	rRes2, err := u.accountRepository.FindSellerSelectedCategories(ctx, dtorepository.FindSellerSelectedCategoriesRequest{SellerId: req.SellerId})
	if err != nil {
		return res, err
	}

	for _, data := range rRes.Products {
		p := dtousecase.SellerProduct{
			Name:           data.Name,
			Price:          data.Price,
			PictureUrl:     data.PictureUrl,
			Stars:          data.Stars,
			TotalSold:      data.TotalSold,
			CreatedAt:      data.CreatedAt,
			CategoryLevel1: data.CategoryLevel1,
			CategoryLevel2: data.CategoryLevel2,
			CategoryLevel3: data.CategoryLevel3,
		}

		res.SellerProducts = append(res.SellerProducts, p)
	}

	for _, data := range rRes2 {
		res.SellerSelectedCategories = append(res.SellerSelectedCategories, dtousecase.SellerSelectedCategory{
			CategoryId:   data.CategoryId,
			CategoryName: data.CategoryName,
		})
	}

	res.SellerId = seller.Id
	res.SellerDistrict = seller.District
	res.SellerName = seller.Name
	res.SellerOperatingHour = dtousecase.SellerOperatingHour{
		Start: seller.StartOperatingHours,
		End:   seller.EndOperatingHours,
	}
	res.SellerPictureUrl = seller.ProfilePicture

	return res, nil
}
