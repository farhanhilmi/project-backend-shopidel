package usecase

import (
	"context"

	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
)

type PromotionUsecase interface {
	GetShopAvailablePromotions(ctx context.Context, req dtousecase.GetShopAvailablePromotionsRequest) (dtousecase.GetShopAvailablePromotionsResponse, error)
	GetMarketplacePromotions(ctx context.Context) (dtousecase.GetMarketplacePromotionsResponse, error)
	GetShopPromotions(ctx context.Context, req dtousecase.GetShopPromotionsRequest) (dtousecase.GetShopPromotionsResponse, error)
	CreateShopPromotions(ctx context.Context, req dtousecase.CreateShopPromotionRequest) (dtousecase.CreateShopPromotionResponse, error)
}

type promotionUsecase struct {
	promotionRepository repository.PromotionRepository
}

type PromotionUsecaseConfig struct {
	PromotionRepository repository.PromotionRepository
}

func NewPromotionUsecase(config PromotionUsecaseConfig) PromotionUsecase {
	au := &promotionUsecase{}

	if config.PromotionRepository != nil {
		au.promotionRepository = config.PromotionRepository
	}

	return au
}

func (u *promotionUsecase) GetShopAvailablePromotions(ctx context.Context, req dtousecase.GetShopAvailablePromotionsRequest) (dtousecase.GetShopAvailablePromotionsResponse, error) {
	res := dtousecase.GetShopAvailablePromotionsResponse{}

	shopPromotions, err := u.promotionRepository.FindShopAvailablePromotions(ctx, req.ShopId)
	if err != nil {
		return res, err
	}

	for _, shopPromotion := range shopPromotions {
		sp := dtousecase.ShopPromotion{
			ID:                 shopPromotion.ID,
			Name:               shopPromotion.Name,
			MinPurchaseAmount:  shopPromotion.MinPurchaseAmount,
			MaxPurchaseAmount:  shopPromotion.MaxPurchaseAmount,
			DiscountPercentage: shopPromotion.DiscountPercentage,
		}

		for _, selectedProduct := range shopPromotion.SelectedProducts {
			sp.SelectedProductsId = append(sp.SelectedProductsId, selectedProduct.ID)
		}

		res.ShopPromotions = append(res.ShopPromotions, sp)
	}

	return res, nil
}

func (u *promotionUsecase) GetMarketplacePromotions(ctx context.Context) (dtousecase.GetMarketplacePromotionsResponse, error) {
	res := dtousecase.GetMarketplacePromotionsResponse{}

	marketplacePromotions, err := u.promotionRepository.FindMarketplacePromotions(ctx)
	if err != nil {
		return res, err
	}

	for _, mp := range marketplacePromotions {
		res.MarketplacePromotions = append(res.MarketplacePromotions, dtousecase.MarketplacePromotion{
			ID:                 mp.ID,
			Name:               mp.Name,
			MaxPurchaseAmount:  mp.MaxPurchaseAmount,
			MinPurchaseAmount:  mp.MinPurchaseAmount,
			DiscountPercentage: mp.DiscountPercentage,
		})
	}

	return res, nil
}

func (u *promotionUsecase) GetShopPromotions(ctx context.Context, req dtousecase.GetShopPromotionsRequest) (dtousecase.GetShopPromotionsResponse, error) {
	res := dtousecase.GetShopPromotionsResponse{}

	marketplacePromotions, err := u.promotionRepository.FindShopPromotions(ctx, req.ShopId)
	if err != nil {
		return res, err
	}

	for _, mp := range marketplacePromotions {
		res.ShopPromotions = append(res.ShopPromotions, dtousecase.ShopPromotion{
			ID:                 mp.ID,
			Name:               mp.Name,
			MaxPurchaseAmount:  mp.MaxPurchaseAmount,
			MinPurchaseAmount:  mp.MinPurchaseAmount,
			DiscountPercentage: mp.DiscountPercentage,
			Quota:              mp.Quota,
			TotalUsed:          mp.TotalUsed,
			StartDate:          mp.StartDate,
			EndDate:            mp.EndDate,
		})
	}

	return res, nil
}

func (u *promotionUsecase) CreateShopPromotions(ctx context.Context, req dtousecase.CreateShopPromotionRequest) (dtousecase.CreateShopPromotionResponse, error) {
	res := dtousecase.CreateShopPromotionResponse{}

	sps := []model.ShopPromotionSelectedProduct{}

	for _, selectedProductId := range req.SelectedProductsId {
		sps = append(sps, model.ShopPromotionSelectedProduct{
			ProductId: selectedProductId,
		})
	}

	sp := model.ShopPromotion{
		ShopId:             req.ShopId,
		Name:               req.Name,
		Quota:              req.Quota,
		StartDate:          req.StartDate,
		EndDate:            req.EndDate,
		MinPurchaseAmount:  req.MinPurchaseAmount,
		MaxPurchaseAmount:  req.MaxPurchaseAmount,
		DiscountPercentage: req.DiscountPercentage,
		SelectedProducts:   sps,
	}

	rRes, err := u.promotionRepository.CreateShopPromotion(ctx, sp)
	if err != nil {
		return res, err
	}

	res = dtousecase.CreateShopPromotionResponse{
		Id:                 rRes.ID,
		ShopId:             rRes.ShopId,
		Name:               rRes.Name,
		Quota:              rRes.Quota,
		StartDate:          rRes.StartDate,
		EndDate:            rRes.EndDate,
		MinPurchaseAmount:  rRes.MinPurchaseAmount,
		MaxPurchaseAmount:  rRes.MaxPurchaseAmount,
		DiscountPercentage: rRes.DiscountPercentage,
		SelectedProductsId: req.SelectedProductsId,
	}

	return res, nil
}
