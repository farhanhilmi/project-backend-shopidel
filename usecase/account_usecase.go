package usecase

import (
	"context"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

type AccountUsecase interface {
	ActivateMyWallet(ctx context.Context, userId int, walletPin string) (*dto.AccountResponse, error)
}

type AccountUsecaseConfig struct {
	AccountRepo repository.AccountRepository
}

func NewAccountUsecase(config AccountUsecaseConfig) AccountUsecase {
	au := &accountUsecase{}
	if config.AccountRepo != nil {
		au.accountRepo = config.AccountRepo
	}

	return au
}

func (u *accountUsecase) ActivateMyWallet(ctx context.Context, userId int, walletPin string) (*dto.AccountResponse, error) {
	if len(walletPin) != 6 {
		return nil, util.ErrBadPIN
	}

	userAccount, err := u.accountRepo.FindById(ctx, userId)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if userAccount.WalletPin != "" {
		return nil, util.ErrWalletAlreadySet
	}
	acc, err := u.accountRepo.ActivateWalletByID(ctx, userId, walletPin)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	account := dto.AccountResponse{
		ID:           acc.ID,
		WalletNumber: acc.WalletNumber,
	}

	return &account, nil
}
