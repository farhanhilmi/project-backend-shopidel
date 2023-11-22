package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"gorm.io/gorm"
)

type promotionRepository struct {
	db *gorm.DB
}

type PromotionRepository interface {
	FindShopAvailablePromotions(ctx context.Context, shopId int) ([]model.ShopPromotion, error)
	FindMarketplacePromotions(ctx context.Context) ([]model.MarketplacePromotion, error)
	FindShopPromotions(ctx context.Context, shopId int) ([]model.ShopPromotion, error)
	CreateShopPromotion(ctx context.Context, req model.ShopPromotion) (model.ShopPromotion, error)
}

func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &promotionRepository{
		db: db,
	}
}

func (r *promotionRepository) FindShopAvailablePromotions(ctx context.Context, shopId int) ([]model.ShopPromotion, error) {
	res := []model.ShopPromotion{}

	err := r.db.WithContext(ctx).Where("shop_id = ? and start_date < Now() and end_date > Now() and Quota > 0", shopId).Preload("SelectedProducts").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *promotionRepository) FindMarketplacePromotions(ctx context.Context) ([]model.MarketplacePromotion, error) {
	res := []model.MarketplacePromotion{}

	err := r.db.WithContext(ctx).Where("start_date < Now() and end_date > Now() and Quota > 0").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *promotionRepository) FindShopPromotions(ctx context.Context, shopId int) ([]model.ShopPromotion, error) {
	res := []model.ShopPromotion{}

	err := r.db.WithContext(ctx).Where("shop_id = ?", shopId).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *promotionRepository) CreateShopPromotion(ctx context.Context, req model.ShopPromotion) (model.ShopPromotion, error) {
	err := r.db.WithContext(ctx).Create(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}
