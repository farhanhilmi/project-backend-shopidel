package repository

import (
	"context"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"gorm.io/gorm"
)

type productDetailRepository struct {
	db *gorm.DB
}

type ProductOrderDetailRepository interface {
	CreateWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductOrderDetails) (dtorepository.ProductOrderDetailResponse, error)
}

func NewProductOrderDetailRepository(db *gorm.DB) ProductOrderDetailRepository {
	return &productDetailRepository{
		db: db,
	}
}

func (r *productDetailRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, req []model.ProductOrderDetails) (dtorepository.ProductOrderDetailResponse, error) {
	res := dtorepository.ProductOrderDetailResponse{}

	err := tx.WithContext(ctx).CreateInBatches(&req, len(req)).Error

	if err != nil {
		return res, err
	}

	return res, err
}
