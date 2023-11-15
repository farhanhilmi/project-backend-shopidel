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
	res := dtorepository.ProductCombinationVariantRespponse{}

	q := `
	select pvsc.id, pvsc.product_id, pvsc.price as individual_price,pvsc.stock, pvs1.name as variant_name1, pvs2.name as variant_name2 from product_variant_selection_combinations pvsc
		left join product_variant_selections pvs1 
			on pvs1.id = pvsc.product_variant_selection_id1
		left join product_variant_selections pvs2 
			on pvs2.id = pvsc.product_variant_selection_id2
	where pvsc.id = ?
	`

	query := r.db.WithContext(ctx).Table("(?) as t", gorm.Expr(q, req.ID))
	err := query.Model(&model.ProductVariantSelectionCombinations{}).First(&res).Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, err
}
