package usecase

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
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
	AddProductToCart(ctx context.Context, req dtousecase.AddProductToCartRequest) (dtousecase.AddProductToCartResponse, error)
	GetAddresses(ctx context.Context, req dtousecase.AddressRequest) (*[]dtousecase.AddressResponse, error)
	RegisterSeller(ctx context.Context, req dtousecase.RegisterSellerRequest) (*dtousecase.RegisterSellerResponse, error)
	UpdateCartQuantity(ctx context.Context, req dtousecase.UpdateCartRequest) (*dtousecase.UpdateCartResponse, error)
	DeleteProductCart(ctx context.Context, req dtousecase.DeleteCartProductRequest) ([]model.AccountCarts, error)
	ValidateWalletPIN(ctx context.Context, req dtousecase.ValidateWAlletPINRequest) (*dtousecase.ValidateWAlletPINResponse, error)
	GetCouriers(ctx context.Context) ([]model.Couriers, error)
	RegisterAccountAddress(ctx context.Context, req dtousecase.RegisterAddressRequest) (dtousecase.RegisterAddressResponse, error)
	GetProvinces(ctx context.Context) (dtousecase.GetProvincesResponse, error)
	GetDistricts(ctx context.Context) (dtousecase.GetDistrictResponse, error)
	GetDistrictsByProvinceId(ctx context.Context, req dtousecase.GetDistrictRequest) (dtousecase.GetDistrictResponse, error)
	RefreshToken(ctx context.Context, req dtousecase.RefreshTokenRequest) (*dtousecase.LoginResponse, error)
	DeleteAddresses(ctx context.Context, req dtousecase.DeleteAddressRequest) error
	UpdateAccountAddress(ctx context.Context, req dtousecase.UpdateAddressRequest) (dtousecase.UpdateAddressResponse, error)
	RequestForgetPassword(ctx context.Context, req dtousecase.ForgetPasswordRequest) (*dtousecase.ForgetPasswordRequest, error)
	RequestForgetChangePassword(ctx context.Context, req dtousecase.ForgetChangePasswordRequest) (*dtousecase.ForgetPasswordRequest, error)
	ChangePassword(ctx context.Context, req dtousecase.ChangePasswordRequest) error
	GetCategories(ctx context.Context) (dtousecase.GetCategoriesResponse, error)
	RequestOTP(ctx context.Context, req dtousecase.ChangePasswordRequest) (*model.Accounts, error)
	UpdatePhotoProfile(ctx context.Context, req dtousecase.UpdatePhoto) (dtousecase.UpdatePhoto, error)
}

type accountUsecase struct {
	accountRepository   repository.AccountRepository
	usedEmailRepository repository.UsedEmailRepository
	courierRepository   repository.CourierRepository
	productRepository   repository.ProductRepository
}

type AccountUsecaseConfig struct {
	AccountRepository   repository.AccountRepository
	UsedEmailRepository repository.UsedEmailRepository
	ProductRepository   repository.ProductRepository
	CourierRepository   repository.CourierRepository
}

func NewAccountUsecase(config AccountUsecaseConfig) AccountUsecase {
	au := &accountUsecase{}
	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
		au.usedEmailRepository = config.UsedEmailRepository
	}
	if config.ProductRepository != nil {
		au.productRepository = config.ProductRepository
	}

	if config.CourierRepository != nil {
		au.courierRepository = config.CourierRepository
	}

	return au
}

func (u *accountUsecase) RegisterSeller(ctx context.Context, req dtousecase.RegisterSellerRequest) (*dtousecase.RegisterSellerResponse, error) {
	res := dtousecase.RegisterSellerResponse{}

	rReq := dtorepository.RegisterSellerRequest{
		UserId:        req.UserId,
		ShopName:      req.ShopName,
		AddressId:     req.AddressId,
		ListCourierId: req.ListCourierId,
	}

	registeredSeller, err := u.accountRepository.CreateSeller(ctx, rReq)
	if errors.Is(err, util.ErrCourierNotAvailable) {
		return nil, util.ErrCourierNotAvailable
	}

	if errors.Is(err, util.ErrAddressNotAvailable) {
		return nil, util.ErrAddressNotAvailable
	}

	if err != nil {
		return nil, err
	}

	res.ShopName = registeredSeller.ShopName

	return &res, nil
}

func (u *accountUsecase) UpdatePhotoProfile(ctx context.Context, req dtousecase.UpdatePhoto) (dtousecase.UpdatePhoto, error) {
	res := dtousecase.UpdatePhoto{}

	currentTime := time.Now().UnixNano()

	file, err := req.ImageHeader.Open()
	if err != nil {
		return res, err
	}

	fileExtension := path.Ext(req.ImageHeader.Filename)
	originalFilename := req.ImageHeader.Filename[:len(req.ImageHeader.Filename)-len(fileExtension)]
	newFilename := fmt.Sprintf("%s_%d", originalFilename, currentTime)

	imageUrl, err := util.UploadToCloudinary(file, newFilename)
	if err != nil {
		return res, err
	}

	_, err = u.accountRepository.UpdatePhotoProfile(ctx, dtorepository.UpdatePhotoProfile{UserID: req.UserID, ImageURL: imageUrl})
	if err != nil {
		return res, err
	}

	res.ImageURL = imageUrl

	return res, nil
}

func (u *accountUsecase) RequestOTP(ctx context.Context, req dtousecase.ChangePasswordRequest) (*model.Accounts, error) {
	res := model.Accounts{}

	token, err := util.GenerateRandomOTP()
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	account, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: req.AccountID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.EmailNotFound
	}
	if err != nil {
		return nil, err
	}
	data := dtousecase.SendEmailPayload{
		RecipientName:  account.FullName,
		RecipientEmail: account.Email,
		Token:          token,
		ExpiresAt:      expirationTime,
	}

	err = util.SendMailOTP(data)
	if err != nil {
		return nil, err
	}

	_, err = u.accountRepository.SaveChangePasswordToken(ctx, dtorepository.RequestChangePasswordRequest{
		UserId:                  account.ID,
		Email:                   account.Email,
		ChangePasswordToken:     token,
		ChangePasswordExpiredAt: expirationTime,
	})
	if err != nil {
		return nil, err
	}

	res.Email = account.Email

	return &res, nil
}

func (u *accountUsecase) ChangePassword(ctx context.Context, req dtousecase.ChangePasswordRequest) error {
	account, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{
		UserId: req.AccountID,
	})
	if err != nil {
		return err
	}

	if !(account.ChangePasswordToken == req.OTP) {
		return util.ErrInvalidOTP
	}

	if account.ChangePasswordExpiredAt.Before(time.Now()) {
		return util.ErrExpiredOTP
	}

	if !util.CheckPasswordHash(req.OldPassword, account.Password) {
		return util.ErrInvalidPassword
	}

	if !util.ValidatePassword(req.NewPassword) {
		return util.ErrWeakPassword
	}

	if len(req.NewPassword) < 8 {
		return util.ErrWeakPassword
	}

	if req.OldPassword == req.NewPassword {
		return util.ErrSamePassword
	}

	if util.CheckPasswordIdentical(account.Username, req.NewPassword) {
		return util.ErrPasswordIdentical
	}

	password, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	rReq := dtorepository.ChangePasswordRequest{
		AccountID:   req.AccountID,
		NewPassword: password,
	}
	_, err = u.accountRepository.ChangePasswordUpdate(ctx, rReq)
	if err != nil {
		return err
	}

	return nil
}

func (u *accountUsecase) GetAddresses(ctx context.Context, req dtousecase.AddressRequest) (*[]dtousecase.AddressResponse, error) {
	res := []dtousecase.AddressResponse{}

	rReq := dtorepository.AddressRequest{
		UserId: req.UserId,
	}
	addresses, err := u.accountRepository.GetAddresses(ctx, rReq)
	if err != nil {
		return nil, err
	}

	for _, data := range *addresses {
		res = append(res, dtousecase.AddressResponse{
			ID:              data.ID,
			FullAddress:     data.FullAddress,
			Detail:          data.Detail,
			ZipCode:         data.ZipCode,
			Kelurahan:       data.Kelurahan,
			SubDistrict:     data.SubDistrict,
			DistrictId:      data.DistrictId,
			District:        data.District,
			ProvinceId:      data.ProvinceId,
			Province:        data.Province,
			IsBuyerDefault:  data.IsBuyerDefault,
			IsSellerDefault: data.IsSellerDefault,
		})
	}

	return &res, nil
}

func (u *accountUsecase) DeleteAddresses(ctx context.Context, req dtousecase.DeleteAddressRequest) error {
	err := u.accountRepository.DeleteAddress(ctx, dtorepository.DeleteAddressRequest{AddressId: req.AddressId})

	return err
}

func (u *accountUsecase) RequestForgetPassword(ctx context.Context, req dtousecase.ForgetPasswordRequest) (*dtousecase.ForgetPasswordRequest, error) {
	res := dtousecase.ForgetPasswordRequest{}

	token, err := util.GenerateRandomToken()
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	account, err := u.accountRepository.FindByEmail(ctx, dtorepository.GetAccountRequest{Email: req.Email})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.EmailNotFound
	}
	if err != nil {
		return nil, err
	}
	data := dtousecase.SendEmailPayload{
		RecipientName:  account.FullName,
		RecipientEmail: req.Email,
		Token:          token,
		ExpiresAt:      expirationTime,
	}

	err = util.SendMail(data)
	if err != nil {
		return nil, err
	}

	_, err = u.accountRepository.SaveForgetPasswordToken(ctx, dtorepository.RequestForgetPasswordRequest{
		UserId:                  account.ID,
		Email:                   req.Email,
		ForgetPasswordToken:     token,
		ForgetPasswordExpiredAt: expirationTime,
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *accountUsecase) RequestForgetChangePassword(ctx context.Context, req dtousecase.ForgetChangePasswordRequest) (*dtousecase.ForgetPasswordRequest, error) {
	res := dtousecase.ForgetPasswordRequest{}

	account, err := u.accountRepository.FindByToken(ctx, dtorepository.RequestForgetPasswordRequest{ForgetPasswordToken: req.Token})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrRequestForgetToken
	}
	if err != nil {
		return nil, err
	}

	if time.Now().After(account.ForgetPasswordExpiredAt) {
		return nil, util.ErrRequestForgetToken
	}

	if util.CheckPasswordHash(req.Password, account.Password) {
		return nil, util.ErrSamePassword
	}

	if strings.Contains(strings.ToLower(req.Password), strings.ToLower(account.Username)) {
		return nil, util.ErrPasswordContainUsername
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = u.accountRepository.UpdatePassword(ctx, dtorepository.RequestForgetPasswordRequest{UserId: account.ID, Password: hashedPassword})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrRequestForgetToken
	}
	if err != nil {
		return nil, err
	}

	return &res, nil
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

	token, err := util.GenerateJWT(userAccount.ID, role, userAccount.WalletNumber)
	if err != nil {
		return nil, err
	}
	refreshToken, err := util.GenerateRefreshJWT(userAccount.ID)
	if err != nil {
		return nil, err
	}

	res.AccessToken = token
	res.RefreshToken = refreshToken

	return &res, nil
}

func (u *accountUsecase) RefreshToken(ctx context.Context, req dtousecase.RefreshTokenRequest) (*dtousecase.LoginResponse, error) {
	token, err := util.ValidateRefreshToken(req.RefreshToken)

	if err != nil || !token.Valid {
		return nil, util.ErrInvalidToken
	}

	claims, ok := token.Claims.(*dtogeneral.ClaimsJWT)

	if !ok {
		return nil, util.ErrUnauthorize
	}
	account, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: claims.UserId})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if err != nil {
		return nil, util.ErrUnauthorize
	}
	role := "buyer"
	if account.ShopName != "" {
		role = "seller"
	}

	accessToken, err := util.GenerateJWT(account.ID, role, account.WalletNumber)
	if err != nil {
		return nil, err
	}
	refreshToken, err := util.GenerateRefreshJWT(account.ID)
	if err != nil {
		return nil, err
	}

	return &dtousecase.LoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil

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

func (u *accountUsecase) ValidateWalletPIN(ctx context.Context, req dtousecase.ValidateWAlletPINRequest) (*dtousecase.ValidateWAlletPINResponse, error) {
	userAccount, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: req.UserID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if userAccount.WalletPin == "" {
		return nil, util.ErrWalletNotSet
	}

	if userAccount.WalletPin != req.WalletPIN {
		return &dtousecase.ValidateWAlletPINResponse{
			IsCorrect: false,
		}, nil
	}

	return &dtousecase.ValidateWAlletPINResponse{
		IsCorrect: true,
	}, nil
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

	isSeller := false

	if userAccount.ShopName != "" {
		isSeller = true
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
	res.IsSeller = isSeller

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
		From:   "5550000012345",
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

	isCorrect := util.CheckPasswordHash(accountReq.Password, userAccount.Password)

	if !isCorrect {
		return nil, util.ErrIncorrectPassword
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
		ci.ProductId = data.ProductId
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

func (u *accountUsecase) UpdateCartQuantity(ctx context.Context, req dtousecase.UpdateCartRequest) (*dtousecase.UpdateCartResponse, error) {
	product, err := u.productRepository.FindProductVariantByID(ctx, dtorepository.ProductCart{ProductID: req.ProductID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if req.Quantity > product.Quantity {
		return nil, util.ErrQtyExceed
	}

	cart, err := u.accountRepository.UpdateCartQuantity(ctx, dtorepository.UpdateCart{ProductID: req.ProductID, Quantity: req.Quantity})
	if err != nil {
		return nil, err
	}

	if cart.ProductID == 0 {
		return nil, util.ErrProductCartNotFound
	}

	return &dtousecase.UpdateCartResponse{
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
	}, nil
}

func (u *accountUsecase) AddProductToCart(ctx context.Context, req dtousecase.AddProductToCartRequest) (dtousecase.AddProductToCartResponse, error) {
	res := dtousecase.AddProductToCartResponse{}

	rReq := dtorepository.AddProductToCartRequest{
		AccountId:                   req.UserId,
		ProductVariantCombinationId: req.ProductVariantId,
		Quantity:                    req.Quantity,
	}

	rRes, err := u.accountRepository.AddProductToCart(ctx, rReq)
	if err != nil {
		return res, err
	}

	res.ProductId = rRes.ProductVariantCombinationId
	res.Quantity = rRes.Quantity

	return res, nil
}

func (u *accountUsecase) GetCouriers(ctx context.Context) ([]model.Couriers, error) {
	response, err := u.courierRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *accountUsecase) DeleteProductCart(ctx context.Context, req dtousecase.DeleteCartProductRequest) ([]model.AccountCarts, error) {

	res, err := u.accountRepository.DeleteCartProduct(ctx, dtorepository.DeleteCartProductRequest{
		ListProductID: req.ListProductID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *accountUsecase) GetProvinces(ctx context.Context) (dtousecase.GetProvincesResponse, error) {
	res := dtousecase.GetProvincesResponse{}

	p, err := u.accountRepository.FindProvinces(ctx)
	if err != nil {
		return res, err
	}

	provinces, err := u.convertProvincesData(ctx, p)
	if err != nil {
		return res, err
	}

	res.Provinces = provinces

	return res, nil
}

func (u *accountUsecase) GetDistricts(ctx context.Context) (dtousecase.GetDistrictResponse, error) {
	res := dtousecase.GetDistrictResponse{}

	d, err := u.accountRepository.FindDistricts(ctx)
	if err != nil {
		return res, err
	}

	districts, err := u.convertDistrictsData(ctx, d)
	if err != nil {
		return res, err
	}

	res.Districts = districts

	return res, nil
}

func (u *accountUsecase) convertProvincesData(ctx context.Context, req []model.Province) ([]dtousecase.Province, error) {
	res := []dtousecase.Province{}

	for _, data := range req {
		p := dtousecase.Province{}

		p.Id = data.ID
		p.Name = data.Name

		res = append(res, p)
	}

	return res, nil
}

func (u *accountUsecase) GetDistrictsByProvinceId(ctx context.Context, req dtousecase.GetDistrictRequest) (dtousecase.GetDistrictResponse, error) {
	res := dtousecase.GetDistrictResponse{}

	d, err := u.accountRepository.FindDistrictsByProvinceId(ctx, req.ProvinceId)
	if err != nil {
		return res, err
	}

	districts, err := u.convertDistrictsData(ctx, d)
	if err != nil {
		return res, err
	}

	res.Districts = districts

	return res, nil
}

func (u *accountUsecase) convertDistrictsData(ctx context.Context, req []model.District) ([]dtousecase.District, error) {
	res := []dtousecase.District{}

	for _, data := range req {
		p := dtousecase.District{}

		p.Id = data.ID
		p.Name = data.Name

		res = append(res, p)
	}

	return res, nil
}

func (u *accountUsecase) RegisterAccountAddress(ctx context.Context, req dtousecase.RegisterAddressRequest) (dtousecase.RegisterAddressResponse, error) {
	res := dtousecase.RegisterAddressResponse{}

	rReq := dtorepository.RegisterAddressRequest{
		AccountId:   req.AccountId,
		ProvinceId:  req.ProvinceId,
		DistrictId:  req.DistrictId,
		SubDistrict: req.SubDistrict,
		Kelurahan:   req.Kelurahan,
		ZipCode:     req.ZipCode,
		Detail:      req.Detail,
	}
	rRes, err := u.accountRepository.CreateAddress(ctx, rReq)
	if err != nil {
		return res, err
	}

	res.Detail = rRes.Detail
	res.DistrictId = rRes.DistrictId
	res.Kelurahan = rRes.Kelurahan
	res.ProvinceId = rRes.ProvinceId
	res.SubDistrict = rRes.SubDistrict
	res.AccountId = rRes.AccountId
	res.ZipCode = rRes.ZipCode

	return res, nil
}

func (u *accountUsecase) UpdateAccountAddress(ctx context.Context, req dtousecase.UpdateAddressRequest) (dtousecase.UpdateAddressResponse, error) {
	res := dtousecase.UpdateAddressResponse{}

	_, err := u.accountRepository.FindAddressByID(ctx, dtorepository.UpdateAddressRequest{AddressId: req.AddressId, AccountId: req.AccountId})
	if errors.Is(err, util.ErrNoRecordFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	rReq := dtorepository.UpdateAddressRequest{
		AddressId:       req.AddressId,
		AccountId:       req.AccountId,
		ProvinceId:      req.ProvinceId,
		DistrictId:      req.DistrictId,
		SubDistrict:     req.SubDistrict,
		Kelurahan:       req.Kelurahan,
		ZipCode:         req.ZipCode,
		Detail:          req.Detail,
		IsBuyerDefault:  req.IsBuyerDefault,
		IsSellerDefault: req.IsSellerDefault,
	}
	rRes, err := u.accountRepository.UpdateAddress(ctx, rReq)
	if err != nil {
		return res, err
	}

	res.Detail = rRes.Detail
	res.DistrictId = rRes.DistrictId
	res.Kelurahan = rRes.Kelurahan
	res.ProvinceId = rRes.ProvinceId
	res.SubDistrict = rRes.SubDistrict
	res.AccountId = rRes.AccountId
	res.ZipCode = rRes.ZipCode
	res.IsBuyerDefault = rReq.IsBuyerDefault
	res.IsSellerDefault = rReq.IsSellerDefault

	return res, nil
}

func (u *accountUsecase) GetCategories(ctx context.Context) (dtousecase.GetCategoriesResponse, error) {
	res := dtousecase.GetCategoriesResponse{}

	categories, err := u.accountRepository.FindCategories(ctx)
	if err != nil {
		return res, err
	}

	lastLevel1 := 0
	lastLevel2 := 0
	c := dtousecase.Category{}
	c2 := dtousecase.Category{}
	for i, category := range categories {
		if category.CategoryLevel1Id != lastLevel1 {
			if c.Id != 0 {
				c.Children = append(c.Children, c2)
				res.Categories = append(res.Categories, c)
			}

			c.Id = category.CategoryLevel1Id
			c.Name = category.CategoryLevel1Name
			c.Children = []dtousecase.Category{}
			lastLevel1 = category.CategoryLevel1Id

			c2.Id = category.CategoryLevel2Id
			c2.Name = category.CategoryLevel2Name
			c2.Children = []dtousecase.Category{}
			lastLevel2 = category.CategoryLevel2Id
		}

		if category.CategoryLevel2Id != lastLevel2 {
			if c2.Id != 0 {
				c.Children = append(c.Children, c2)
			}

			c2.Id = category.CategoryLevel2Id
			c2.Name = category.CategoryLevel2Name
			c2.Children = []dtousecase.Category{}
			lastLevel2 = category.CategoryLevel2Id
		}

		if category.CategoryLevel3Id != 0 {
			c2.Children = append(c2.Children, dtousecase.Category{
				Id:   category.CategoryLevel3Id,
				Name: category.CategoryLevel3Name,
			})
		}

		if i == len(categories)-1 {
			c.Children = append(c.Children, c2)
		}
	}

	return res, nil
}
