package repository

import (
	"context"
	"errors"
	"math"
	"time"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	CreateShopPromotion(ctx context.Context, req model.ShopPromotion) (model.ShopPromotion, error)
	DeleteShopPromotion(ctx context.Context, shopPromotionId int, shopId int) error
	FindShopAvailablePromotions(ctx context.Context, shopId int) ([]model.ShopPromotion, error)
	FindMarketplacePromotions(ctx context.Context) ([]model.MarketplacePromotion, error)
	FindShopPromotions(ctx context.Context, req dtorepository.FindShopPromotionsRequest) (dtorepository.FindShopPromotionsResponse, error)
	FindShopPromotion(ctx context.Context, shopPromotionId int) (model.ShopPromotion, error)
	UpdateShopPromotion(ctx context.Context, req model.ShopPromotion) (model.ShopPromotion, error)
}

type promotionRepository struct {
	db *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &promotionRepository{
		db: db,
	}
}

func (r *promotionRepository) CreateShopPromotion(ctx context.Context, req model.ShopPromotion) (model.ShopPromotion, error) {
	err := r.db.WithContext(ctx).Create(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (r *promotionRepository) DeleteShopPromotion(ctx context.Context, shopPromotionId int, shopId int) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sp := model.ShopPromotion{}
		if err := tx.Where("id = ? and shop_id = ? and deleted_at is null", shopPromotionId, shopId).First(&sp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return util.ErrShopPromotionNotFound
			}

			return err
		}

		sp.DeletedAt = time.Now()

		if err := tx.Save(&sp).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *promotionRepository) FindShopAvailablePromotions(ctx context.Context, shopId int) ([]model.ShopPromotion, error) {
	res := []model.ShopPromotion{}

	err := r.db.WithContext(ctx).Where("shop_id = ? and start_date < Now() and end_date > Now() and Quota > 0 and deleted_at is null", shopId).Preload("SelectedProducts").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *promotionRepository) FindMarketplacePromotions(ctx context.Context) ([]model.MarketplacePromotion, error) {
	res := []model.MarketplacePromotion{}

	err := r.db.WithContext(ctx).Where("start_date < Now() and end_date > Now() and Quota > 0 and deleted_at is null").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *promotionRepository) FindShopPromotions(ctx context.Context, req dtorepository.FindShopPromotionsRequest) (dtorepository.FindShopPromotionsResponse, error) {
	res := dtorepository.FindShopPromotionsResponse{}
	sp := []model.ShopPromotion{}
	limit := 10
	type count struct{ Count int }
	c := count{}

	q := `
		select
			count(sp.id) as "Count"
		from shop_promotions sp
		where shop_id = ? and deleted_at is null
	`

	if err := r.db.WithContext(ctx).Raw(q, req.ShopId).Scan(&c).Error; err != nil {
		return res, err
	}

	err := r.db.WithContext(ctx).Where("shop_id = ? and deleted_at is null", req.ShopId).Limit(limit).Offset((req.Page - 1) * limit).Find(&sp).Error
	if err != nil {
		return res, err
	}

	res.Limit = limit
	res.CurrentPage = req.Page
	res.TotalItems = c.Count
	res.TotalPages = int(math.Ceil(float64(c.Count) / float64(res.Limit)))
	res.ShopPromotions = sp

	return res, nil
}

func (r *promotionRepository) FindShopPromotion(ctx context.Context, shopPromotionId int) (model.ShopPromotion, error) {
	res := model.ShopPromotion{}

	err := r.db.WithContext(ctx).Where("id = ? and deleted_at is null", shopPromotionId).Preload("SelectedProducts").Preload("SelectedProducts.Product").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *promotionRepository) UpdateShopPromotion(ctx context.Context, req model.ShopPromotion) (model.ShopPromotion, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sp := model.ShopPromotion{}
		if err := tx.Where("id = ?", req.ID).First(&sp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return util.ErrShopPromotionNotFound
			}

			return err
		}

		sps := []model.ShopPromotionSelectedProduct{}
		if err := tx.Where("shop_promotion_id = ?", req.ID).Find(&sps).Error; err != nil {

			return err
		}

		sp.Name = req.Name
		sp.Quota = req.Quota
		sp.StartDate = req.StartDate
		sp.EndDate = req.EndDate
		sp.MinPurchaseAmount = req.MinPurchaseAmount
		sp.MaxPurchaseAmount = req.MaxPurchaseAmount
		sp.DiscountPercentage = req.DiscountPercentage
		sp.SelectedProducts = req.SelectedProducts

		if err := tx.Delete(&sps).Error; err != nil {
			return err
		}

		if err := tx.Save(&sp).Error; err != nil {
			return err
		}

		req = sp

		return nil
	})

	if err != nil {
		return req, err
	}

	return req, nil
}
