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
	CreateWithTx(ctx context.Context, tx *gorm.DB, req model.MyWalletTransactionHistories) (dtorepository.MyWalletTransactionHistoriesResponse, error)
	FindAllTransactions(ctx context.Context, req dtorepository.WalletHistoriesParams) ([]model.MyWalletTransactionHistories, int64, error)
}

func NewWalletTransactionHistoryRepository(db *gorm.DB) WalletTransactionHistoryRepository {
	return &walletTransactionHistoryRepository{
		db: db,
	}
}

func (r *walletTransactionHistoryRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, req model.MyWalletTransactionHistories) (dtorepository.MyWalletTransactionHistoriesResponse, error) {
	res := dtorepository.MyWalletTransactionHistoriesResponse{}

	// walletTx := model.MyWalletTransactionHistories{
	// 	AccountID:      req.AccountID,
	// 	Type:           req.Type,
	// 	Amount:         req.Amount,
	// 	ProductOrderID: req.ProductOrderID,
	// 	From:           req.From,
	// }

	err := tx.WithContext(ctx).Model(&model.MyWalletTransactionHistories{}).Create(&req).Scan(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	// res.AccountID = walletTx.AccountID
	// res.Amount = walletTx.Amount
	// res.Type = walletTx.Type
	// res.ID = walletTx.ID
	// res.ProductOrderID = walletTx.ProductOrderID

	return res, err
}

func (r *walletTransactionHistoryRepository) FindAllTransactions(ctx context.Context, req dtorepository.WalletHistoriesParams) ([]model.MyWalletTransactionHistories, int64, error) {
	transactions := []model.MyWalletTransactionHistories{}
	var totalItems int64

	subQuery := r.db.WithContext(ctx).Model(&transactions).Where("account_id", req.AccountID)
	query := r.db.WithContext(ctx).Table("(?) as t", subQuery)

	if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	}

	if req.EndDate != "" {
		req.EndDate += " 23:59:59"
		query = query.Where("created_at <= ?", req.EndDate)
	}

	if req.TransactionType != "" {
		query = query.Where("type ilike ?", req.TransactionType)
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order(req.SortBy + " " + req.Sort)
	offset := (req.Page - 1) * req.Limit
	query = query.Offset(offset).Limit(req.Limit)

	if err := query.Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, totalItems, nil
}
