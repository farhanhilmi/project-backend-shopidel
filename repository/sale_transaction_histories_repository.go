package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type saleWalletTransactionHistoryRepository struct {
	db *gorm.DB
}

type SaleWalletTransactionHistoryRepository interface {
	CreateWithTx(ctx context.Context, tx *gorm.DB, req model.SaleWalletTransactionHistories) (dtorepository.SaleWalletTransactionHistoriesResponse, error)
}

func NewSaleWalletTransactionHistoryRepository(db *gorm.DB) SaleWalletTransactionHistoryRepository {
	return &saleWalletTransactionHistoryRepository{
		db: db,
	}
}

func (r *saleWalletTransactionHistoryRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, req model.SaleWalletTransactionHistories) (dtorepository.SaleWalletTransactionHistoriesResponse, error) {
	res := dtorepository.SaleWalletTransactionHistoriesResponse{}

	err := tx.WithContext(ctx).Model(&model.SaleWalletTransactionHistories{}).Create(&req).Scan(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	return res, err
}
