package usecase

import (
	"context"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type AccountUsecase interface {
	ActivateMyWallet(ctx context.Context, userId int, walletPin string) (*dto.AccountResponse, error)
	CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (dtousecase.CreateAccountResponse, error)
	ChangeMyWalletPIN(ctx context.Context, walletReq dtousecase.UpdateWalletPINRequest) (*dtousecase.UpdateWalletPINResponse, error)
	CheckPasswordCorrect(ctx context.Context, accountReq dtousecase.AccountRequest) (*dtousecase.CheckPasswordResponse, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
}

type AccountUsecaseConfig struct {
	AccountRepository repository.AccountRepository
}

func NewAccountUsecase(config AccountUsecaseConfig) AccountUsecase {
	au := &accountUsecase{}
	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
	}

	return au
}

func (u *accountUsecase) CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (dtousecase.CreateAccountResponse, error) {
	res := dtousecase.CreateAccountResponse{}

	rReq := dtorepository.CreateAccountRequest{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	rRes, err := u.accountRepository.Create(ctx, rReq)

	if err != nil {
		return res, err
	}

	res.Email = rRes.Email
	res.FullName = rRes.FullName
	res.Username = rRes.Username

	return res, nil
}

func (u *accountUsecase) ActivateMyWallet(ctx context.Context, userId int, walletPin string) (*dto.AccountResponse, error) {
	if len(walletPin) != 6 {
		return nil, util.ErrBadPIN
	}

	userAccount, err := u.accountRepository.FindById(ctx, userId)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if userAccount.WalletPin != "" {
		return nil, util.ErrWalletAlreadySet
	}
	acc, err := u.accountRepository.ActivateWalletByID(ctx, userId, walletPin)
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

func (u *accountUsecase) ChangeMyWalletPIN(ctx context.Context, walletReq dtousecase.UpdateWalletPINRequest) (*dtousecase.UpdateWalletPINResponse, error) {
	userAccount, err := u.accountRepository.FindById(ctx, walletReq.UserID)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if userAccount.WalletPin == "" {
		return nil, util.ErrWalletNotSet
	}

	if len(walletReq.WalletNewPIN) != 6 {
		return nil, util.ErrBadPIN
	}

	if userAccount.WalletPin != walletReq.WalletPIN {
		return nil, util.ErrWalletPINNotMatch
	}

	if len(walletReq.WalletNewPIN) != 6 {
		return nil, util.ErrBadPIN
	}

	if userAccount.WalletPin == walletReq.WalletNewPIN {
		return nil, util.ErrSameWalletPIN
	}

	acc, err := u.accountRepository.UpdateWalletPINByID(ctx, dtorepository.UpdateWalletPINRequest{
		UserID:       walletReq.UserID,
		WalletNewPIN: walletReq.WalletNewPIN,
	})

	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	return &dtousecase.UpdateWalletPINResponse{
		WalletNewPIN: acc.WalletNewPIN,
	}, nil
}

func (u *accountUsecase) CheckPasswordCorrect(ctx context.Context, accountReq dtousecase.AccountRequest) (*dtousecase.CheckPasswordResponse, error) {
	userAccount, err := u.accountRepository.FindById(ctx, accountReq.ID)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	return &dtousecase.CheckPasswordResponse{
		IsCorrect: util.CheckPasswordHash(accountReq.Password, userAccount.Password),
	}, nil
}
