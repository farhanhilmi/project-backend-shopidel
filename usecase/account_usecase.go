package usecase

import (
	"context"
	"errors"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type AccountUsecase interface {
	ActivateMyWallet(ctx context.Context, req dtousecase.GetAccountRequest, walletPin string) (*dto.AccountResponse, error)
	CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (*dtousecase.CreateAccountResponse, error)
	ChangeMyWalletPIN(ctx context.Context, walletReq dtousecase.UpdateWalletPINRequest) (*dtousecase.UpdateWalletPINResponse, error)
	CheckPasswordCorrect(ctx context.Context, accountReq dtousecase.AccountRequest) (*dtousecase.CheckPasswordResponse, error)
	GetProfile(ctx context.Context, req dtousecase.GetAccountRequest) (*dtousecase.GetAccountResponse, error)
}

type accountUsecase struct {
	accountRepository   repository.AccountRepository
	usedEmailRepository repository.UsedEmailRepository
}

type AccountUsecaseConfig struct {
	AccountRepository   repository.AccountRepository
	UsedEmailRepository repository.UsedEmailRepository
}

func NewAccountUsecase(config AccountUsecaseConfig) AccountUsecase {
	au := &accountUsecase{}
	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
		au.usedEmailRepository = config.UsedEmailRepository
	}

	return au
}

func (u *accountUsecase) CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (*dtousecase.CreateAccountResponse, error) {
	res := dtousecase.CreateAccountResponse{}

	userAccount, err := u.accountRepository.FindByEmail(ctx, dtorepository.GetAccountRequest{Email: req.Email})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}

	if strings.EqualFold(userAccount.Email, req.Email) {
		return nil, util.ErrEmailAlreadyExist
	}

	if strings.Contains(strings.ToLower(req.Username), strings.ToLower(req.Password)) {
		return nil, util.ErrPasswordContainUsername
	}

	usedEmail, err := u.usedEmailRepository.FindByEmail(ctx, dtorepository.UsedEmailRequest{Email: req.Email})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if usedEmail.Email == req.Email {
		return nil, util.ErrCantUseThisEmail
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	rReq := dtorepository.CreateAccountRequest{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
	}

	rRes, err := u.accountRepository.Create(ctx, rReq)
	if err != nil {
		return nil, err
	}

	res.Email = rRes.Email
	res.FullName = rRes.FullName
	res.Username = rRes.Username

	return &res, nil
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

func (u *accountUsecase) ChangeMyWalletPIN(ctx context.Context, walletReq dtousecase.UpdateWalletPINRequest) (*dtousecase.UpdateWalletPINResponse, error) {
	userAccount, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: walletReq.UserID})
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
	userAccount, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: accountReq.ID})
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
