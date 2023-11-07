package usecase

import (
	"context"
	"errors"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type MyWalletTransactionUsecase interface {
	GetTransactions(ctx context.Context, req dtousecase.WalletHistoriesParams) ([]model.MyWalletTransactionHistories, *dtogeneral.PaginationData, error)
}

type myWalletTransactionUsecase struct {
	walletTransactionRepo repository.WalletTransactionHistoryRepository
	accountRepository     repository.AccountRepository
}

type MyWalletTransactionUsecaseConfig struct {
	WalletTransactionRepo repository.WalletTransactionHistoryRepository
	AccountRepository     repository.AccountRepository
}

func NewMyWalletTransactionUsecase(config MyWalletTransactionUsecaseConfig) MyWalletTransactionUsecase {
	au := &myWalletTransactionUsecase{}
	if config.WalletTransactionRepo != nil {
		au.walletTransactionRepo = config.WalletTransactionRepo
	}
	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
	}

	return au
}

func (u *myWalletTransactionUsecase) GetTransactions(ctx context.Context, req dtousecase.WalletHistoriesParams) ([]model.MyWalletTransactionHistories, *dtogeneral.PaginationData, error) {
	userAccount, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: req.AccountID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, nil, err
	}

	if userAccount.WalletPin == "" {
		return nil, nil, util.ErrWalletNotSet
	}

	transactions, totalItems, err := u.walletTransactionRepo.FindAllTransactions(ctx, dtorepository.WalletHistoriesParams{
		AccountID:       req.AccountID,
		SortBy:          req.SortBy,
		Sort:            req.Sort,
		Limit:           req.Limit,
		Page:            req.Page,
		StartDate:       req.StartDate,
		EndDate:         req.EndDate,
		TransactionType: req.TransactionType,
	})
	if err != nil {
		return nil, nil, err
	}

	pagination := dtogeneral.PaginationData{
		TotalItem:   int(totalItems),
		TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}

	return transactions, &pagination, nil
}
