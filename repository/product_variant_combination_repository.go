package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"gorm.io/gorm"
)

type productVariantCombinationRepository struct {
	db *gorm.DB
}

type ProductVariantCombinationRepository interface {
	IncreaseStockWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductCombinationVariant) (model.ProductCombinationVariant, error)
}

func NewProductVariantCombinationRepository(db *gorm.DB) ProductVariantCombinationRepository {
	return &productVariantCombinationRepository{
		db: db,
	}
}

func (r *productVariantCombinationRepository) IncreaseStockWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductCombinationVariant) (model.ProductCombinationVariant, error) {
	res := model.ProductCombinationVariant{}

	for _, v := range req {
		err := tx.WithContext(ctx).Model(&model.ProductVariantSelectionCombinations{}).Where("id = ?", v.ID).Update("stock", v.Stock).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
