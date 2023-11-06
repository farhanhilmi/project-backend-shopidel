package usecase

import (
	"context"
	"errors"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/shopspring/decimal"
)

type AccountUsecase interface {
	ActivateMyWallet(ctx context.Context, req dtousecase.GetAccountRequest, walletPin string) (*dto.AccountResponse, error)
	CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (*dtousecase.CreateAccountResponse, error)
	ChangeMyWalletPIN(ctx context.Context, walletReq dtousecase.UpdateWalletPINRequest) (*dtousecase.UpdateWalletPINResponse, error)
	CheckPasswordCorrect(ctx context.Context, accountReq dtousecase.AccountRequest) (*dtousecase.CheckPasswordResponse, error)
	GetProfile(ctx context.Context, req dtousecase.GetAccountRequest) (*dtousecase.GetAccountResponse, error)
	EditProfile(ctx context.Context, req dtousecase.EditAccountRequest) (*dtousecase.EditAccountResponse, error)
	GetWallet(ctx context.Context, req dtousecase.AccountRequest) (*dtousecase.WalletResponse, error)
	Login(ctx context.Context, req dtousecase.LoginRequest) (*dtousecase.LoginResponse, error)
	TopUpBalanceWallet(ctx context.Context, walletReq dtousecase.TopUpBalanceWalletRequest) (*dtousecase.TopUpBalanceWalletResponse, error)
	GetCart(ctx context.Context, req dtousecase.GetCartRequest) (dtousecase.GetCartResponse, error)
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

func (u *accountUsecase) Login(ctx context.Context, req dtousecase.LoginRequest) (*dtousecase.LoginResponse, error) {
	res := dtousecase.LoginResponse{}

	userAccount, err := u.accountRepository.FindByEmail(ctx, dtorepository.GetAccountRequest{Email: req.Email})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}

	if valid := util.CheckPasswordHash(req.Password, userAccount.Password); !valid {
		return nil, util.ErrInvalidPassword
	}

	role := "buyer"
	if userAccount.ShopName != "" {
		role = "seller"
	}

	token, err := util.GenerateJWT(userAccount.ID, role)
	if err != nil {
		return nil, err
	}

	res.AccessToken = token

	return &res, nil
}

func (u *accountUsecase) CreateAccount(ctx context.Context, req dtousecase.CreateAccountRequest) (*dtousecase.CreateAccountResponse, error) {
	res := dtousecase.CreateAccountResponse{}

	if !util.ValidatePassword(req.Password) {
		return nil, util.ErrWeakPassword
	}

	userAccount, err := u.accountRepository.FindByEmail(ctx, dtorepository.GetAccountRequest{Email: req.Email})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}

	if strings.EqualFold(userAccount.Email, req.Email) {
		return nil, util.ErrEmailAlreadyExist
	}

	uAcc, err := u.accountRepository.FindByUsername(ctx, dtorepository.GetAccountRequest{Username: req.Username})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}

	if strings.EqualFold(uAcc.Username, req.Username) {
		return nil, util.ErrUsernameAlreadyExist
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

func (u *accountUsecase) EditProfile(ctx context.Context, req dtousecase.EditAccountRequest) (*dtousecase.EditAccountResponse, error) {
	res := dtousecase.EditAccountResponse{}

	oldAccount, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: req.UserId})
	if err != nil {
		return &res, err
	}

	usedUsername, err := u.accountRepository.FindByUsername(ctx, dtorepository.GetAccountRequest{
		UserId:   req.UserId,
		Email:    req.Email,
		Username: req.Username,
	})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if usedUsername.Username == req.Username {
		return nil, util.ErrCantUseThisUsername
	}

	usedPhonenumber, err := u.accountRepository.FindByPhoneNumber(ctx, dtorepository.GetAccountRequest{
		UserId:      req.UserId,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if usedPhonenumber.PhoneNumber == req.PhoneNumber {
		return nil, util.ErrCantUseThisPhonenumber
	}

	usedEmail, err := u.usedEmailRepository.FindByEmail(ctx, dtorepository.UsedEmailRequest{Email: req.Email})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if usedEmail.Email == req.Email {
		return nil, util.ErrCantUseThisEmail
	}

	rReq := dtorepository.EditAccountRequest{
		UserId:         req.UserId,
		FullName:       req.FullName,
		Username:       req.Username,
		UsedEmail:      oldAccount.Email,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Gender:         req.Gender,
		Birthdate:      req.Birthdate,
		ProfilePicture: req.ProfilePicture,
	}

	if oldAccount.Email == req.Email {
		_, err := u.accountRepository.UpdateAccountWithoutEmail(ctx, rReq)
		if err != nil {
			return &res, err
		}
	} else {
		_, err = u.accountRepository.UpdateAccount(ctx, rReq)
		if err != nil {
			return &res, err
		}
	}

	res.ID = rReq.UserId
	res.FullName = rReq.FullName
	res.Username = rReq.Username
	res.Email = rReq.Email
	res.PhoneNumber = rReq.PhoneNumber
	res.Gender = rReq.Gender
	res.Birthdate = rReq.Birthdate
	res.ProfilePicture = rReq.ProfilePicture

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

	res.ID = userAccount.ID
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

func (u *accountUsecase) GetWallet(ctx context.Context, req dtousecase.AccountRequest) (*dtousecase.WalletResponse, error) {
	userAccount, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: req.ID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	IsActive := true
	if userAccount.WalletPin == "" {
		IsActive = false
	}

	return &dtousecase.WalletResponse{
		Balance:      userAccount.Balance,
		WalletNumber: userAccount.WalletNumber,
		IsActive:     IsActive,
	}, nil
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

func (u *accountUsecase) TopUpBalanceWallet(ctx context.Context, walletReq dtousecase.TopUpBalanceWalletRequest) (*dtousecase.TopUpBalanceWalletResponse, error) {
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

	if walletReq.Amount.LessThan(constant.TopupAmountMin) || walletReq.Amount.GreaterThan(constant.TopupAmountMax) {
		return nil, util.ErrInvalidAmountRange
	}

	acc, err := u.accountRepository.TopUpWalletBalanceByID(ctx, dtorepository.TopUpWalletRequest{
		UserID: walletReq.UserID,
		Amount: walletReq.Amount,
		Type:   "TOP UP",
	})

	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	return &dtousecase.TopUpBalanceWalletResponse{
		WalletNumber: acc.WalletNumber,
		Balance:      acc.Balance,
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

func (u *accountUsecase) GetCart(ctx context.Context, req dtousecase.GetCartRequest) (dtousecase.GetCartResponse, error) {
	res := dtousecase.GetCartResponse{}

	rReq := dtorepository.GetAccountCartItemsRequest{
		AccountId: req.UserId,
	}

	rRes, err := u.accountRepository.FindAccountCartItems(ctx, rReq)
	if err != nil {
		return res, err
	}

	cartShop, err := u.convertCartItems(ctx, rRes)
	if err != nil {
		return res, err
	}

	res.CartShops = cartShop

	return res, nil
}

func (u *accountUsecase) convertCartItems(ctx context.Context, rRes dtorepository.GetAccountCartItemsResponse) ([]dtousecase.CartShop, error) {
	res := []dtousecase.CartShop{}
	cs := dtousecase.CartShop{}

	for _, data := range rRes.CartItems {
		if cs.ShopId != data.ShopId {
			if cs.ShopId != 0 {
				res = append(res, cs)
			}

			cs.ShopId = data.ShopId
			cs.ShopName = data.ShopName
			cs.CartItems = []dtousecase.CartItem{}
		}

		ci := dtousecase.CartItem{}
		ci.ProductImageUrl = data.ProductUrl
		ci.ProductName = data.ProductName
		ci.ProductQuantity = data.Quantity
		ci.ProductUnitPrice = data.ProductPrice
		ci.ProductTotalPrice = data.ProductPrice.Mul(decimal.NewFromInt(int64(ci.ProductQuantity)))
		cs.CartItems = append(cs.CartItems, ci)
	}

	if cs.ShopId != 0 {
		res = append(res, cs)
	}

	return res, nil
}
