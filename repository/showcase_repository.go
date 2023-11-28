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

type ShowcaseRepository interface {
	CreateShowcase(ctx context.Context, req model.SellerShowcase) (model.SellerShowcase, error)
	DeleteShowcase(ctx context.Context, showcaseId int, shopId int) error
	FindShowcases(ctx context.Context, req dtorepository.FindShowcaseRequest) (dtorepository.FindShowcaseResponse, error)
	FindShowcase(ctx context.Context, showcaseId int) (model.SellerShowcase, error)
	UpdateShowcase(ctx context.Context, req model.SellerShowcase) (model.SellerShowcase, error)
}

type showcaseRepository struct {
	db *gorm.DB
}

func NewShowcaseRepository(db *gorm.DB) ShowcaseRepository {
	return &showcaseRepository{
		db: db,
	}
}

func (r *showcaseRepository) CreateShowcase(ctx context.Context, req model.SellerShowcase) (model.SellerShowcase, error) {
	err := r.db.WithContext(ctx).Create(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (r *showcaseRepository) DeleteShowcase(ctx context.Context, showcaseId int, shopId int) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sp := model.SellerShowcase{}
		if err := tx.Where("id = ? and seller_id = ? and deleted_at is null", showcaseId, shopId).First(&sp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return util.ErrShowcaseNotFound
			}

			return err
		}

		sp.DeletedAt = time.Now()

		if err := tx.Updates(&sp).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *showcaseRepository) FindShowcases(ctx context.Context, req dtorepository.FindShowcaseRequest) (dtorepository.FindShowcaseResponse, error) {
	res := dtorepository.FindShowcaseResponse{}
	sp := []model.SellerShowcase{}
	limit := 10
	type count struct{ Count int }
	c := count{}

	q := `
		select
			count(ss.id) as "Count"
		from seller_showcases ss
		where seller_id = ? and deleted_at is null
	`

	if err := r.db.WithContext(ctx).Raw(q, req.ShopId).Scan(&c).Error; err != nil {
		return res, err
	}

	err := r.db.WithContext(ctx).Where("seller_id = ? and deleted_at is null", req.ShopId).Limit(limit).Offset((req.Page - 1) * limit).Find(&sp).Error
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

func (r *showcaseRepository) FindShowcase(ctx context.Context, showcaseId int) (model.SellerShowcase, error) {
	res := model.SellerShowcase{}

	err := r.db.WithContext(ctx).Where("id = ? and deleted_at is null", showcaseId).Preload("ShowcaseProducts").Preload("ShowcaseProducts.Product").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *showcaseRepository) UpdateShowcase(ctx context.Context, req model.SellerShowcase) (model.SellerShowcase, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sp := model.SellerShowcase{}
		if err := tx.Where("id = ?", req.ID).First(&sp).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return util.ErrShopPromotionNotFound
			}

			return err
		}

		sps := []model.SellerShowcaseProduct{}
		if err := tx.Where("seller_showcase_id = ?", req.ID).Find(&sps).Error; err != nil {
			return err
		}

		sp.Name = req.Name
		sp.ShowcaseProducts = req.ShowcaseProducts

		if err := tx.Delete(&sps).Error; err != nil {
			return err
		}

		if err := tx.Updates(&sp).Error; err != nil {
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
