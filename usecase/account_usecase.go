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
	ActivateMyWallet(ctx context.Context, req dtousecase.GetAccountRequest, walletPin string) (*dto.AccountResponse, error)
	CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (dtousecase.CreateAccountResponse, error)
	GetProfile(ctx context.Context, req dtousecase.GetAccountRequest) (*dtousecase.GetAccountResponse, error)
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

func (u *accountUsecase) GetProfile(ctx context.Context, req dtousecase.GetAccountRequest) (*dtousecase.GetAccountResponse, error) {
	res := dtousecase.GetAccountResponse{}

	rReq := dtorepository.GetAccountRequest{
		UserId: req.UserId,
	}

	userAccount, err := u.accountRepository.FindById(ctx, rReq)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	res.FullName = userAccount.FullName
	res.Username = userAccount.Username
	res.Email = userAccount.Email
	res.PhoneNumber = userAccount.PhoneNumber
	res.Gender = userAccount.Gender
	res.Birthdate = userAccount.Birthdate
	res.ProfilePicture = userAccount.ProfilePicture
	res.WalletNumber = userAccount.WalletNumber
	res.Balance = userAccount.Balance

	return &res, nil
}

func (u *accountUsecase) ActivateMyWallet(ctx context.Context, req dtousecase.GetAccountRequest, walletPin string) (*dto.AccountResponse, error) {
	if len(walletPin) != 6 {
		return nil, util.ErrBadPIN
	}

	rReq := dtorepository.GetAccountRequest{
		UserId: req.UserId,
	}

	userAccount, err := u.accountRepository.FindById(ctx, rReq)
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if userAccount.WalletPin != "" {
		return nil, util.ErrWalletAlreadySet
	}
	acc, err := u.accountRepository.ActivateWalletByID(ctx, req.UserId, walletPin)
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
