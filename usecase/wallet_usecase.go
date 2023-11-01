package usecase

// import (
// 	"context"
// 	"errors"

// 	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
// 	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
// )

// type WalletUsecase interface {
// 	ActivateMyWallet(ctx context.Context, userId int, walletPin string) (*dto.AccountResponse, error)
// 	ChangeMyWalletPIN(ctx context.Context, walletReq dtohttp.ChangeWalletPINRequest) (*dto.AccountResponse, error)
// }

// type walletUsecase struct {
// 	accountRepository repository.AccountRepository
// }

// type WalletUsecaseConfig struct {
// 	AccountRepository repository.AccountRepository
// }

// func NewWalletUsecase(config WalletUsecaseConfig) WalletUsecase {
// 	au := &walletUsecase{}
// 	if config.AccountRepository != nil {
// 		au.accountRepository = config.AccountRepository
// 	}

// 	return au
// }

// func (u *walletUsecase) ActivateMyWallet(ctx context.Context, userId int, walletPin string) (*dto.AccountResponse, error) {
// 	if len(walletPin) != 6 {
// 		return nil, util.ErrBadPIN
// 	}

// 	userAccount, err := u.accountRepository.FindById(ctx, userId)
// 	if errors.Is(err, util.ErrNoRecordFound) {
// 		return nil, util.ErrNoRecordFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	if userAccount.WalletPin != "" {
// 		return nil, util.ErrWalletAlreadySet
// 	}
// 	acc, err := u.accountRepository.ActivateWalletByID(ctx, userId, walletPin)
// 	if errors.Is(err, util.ErrNoRecordFound) {
// 		return nil, util.ErrNoRecordFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	account := dto.AccountResponse{
// 		ID:           acc.ID,
// 		WalletNumber: acc.WalletNumber,
// 	}

// 	return &account, nil
// }

// func (u *accountUsecase) ChangeMyWalletPIN(ctx context.Context, walletReq dtohttp.ChangeWalletPINRequest) (*dto.AccountResponse, error) {
// 	if len(walletReq.WalletNewPIN) != 6 {
// 		return nil, util.ErrBadPIN
// 	}

// 	userAccount, err := u.accountRepository.FindById(ctx, userId)
// 	if errors.Is(err, util.ErrNoRecordFound) {
// 		return nil, util.ErrNoRecordFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	if userAccount.WalletPin != "" {
// 		return nil, util.ErrWalletAlreadySet
// 	}
// 	acc, err := u.accountRepository.ActivateWalletByID(ctx, userId, walletPin)
// 	if errors.Is(err, util.ErrNoRecordFound) {
// 		return nil, util.ErrNoRecordFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	account := dto.AccountResponse{
// 		ID:           acc.ID,
// 		WalletNumber: acc.WalletNumber,
// 	}

// 	return &account, nil
// }
