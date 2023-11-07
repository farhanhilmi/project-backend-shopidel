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
	CreateWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.SaleWalletTransactionHistoriesRequest) (dtorepository.SaleWalletTransactionHistoriesResponse, error)
}

func NewSaleWalletTransactionHistoryRepository(db *gorm.DB) SaleWalletTransactionHistoryRepository {
	return &saleWalletTransactionHistoryRepository{
		db: db,
	}
}

func (r *saleWalletTransactionHistoryRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.SaleWalletTransactionHistoriesRequest) (dtorepository.SaleWalletTransactionHistoriesResponse, error) {
	res := dtorepository.SaleWalletTransactionHistoriesResponse{}

	walletTx := model.SaleWalletTransactionHistories{
		AccountID:      req.AccountID,
		Type:           req.Type,
		Amount:         req.Amount,
		ProductOrderID: req.ProductOrderID,
		To:             req.To,
	}

	err := tx.WithContext(ctx).Model(&walletTx).Create(&walletTx).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.AccountID = walletTx.AccountID
	res.Amount = walletTx.Amount
	res.Type = walletTx.Type
	res.ID = walletTx.ID
	res.ProductOrderID = walletTx.ProductOrderID

	return res, err
}
