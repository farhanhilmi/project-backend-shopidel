package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type walletTransactionHistoryRepository struct {
	db *gorm.DB
}

type WalletTransactionHistoryRepository interface {
	CreateWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletTransactionHistoriesRequest) (dtorepository.MyWalletTransactionHistoriesResponse, error)
}

func NewWalletTransactionHistoryRepository(db *gorm.DB) WalletTransactionHistoryRepository {
	return &walletTransactionHistoryRepository{
		db: db,
	}
}

func (r *walletTransactionHistoryRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletTransactionHistoriesRequest) (dtorepository.MyWalletTransactionHistoriesResponse, error) {
	res := dtorepository.MyWalletTransactionHistoriesResponse{}

	walletTx := model.MyWalletTransactionHistories{
		AccountID:      req.AccountID,
		Type:           req.Type,
		Amount:         req.Amount,
		ProductOrderID: req.ProductOrderID,
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
