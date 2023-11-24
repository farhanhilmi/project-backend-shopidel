package usecase

import (
	"context"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
)

type ShowcaseUsecase interface {
	CreateShopPromotions(ctx context.Context, req dtousecase.CreateShowcaseRequest) (dtousecase.CreateShowcaseResponse, error)
	DeleteShopPromotions(ctx context.Context, showcaseId int, shopId int) error
	GetShowcases(ctx context.Context, req dtousecase.GetShowcasesRequest) (dtousecase.GetShowcasesResponse, error)
	GetShowcaseDetail(ctx context.Context, showcaseId int) (dtousecase.GetShowcaseDetailResponse, error)
	UpdateShowcase(ctx context.Context, req dtousecase.UpdateShowcaseRequest) (dtousecase.UpdateShowcaseResponse, error)
}

type showcaseUsecase struct {
	showcaseRepository repository.ShowcaseRepository
}

type ShowcaseUsecaseConfig struct {
	ShowcaseRepository repository.ShowcaseRepository
}

func NewShowcaseUsecase(config ShowcaseUsecaseConfig) ShowcaseUsecase {
	au := &showcaseUsecase{}

	if config.ShowcaseRepository != nil {
		au.showcaseRepository = config.ShowcaseRepository
	}

	return au
}

func (u *showcaseUsecase) CreateShopPromotions(ctx context.Context, req dtousecase.CreateShowcaseRequest) (dtousecase.CreateShowcaseResponse, error) {
	res := dtousecase.CreateShowcaseResponse{}
	res.SelectedProductsId = []int{}

	sps := []model.SellerShowcaseProduct{}

	for _, selectedProductId := range req.SelectedProductsId {
		sps = append(sps, model.SellerShowcaseProduct{
			ProductId: selectedProductId,
		})
	}

	sp := model.SellerShowcase{
		Name:             req.Name,
		SellerId:         req.ShopId,
		ShowcaseProducts: sps,
	}

	rRes, err := u.showcaseRepository.CreateShowcase(ctx, sp)
	if err != nil {
		return res, err
	}

	res = dtousecase.CreateShowcaseResponse{
		Id:                 rRes.ID,
		Name:               rRes.Name,
		ShopId:             rRes.SellerId,
		SelectedProductsId: req.SelectedProductsId,
		CreatedAt:          rRes.CreatedAt,
	}

	return res, nil
}

func (u *showcaseUsecase) DeleteShopPromotions(ctx context.Context, showcaseId int, shopId int) error {
	err := u.showcaseRepository.DeleteShowcase(ctx, showcaseId, shopId)
	if err != nil {
		return err
	}

	return nil
}

func (u *showcaseUsecase) GetShowcases(ctx context.Context, req dtousecase.GetShowcasesRequest) (dtousecase.GetShowcasesResponse, error) {
	res := dtousecase.GetShowcasesResponse{}
	res.Showcases = []dtousecase.Showcase{}

	rRes, err := u.showcaseRepository.FindShowcases(ctx, dtorepository.FindShowcaseRequest{ShopId: req.ShopId, Page: req.Page})
	if err != nil {
		return res, err
	}

	res.CurrentPage = rRes.CurrentPage
	res.Limit = rRes.Limit
	res.TotalItems = rRes.TotalItems
	res.TotalPages = rRes.TotalPages

	for _, mp := range rRes.ShopPromotions {
		res.Showcases = append(res.Showcases, dtousecase.Showcase{
			ID:   mp.ID,
			Name: mp.Name,
		})
	}

	return res, nil
}

func (u *showcaseUsecase) GetShowcaseDetail(ctx context.Context, showcaseId int) (dtousecase.GetShowcaseDetailResponse, error) {
	res := dtousecase.GetShowcaseDetailResponse{}
	selectedProducts := []dtousecase.ShowcaseSelectedProduct{}

	rRes, err := u.showcaseRepository.FindShowcase(ctx, showcaseId)
	if err != nil {
		return res, nil
	}

	for _, selectedProduct := range rRes.ShowcaseProducts {
		selectedProducts = append(selectedProducts, dtousecase.ShowcaseSelectedProduct{
			ProductId:   selectedProduct.ProductId,
			ProductName: selectedProduct.Product.Name,
			CreatedAt:   selectedProduct.CreatedAt,
		})
	}

	res = dtousecase.GetShowcaseDetailResponse{
		ID:               rRes.ID,
		Name:             rRes.Name,
		CreatedAt:        rRes.CreatedAt,
		SelectedProducts: selectedProducts,
	}
	return res, nil
}

func (u *showcaseUsecase) UpdateShowcase(ctx context.Context, req dtousecase.UpdateShowcaseRequest) (dtousecase.UpdateShowcaseResponse, error) {
	res := dtousecase.UpdateShowcaseResponse{}

	sps := []model.SellerShowcaseProduct{}

	for _, selectedProductId := range req.SelectedProductsId {
		sps = append(sps, model.SellerShowcaseProduct{
			SellerShowcaseId: req.Id,
			ProductId:        selectedProductId,
		})
	}

	sp := model.SellerShowcase{
		ID:               req.Id,
		Name:             req.Name,
		ShowcaseProducts: sps,
	}

	rRes, err := u.showcaseRepository.UpdateShowcase(ctx, sp)
	if err != nil {
		return res, err
	}

	res = dtousecase.UpdateShowcaseResponse{
		Id:                 rRes.ID,
		ShopId:             rRes.SellerId,
		Name:               rRes.Name,
		SelectedProductsId: req.SelectedProductsId,
		CreatedAt:          rRes.CreatedAt,
	}

	return res, nil
}
