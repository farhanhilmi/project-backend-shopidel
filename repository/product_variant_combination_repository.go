package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type productVariantCombinationRepository struct {
	db *gorm.DB
}

type ProductVariantCombinationRepository interface {
	UpdateStockWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductCombinationVariant) (model.ProductCombinationVariant, error)
	FindProductsByID(ctx context.Context, req dtorepository.ProductCombinationVariantListRequest) ([]dtorepository.ProductVariant, error)
	FindById(ctx context.Context, req dtorepository.ProductCombinationVariantRequest) (dtorepository.ProductCombinationVariantRespponse, error)
	DecreaseStockWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductCombinationVariant) (model.ProductCombinationVariant, error)
}

func NewProductVariantCombinationRepository(db *gorm.DB) ProductVariantCombinationRepository {
	return &productVariantCombinationRepository{
		db: db,
	}
}

func (r *productVariantCombinationRepository) UpdateStockWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductCombinationVariant) (model.ProductCombinationVariant, error) {
	res := model.ProductCombinationVariant{}

	for _, v := range req {
		err := tx.WithContext(ctx).Model(&model.ProductVariantSelectionCombinations{}).Where("id = ?", v.ID).Update("stock", v.Stock).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *productVariantCombinationRepository) DecreaseStockWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductCombinationVariant) (model.ProductCombinationVariant, error) {
	res := model.ProductCombinationVariant{}

	for _, v := range req {
		err := tx.WithContext(ctx).Model(&model.ProductVariantSelectionCombinations{}).Where("id = ?", v.ID).Update("stock", gorm.Expr("stock - ?", v.Stock)).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *productVariantCombinationRepository) FindProductsByID(ctx context.Context, req dtorepository.ProductCombinationVariantListRequest) ([]dtorepository.ProductVariant, error) {
	res := []dtorepository.ProductVariant{}

	err := r.db.WithContext(ctx).Model(&model.ProductVariantSelectionCombinations{}).Where("id IN ?", req.ID).Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	if err != nil {
		return res, err
	}

	return res, err
}

func (r *productVariantCombinationRepository) FindById(ctx context.Context, req dtorepository.ProductCombinationVariantRequest) (dtorepository.ProductCombinationVariantRespponse, error) {
	product := model.ProductVariantSelectionCombinations{}
	res := dtorepository.ProductCombinationVariantRespponse{}

	err := r.db.WithContext(ctx).Where("id = ?", req.ID).First(&product).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	res.ID = product.ID
	res.IndividualPrice = product.Price
	res.Stock = product.Stock
	res.ProductID = product.ProductID

	return res, err
}
