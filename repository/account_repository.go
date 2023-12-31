package repository

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository interface {
	ActivateWalletByID(ctx context.Context, userId int, walletPin string) (model.Accounts, error)
	FindById(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error)
	Create(ctx context.Context, req dtorepository.CreateAccountRequest) (dtorepository.CreateAccountResponse, error)
	UpdateWalletPINByID(ctx context.Context, req dtorepository.UpdateWalletPINRequest) (dtorepository.UpdateWalletPINResponse, error)
	UpdateAccount(ctx context.Context, req dtorepository.EditAccountRequest) (*dtorepository.EditAccountResponse, error)
	FindByEmail(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error)
	TopUpWalletBalanceByID(ctx context.Context, req dtorepository.TopUpWalletRequest) (dtorepository.WalletResponse, error)
	FindByUsername(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error)
	FindByPhoneNumber(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error)
	UpdateAccountWithoutEmail(ctx context.Context, req dtorepository.EditAccountRequest) (*dtorepository.EditAccountResponse, error)
	RefundBalance(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error)
	DecreaseBalanceSellerWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error)
	DecreaseBalanceBuyerWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error)
	IncreaseBalanceSallerWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error)
	FindAccountCartItems(ctx context.Context, req dtorepository.GetAccountCartItemsRequest) (dtorepository.GetAccountCartItemsResponse, error)
	AddProductToCart(ctx context.Context, req dtorepository.AddProductToCartRequest) (dtorepository.AddProductToCartResponse, error)
	GetAddresses(ctx context.Context, req dtorepository.AddressRequest) (*[]dtorepository.AddressResponse, error)
	CreateSeller(ctx context.Context, req dtorepository.RegisterSellerRequest) (*dtorepository.RegisterSellerResponse, error)
	UpdateCartQuantity(ctx context.Context, req dtorepository.UpdateCart) (dtorepository.UpdateCart, error)
	AddressAndCourierIsAvailable(ctx context.Context, tx *gorm.DB, req dtorepository.RegisterSellerRequest) error
	AlreadyRegisteredAsSeller(ctx context.Context, tx *gorm.DB, req dtorepository.RegisterSellerRequest) error
	UpdateShopNameAndSellerDefaultAddress(ctx context.Context, tx *gorm.DB, req dtorepository.RegisterSellerRequest) error
	ConvertListCourierIdToListCourierModel(ctx context.Context, req dtorepository.RegisterSellerRequest) []model.SellerCouriers
	DeleteCartProduct(ctx context.Context, req dtorepository.DeleteCartProductRequest) ([]model.AccountCarts, error)
	CreateAddress(ctx context.Context, req dtorepository.RegisterAddressRequest) (dtorepository.RegisterAddressResponse, error)
	FindProvinces(ctx context.Context) ([]model.Province, error)
	FindDistricts(ctx context.Context) ([]model.District, error)
	FindDistrictsByProvinceId(ctx context.Context, ProvinceId int) ([]model.District, error)
	DeleteAddress(ctx context.Context, req dtorepository.DeleteAddressRequest) error
	UpdateAddress(ctx context.Context, req dtorepository.UpdateAddressRequest) (dtorepository.UpdateAddressResponse, error)
	FindAddressByID(ctx context.Context, req dtorepository.UpdateAddressRequest) (dtorepository.UpdateAddressResponse, error)
	FirstSeller(ctx context.Context, req dtorepository.SellerDataRequest) (dtorepository.SellerDataResponse, error)
	FindSellerBestSelling(ctx context.Context, req dtorepository.FindSellerBestSellingRequest) (dtorepository.FindSellerBestSellingResponse, error)
	FindSellerShowcases(ctx context.Context, req dtorepository.FindSellerShowcasesRequest) ([]dtorepository.FindSellerShowcasesResponse, error)
	FindSellerShowcaseProduct(ctx context.Context, req dtorepository.FindSellerShowcaseProductRequest) (dtousecase.GetSellerShowcaseProductResponse, error)
	FindByToken(ctx context.Context, req dtorepository.RequestForgetPasswordRequest) (dtorepository.GetAccountResponse, error)
	SaveForgetPasswordToken(ctx context.Context, req dtorepository.RequestForgetPasswordRequest) (dtorepository.GetAccountResponse, error)
	UpdatePassword(ctx context.Context, req dtorepository.RequestForgetPasswordRequest) (dtorepository.GetAccountResponse, error)
	ChangePasswordUpdate(ctx context.Context, req dtorepository.ChangePasswordRequest) (dtorepository.GetAccountResponse, error)
	FindCategories(ctx context.Context) ([]dtorepository.Category, error)
	FindCategoryByID(ctx context.Context, categoryID int) (dtorepository.Category, error)
	SaveChangePasswordToken(ctx context.Context, req dtorepository.RequestChangePasswordRequest) (dtorepository.GetAccountResponse, error)
	UpdateShopProfile(ctx context.Context, req dtousecase.UpdateShopProfileRequest) (dtousecase.UpdateShopProfileResponse, error)
	UpdatePhotoProfile(ctx context.Context, req dtorepository.UpdatePhotoProfile) (dtorepository.UpdatePhotoProfile, error)
}

type accountRepository struct {
	db                         *gorm.DB
	usedEmailRepo              usedEmailRepository
	walletTransactionHistories walletTransactionHistoryRepository
	saleTransactionHistories   saleWalletTransactionHistoryRepository
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) AddressAndCourierIsAvailable(ctx context.Context, tx *gorm.DB, req dtorepository.RegisterSellerRequest) error {
	couriers := []model.Couriers{}
	err := tx.WithContext(ctx).Where("id = ?", req.AddressId).First(&model.AccountAddress{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return util.ErrAddressNotAvailable
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	err = r.db.WithContext(ctx).Model(&model.Couriers{}).Where("id IN ?", req.ListCourierId).Scan(&couriers).Error
	if len(couriers) < len(req.ListCourierId) {
		tx.Rollback()
		return util.ErrCourierNotAvailable
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *accountRepository) AlreadyRegisteredAsSeller(ctx context.Context, tx *gorm.DB, req dtorepository.RegisterSellerRequest) error {
	account := model.Accounts{}
	err := tx.WithContext(ctx).Where("id = ?", req.UserId).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return util.ErrNoRecordFound
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	if account.ShopName != "" {
		tx.Rollback()
		return util.ErrAlreadyRegisteredAsSeller
	}

	return nil
}

func (r *accountRepository) UpdateShopNameAndSellerDefaultAddress(ctx context.Context, tx *gorm.DB, req dtorepository.RegisterSellerRequest) error {
	account := model.Accounts{}
	accountAddress := model.AccountAddress{}

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserId).Update("shop_name", req.ShopName).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(accountAddress).Where("id = ?", req.AddressId).Updates(
		model.AccountAddress{
			IsBuyerDefault:  false,
			IsSellerDefault: true,
		}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r accountRepository) ConvertListCourierIdToListCourierModel(ctx context.Context, req dtorepository.RegisterSellerRequest) []model.SellerCouriers {
	seller_couriers := []model.SellerCouriers{}
	for _, seller_courier_id := range req.ListCourierId {
		seller_couriers = append(seller_couriers, model.SellerCouriers{
			AccountID: req.UserId,
			CourierID: seller_courier_id,
		})
	}

	return seller_couriers
}

func (r *accountRepository) CreateSeller(ctx context.Context, req dtorepository.RegisterSellerRequest) (*dtorepository.RegisterSellerResponse, error) {
	res := dtorepository.RegisterSellerResponse{}
	seller_couriers := r.ConvertListCourierIdToListCourierModel(ctx, req)

	tx := r.db.Begin()
	err := r.AddressAndCourierIsAvailable(ctx, tx, req)
	if err != nil {
		tx.Rollback()
		return &res, err
	}

	err = r.AlreadyRegisteredAsSeller(ctx, tx, req)
	if err != nil {
		tx.Rollback()
		return &res, err
	}

	err = r.UpdateShopNameAndSellerDefaultAddress(ctx, tx, req)
	if err != nil {
		tx.Rollback()
		return &res, err
	}

	err = tx.WithContext(ctx).Create(&seller_couriers).Error
	if err != nil {
		tx.Rollback()
		return &res, err
	}

	tx.Commit()

	res.ShopName = req.ShopName

	return &res, nil
}

func (r *accountRepository) ChangePasswordUpdate(ctx context.Context, req dtorepository.ChangePasswordRequest) (dtorepository.GetAccountResponse, error) {
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Model(&model.Accounts{}).
		Where("id = ?", req.AccountID).
		Update("password", req.NewPassword).
		Update("change_password_token", gorm.Expr("NULL")).
		Update("change_password_expired_at", gorm.Expr("NULL")).
		Scan(&res).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) GetAddresses(ctx context.Context, req dtorepository.AddressRequest) (*[]dtorepository.AddressResponse, error) {
	res := []dtorepository.AddressResponse{}

	q := `
		select
			aa.id as "ID",
			concat(aa.detail, aa.kelurahan, aa.sub_district, aa.district, aa.province) as "FullAddress",
			aa.detail as "Detail",          
			aa.zip_code as "ZipCode",         
			aa.kelurahan as "Kelurahan",       
			aa.sub_district as "SubDistrict",     
			d.id as "DistrictId",
			aa.district as "District",        
			p.id as "ProvinceId",
			aa.province as "Province",        
			aa.is_buyer_default as "IsBuyerDefault",  
			aa.is_seller_default as "IsSellerDefault" 
		from account_addresses aa 
		left join provinces p 
			on p."name" = aa.province 
		left join districts d 
			on d."name" = aa.district
		where aa.account_id = ?
			and aa.deleted_at is null
	`

	err := r.db.WithContext(ctx).Raw(q, req.UserId).Scan(&res).Error
	if err != nil {
		return &res, util.ErrNoRecordFound
	}

	return &res, nil
}

func (r *accountRepository) DeleteAddress(ctx context.Context, req dtorepository.DeleteAddressRequest) error {
	ad := model.AccountAddress{}

	err := r.db.WithContext(ctx).Where("id = ?", req.AddressId).First(&ad).Error
	if err != nil {
		return err
	}

	ad.DeletedAt = time.Now()

	err = r.db.WithContext(ctx).Updates(&ad).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) UpdateAccount(ctx context.Context, req dtorepository.EditAccountRequest) (*dtorepository.EditAccountResponse, error) {
	res := &dtorepository.EditAccountResponse{}

	tx := r.db.Begin()

	a := model.Accounts{}
	err := tx.WithContext(ctx).Where("id = ?", req.UserId).First(&a).Error

	if err != nil {
		tx.Rollback()
		return res, err
	}

	a.FullName = req.FullName
	a.Username = req.Username
	a.Email = req.Email
	a.PhoneNumber = req.PhoneNumber
	a.Gender = req.Gender
	a.Birthdate = req.Birthdate

	err = tx.WithContext(ctx).Updates(&a).Error

	if err != nil {
		tx.Rollback()
		return res, err
	}

	ueReq := dtorepository.UsedEmailRequest{
		AccountID: req.UserId,
		Email:     req.UsedEmail,
	}

	_, err = r.usedEmailRepo.CreateEmail(ctx, tx, ueReq)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	res.ID = a.ID
	res.FullName = a.FullName
	res.Username = a.Username
	res.Email = a.Email
	res.PhoneNumber = a.PhoneNumber
	res.Gender = a.Gender
	res.Birthdate = a.Birthdate

	return res, nil
}

func (r *accountRepository) UpdateAccountWithoutEmail(ctx context.Context, req dtorepository.EditAccountRequest) (*dtorepository.EditAccountResponse, error) {
	res := &dtorepository.EditAccountResponse{}

	a := model.Accounts{}
	err := r.db.WithContext(ctx).Where("id = ?", req.UserId).First(&a).Error

	if err != nil {
		return res, err
	}

	a.FullName = req.FullName
	a.Username = req.Username
	a.Email = req.Email
	a.PhoneNumber = req.PhoneNumber
	a.Gender = req.Gender
	a.Birthdate = req.Birthdate

	err = r.db.WithContext(ctx).Updates(&a).Error

	if err != nil {
		return res, err
	}

	res.ID = a.ID
	res.FullName = a.FullName
	res.Username = a.Username
	res.Email = a.Email
	res.PhoneNumber = a.PhoneNumber
	res.Gender = a.Gender
	res.Birthdate = a.Birthdate
	res.ProfilePicture = a.ProfilePicture

	return res, nil
}

func (r *accountRepository) ActivateWalletByID(ctx context.Context, userId int, walletPin string) (model.Accounts, error) {
	account := model.Accounts{}

	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", userId).Update("wallet_pin", walletPin).Error

	if err != nil {
		return model.Accounts{}, err
	}

	return account, nil
}

func (r *accountRepository) UpdateCartQuantity(ctx context.Context, req dtorepository.UpdateCart) (dtorepository.UpdateCart, error) {
	account := model.AccountCarts{}

	err := r.db.WithContext(ctx).Model(&account).Where("product_variant_selection_combination_id = ?", req.ProductID).Update("quantity", req.Quantity).Scan(&account).Error

	if err != nil {
		return dtorepository.UpdateCart{}, err
	}

	return dtorepository.UpdateCart{
		ProductID: account.ProductVariantSelectionCombinationId,
		Quantity:  account.Quantity,
	}, nil
}

func (r *accountRepository) DeleteCartProduct(ctx context.Context, req dtorepository.DeleteCartProductRequest) ([]model.AccountCarts, error) {
	account := []model.AccountCarts{}

	err := r.db.WithContext(ctx).Where("product_variant_selection_combination_id IN ?", req.ListProductID).Delete(&account).Scan(&account).Error

	if err != nil {
		return account, err
	}

	return account, nil
}

func (r *accountRepository) UpdatePhotoProfile(ctx context.Context, req dtorepository.UpdatePhotoProfile) (dtorepository.UpdatePhotoProfile, error) {
	account := model.Accounts{}

	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("profile_picture", req.ImageURL).Error

	if err != nil {
		return dtorepository.UpdatePhotoProfile{}, err
	}

	return dtorepository.UpdatePhotoProfile{}, nil
}

func (r *accountRepository) UpdateWalletPINByID(ctx context.Context, req dtorepository.UpdateWalletPINRequest) (dtorepository.UpdateWalletPINResponse, error) {
	account := model.Accounts{}

	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("wallet_pin", req.WalletNewPIN).Error

	if err != nil {
		return dtorepository.UpdateWalletPINResponse{}, err
	}

	return dtorepository.UpdateWalletPINResponse{
		UserID:       account.ID,
		WalletNewPIN: account.WalletPin,
	}, nil
}

func (r *accountRepository) RefundBalance(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error) {
	account := model.Accounts{}

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("balance", gorm.Expr("balance + ?", req.Balance)).Scan(&account).Error

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	_, err = r.walletTransactionHistories.CreateWithTx(ctx, tx, model.MyWalletTransactionHistories{
		AccountID:      req.UserID,
		Amount:         req.Balance,
		Type:           "Refund",
		From:           req.WalletNumber,
		ProductOrderID: req.ProductOrderID,
	})

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	return dtorepository.WalletResponse{
		UserID:       account.ID,
		WalletNumber: account.WalletNumber,
		Balance:      account.Balance,
	}, nil
}

func (r *accountRepository) DecreaseBalanceSellerWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error) {
	account := model.Accounts{}

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("seller_balance", gorm.Expr("seller_balance - ?", req.Balance)).Error

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	_, err = r.saleTransactionHistories.CreateWithTx(ctx, tx, model.SaleWalletTransactionHistories{
		AccountID:      req.UserID,
		Amount:         req.Balance.Neg(),
		Type:           "Refund",
		To:             req.WalletNumber,
		ProductOrderID: req.ProductOrderID,
	})

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	return dtorepository.WalletResponse{
		UserID:       account.ID,
		WalletNumber: account.WalletNumber,
		Balance:      account.Balance,
	}, nil
}

func (r *accountRepository) DecreaseBalanceBuyerWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error) {
	account := model.Accounts{}

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("balance", gorm.Expr("balance - ?", req.Balance)).Error

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	_, err = r.walletTransactionHistories.CreateWithTx(ctx, tx, model.MyWalletTransactionHistories{
		AccountID:      req.UserID,
		Amount:         req.Balance.Neg(),
		Type:           req.TransactionType,
		ProductOrderID: req.ProductOrderID,
		To:             req.WalletNumber,
	})

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	return dtorepository.WalletResponse{
		UserID:       account.ID,
		WalletNumber: account.WalletNumber,
		Balance:      account.Balance,
	}, nil
}

func (r *accountRepository) IncreaseBalanceSallerWithTx(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error) {
	account := model.Accounts{}

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("seller_balance", gorm.Expr("seller_balance + ?", req.Balance)).Error

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	_, err = r.saleTransactionHistories.CreateWithTx(ctx, tx, model.SaleWalletTransactionHistories{
		AccountID:      req.UserID,
		Amount:         req.Balance,
		Type:           req.TransactionType,
		ProductOrderID: req.ProductOrderID,
		From:           req.WalletNumber,
	})

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	return dtorepository.WalletResponse{
		UserID:       account.ID,
		WalletNumber: account.WalletNumber,
		Balance:      account.Balance,
	}, nil
}

func (r *accountRepository) TopUpWalletBalanceByID(ctx context.Context, req dtorepository.TopUpWalletRequest) (dtorepository.WalletResponse, error) {
	account := model.Accounts{}
	tx := r.db.Begin()

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&account).Where("id = ?", req.UserID).Update("balance", gorm.Expr("balance + ?", req.Amount)).Error

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	_, err = r.walletTransactionHistories.CreateWithTx(ctx, tx, model.MyWalletTransactionHistories{
		AccountID: req.UserID,
		Amount:    req.Amount,
		Type:      req.Type,
		From:      req.From,
	})

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	if req.Type == constant.Withdraw {
		err := tx.WithContext(ctx).Model(&model.SaleWalletTransactionHistories{}).
			Where("product_order_id = ?", req.OrderID).
			Update("is_withdrawn", true).Error

		if err != nil {
			tx.Rollback()
			return dtorepository.WalletResponse{}, err
		}
	}

	tx.Commit()

	return dtorepository.WalletResponse{
		UserID:       account.ID,
		WalletNumber: account.WalletNumber,
		Balance:      account.Balance,
	}, nil
}

func (r *accountRepository) FindById(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error) {
	account := model.Accounts{}
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Where("id = ?", req.UserId).First(&account).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	if err != nil {
		return res, err
	}

	res.FullName = account.FullName
	res.Username = account.Username
	res.Email = account.Email
	res.PhoneNumber = account.PhoneNumber
	res.Gender = account.Gender
	res.Birthdate = account.Birthdate
	res.ProfilePicture = account.ProfilePicture
	res.WalletNumber = account.WalletNumber
	res.Balance = account.Balance
	res.Password = account.Password
	res.WalletPin = account.WalletPin
	res.ID = account.ID
	res.ForgetPasswordExpiredAt = account.ForgetPasswordExpiredAt
	res.ForgetPasswordToken = account.ForgetPasswordToken
	res.ShopName = account.ShopName
	res.ChangePasswordToken = account.ChangePasswordToken
	res.ChangePasswordExpiredAt = account.ChangePasswordExpiredAt

	return res, err
}

func (r *accountRepository) FindByPhoneNumber(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error) {
	account := model.Accounts{}
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Where("phone_number = ?", req.PhoneNumber).Where("id != ?", req.UserId).First(&account).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.FullName = account.FullName
	res.Username = account.Username
	res.Email = account.Email
	res.PhoneNumber = account.PhoneNumber
	res.Gender = account.Gender
	res.Birthdate = account.Birthdate
	res.ProfilePicture = account.ProfilePicture
	res.WalletNumber = account.WalletNumber
	res.Balance = account.Balance
	res.Password = account.Password
	res.WalletPin = account.WalletPin
	res.ID = account.ID
	res.ForgetPasswordExpiredAt = account.ForgetPasswordExpiredAt
	res.ForgetPasswordToken = account.ForgetPasswordToken

	return res, err
}

func (r *accountRepository) FindByEmail(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error) {
	account := model.Accounts{}
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Where("email = ?", req.Email).First(&account).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.FullName = account.FullName
	res.Username = account.Username
	res.Email = account.Email
	res.PhoneNumber = account.PhoneNumber
	res.Gender = account.Gender
	res.Birthdate = account.Birthdate
	res.ProfilePicture = account.ProfilePicture
	res.WalletNumber = account.WalletNumber
	res.Balance = account.Balance
	res.Password = account.Password
	res.WalletPin = account.WalletPin
	res.ID = account.ID
	res.ShopName = account.ShopName
	res.ForgetPasswordExpiredAt = account.ForgetPasswordExpiredAt
	res.ForgetPasswordToken = account.ForgetPasswordToken

	return res, err
}

func (r *accountRepository) FindByUsername(ctx context.Context, req dtorepository.GetAccountRequest) (dtorepository.GetAccountResponse, error) {
	account := model.Accounts{}
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Where("username = ?", req.Username).Where("id != ?", req.UserId).First(&account).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.FullName = account.FullName
	res.Username = account.Username
	res.Email = account.Email
	res.PhoneNumber = account.PhoneNumber
	res.Gender = account.Gender
	res.Birthdate = account.Birthdate
	res.ProfilePicture = account.ProfilePicture
	res.WalletNumber = account.WalletNumber
	res.Balance = account.Balance
	res.Password = account.Password
	res.WalletPin = account.WalletPin
	res.ID = account.ID
	res.ForgetPasswordExpiredAt = account.ForgetPasswordExpiredAt
	res.ForgetPasswordToken = account.ForgetPasswordToken

	return res, err
}

func (r *accountRepository) Create(ctx context.Context, req dtorepository.CreateAccountRequest) (dtorepository.CreateAccountResponse, error) {
	res := dtorepository.CreateAccountResponse{}

	a := model.Accounts{
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := r.db.WithContext(ctx).Create(&a).Error

	res.Email = a.Email
	res.FullName = a.FullName
	res.Username = a.Username

	return res, err
}

func (r *accountRepository) FindAccountCartItems(ctx context.Context, req dtorepository.GetAccountCartItemsRequest) (dtorepository.GetAccountCartItemsResponse, error) {
	res := dtorepository.GetAccountCartItemsResponse{}

	q := `
		select 
			seller.id as "ShopId",
			seller.shop_name as "ShopName",
			pvsc.id as "ProductId",
			case 
				when pvs.name = 'default_reserved_keyword' then product_image.url
				else pvsc.picture_url
			end as "ProductUrl",
			case 
				when pvs."name" = 'default_reserved_keyword' then p."name"
				when pvs2."name" is null then concat(p."name", ' - ', pvs."name")
				when pvs2."name" is not null then concat(p."name", ' - ', pvs."name", ', ', pvs2."name")
			end as "ProductName",
			pvsc.price as "ProductPrice",
			ac.quantity as "Quantity"
		from account_carts ac 
			left join product_variant_selection_combinations pvsc 
				on pvsc.id = ac.product_variant_selection_combination_id 
			left join product_variant_selections pvs
				on pvs.id = pvsc.product_variant_selection_id1 
			left join product_variant_selections pvs2 
				on pvs2.id = pvsc.product_variant_selection_id2 
			left join products p 
				on p.id = pvsc.product_id 
			left join lateral (
					select
						pi2.product_id,
						pi2.url 
					from product_images pi2 
					where pi2.product_id = p.id
					limit 1
				) product_image on product_image.product_id = p.id 
			left join accounts seller
				on seller.id = p.seller_id 
		where ac.account_id = ?
		order by seller.id asc
	`

	err := r.db.WithContext(ctx).Raw(q, req.AccountId).Scan(&res.CartItems).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) AddProductToCart(ctx context.Context, req dtorepository.AddProductToCartRequest) (dtorepository.AddProductToCartResponse, error) {
	res := dtorepository.AddProductToCartResponse{}

	pvc := model.ProductVariantSelectionCombinations{}

	err := r.db.WithContext(ctx).Where("id = ?", req.ProductVariantCombinationId).First(&pvc).Error
	if err != nil {
		return res, err
	}

	if pvc.ID == 0 {
		return res, errors.New("product not found")
	}

	product := model.Products{}

	err = r.db.WithContext(ctx).Model(&model.Products{}).Where("id", pvc.ProductID).First(&product).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if product.SellerID == req.AccountId {
		return res, util.ErrSameShop
	}

	c := model.AccountCarts{}

	err = r.db.WithContext(ctx).Where("account_id = $1 and product_variant_selection_combination_id = $2", req.AccountId, req.ProductVariantCombinationId).First(&c).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if c.ID == 0 {
		c1 := model.AccountCarts{
			AccountID:                            req.AccountId,
			ProductVariantSelectionCombinationId: req.ProductVariantCombinationId,
			Quantity:                             req.Quantity,
		}

		err = r.db.WithContext(ctx).Create(&c1).Error
		if err != nil {
			return res, err
		}

		c = c1
	} else {
		c.Quantity += req.Quantity

		err = r.db.WithContext(ctx).Updates(&c).Error
		if err != nil {
			return res, err
		}
	}

	res.AccountId = req.AccountId
	res.Quantity = c.Quantity
	res.ProductVariantCombinationId = c.ProductVariantSelectionCombinationId

	return res, nil
}

func (r *accountRepository) CreateAddress(ctx context.Context, req dtorepository.RegisterAddressRequest) (dtorepository.RegisterAddressResponse, error) {
	res := dtorepository.RegisterAddressResponse{}

	p := model.Province{}
	if err := r.db.WithContext(ctx).Where("id = ?", req.ProvinceId).Find(&p).Error; err != nil {
		return res, err
	}

	d := model.District{}
	if err := r.db.WithContext(ctx).Where("id = ?", req.DistrictId).Find(&d).Error; err != nil {
		return res, err
	}

	ad := model.AccountAddress{
		AccountID:            req.AccountId,
		Province:             p.Name,
		District:             d.Name,
		RajaOngkirDistrictId: d.RajaOngkirDistrictId,
		SubDistrict:          req.SubDistrict,
		Kelurahan:            req.Kelurahan,
		ZipCode:              req.ZipCode,
		Detail:               req.Detail,
	}

	ads := []model.AccountAddress{}
	if err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountId).Find(&ads).Error; err != nil {
		return res, err
	}
	if len(ads) == 0 {
		ad.IsBuyerDefault = true
	}

	if err := r.db.WithContext(ctx).Create(&ad).Error; err != nil {
		return res, err
	}

	res.AccountId = req.AccountId
	res.ProvinceId = req.ProvinceId
	res.DistrictId = req.DistrictId
	res.Kelurahan = req.Kelurahan
	res.SubDistrict = req.SubDistrict
	res.ZipCode = req.ZipCode
	res.Detail = req.Detail

	return res, nil
}

func (r *accountRepository) SaveForgetPasswordToken(ctx context.Context, req dtorepository.RequestForgetPasswordRequest) (dtorepository.GetAccountResponse, error) {
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Model(&model.Accounts{}).Where("id = ?", req.UserId).Update("forget_password_token", req.ForgetPasswordToken).Update("forget_password_expired_at", req.ForgetPasswordExpiredAt).Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) SaveChangePasswordToken(ctx context.Context, req dtorepository.RequestChangePasswordRequest) (dtorepository.GetAccountResponse, error) {
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Model(&model.Accounts{}).
		Where("id = ?", req.UserId).
		Update("change_password_token", req.ChangePasswordToken).
		Update("change_password_expired_at", req.ChangePasswordExpiredAt).
		Scan(&res).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) UpdatePassword(ctx context.Context, req dtorepository.RequestForgetPasswordRequest) (dtorepository.GetAccountResponse, error) {
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Model(&model.Accounts{}).
		Where("id = ?", req.UserId).
		Update("password", req.Password).
		Update("forget_password_token", gorm.Expr("NULL")).
		Update("forget_password_expired_at", gorm.Expr("NULL")).
		Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) FindByToken(ctx context.Context, req dtorepository.RequestForgetPasswordRequest) (dtorepository.GetAccountResponse, error) {
	res := dtorepository.GetAccountResponse{}

	err := r.db.WithContext(ctx).Model(&model.Accounts{}).Where("forget_password_token = ?", req.ForgetPasswordToken).Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

type ChangeDefaultAddress struct {
	ID              int
	AccountID       int
	IsBuyerDefault  bool
	IsSellerDefault bool
}

func (r *accountRepository) UpdateAddress(ctx context.Context, req dtorepository.UpdateAddressRequest) (dtorepository.UpdateAddressResponse, error) {
	res := dtorepository.UpdateAddressResponse{}

	p := model.Province{}
	if err := r.db.WithContext(ctx).Where("id = ?", req.ProvinceId).Find(&p).Error; err != nil {
		return res, err
	}

	d := model.District{}
	if err := r.db.WithContext(ctx).Where("id = ?", req.DistrictId).Find(&d).Error; err != nil {
		return res, err
	}

	ad := model.AccountAddress{
		Province:             p.Name,
		District:             d.Name,
		RajaOngkirDistrictId: d.RajaOngkirDistrictId,
		SubDistrict:          req.SubDistrict,
		Kelurahan:            req.Kelurahan,
		ZipCode:              req.ZipCode,
		Detail:               req.Detail,
		IsBuyerDefault:       req.IsBuyerDefault,
		IsSellerDefault:      req.IsSellerDefault,
	}

	ads := []model.AccountAddress{}
	if err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountId).Find(&ads).Error; err != nil {
		return res, err
	}
	if len(ads) == 0 {
		ad.IsBuyerDefault = true
	}

	tx := r.db.Begin()

	if err := tx.WithContext(ctx).Where("account_id = ?", req.AccountId).Where("id = ?", req.AddressId).Updates(&ad).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	if req.IsBuyerDefault {
		_, err := r.updateBuyerDefaultSetToFalse(ctx, tx, req.AccountId, req.AddressId)
		if err != nil {
			tx.Rollback()
			return res, err
		}
	}

	if req.IsSellerDefault {
		_, err := r.updateSellerDefaultSetToFalse(ctx, tx, req.AccountId, req.AddressId)
		if err != nil {
			tx.Rollback()
			return res, err
		}
	}

	tx.Commit()

	res.AccountId = req.AccountId
	res.ProvinceId = req.ProvinceId
	res.DistrictId = req.DistrictId
	res.Kelurahan = req.Kelurahan
	res.SubDistrict = req.SubDistrict
	res.ZipCode = req.ZipCode
	res.Detail = req.Detail

	return res, nil
}

func (r *accountRepository) updateSellerDefaultSetToFalse(ctx context.Context, tx *gorm.DB, accountId, addressId int) (dtorepository.UpdateAddressResponse, error) {
	res := dtorepository.UpdateAddressResponse{}

	if err := tx.WithContext(ctx).Model(&model.AccountAddress{}).Where("account_id = ?", accountId).Where("id != ?", addressId).Update("is_seller_default", false).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) updateBuyerDefaultSetToFalse(ctx context.Context, tx *gorm.DB, accountId, addressId int) (dtorepository.UpdateAddressResponse, error) {
	res := dtorepository.UpdateAddressResponse{}

	if err := tx.WithContext(ctx).Model(&model.AccountAddress{}).Where("account_id = ?", accountId).Where("id != ?", addressId).Update("is_buyer_default", false).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) FindAddressByID(ctx context.Context, req dtorepository.UpdateAddressRequest) (dtorepository.UpdateAddressResponse, error) {
	res := dtorepository.UpdateAddressResponse{}
	ads := model.AccountAddress{}
	err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountId).Where("id = ?", req.AddressId).First(&ads).Scan(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil

}

func (r *accountRepository) FindProvinces(ctx context.Context) ([]model.Province, error) {
	p := []model.Province{}
	err := r.db.WithContext(ctx).Order("id asc").Find(&p).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *accountRepository) FindDistricts(ctx context.Context) ([]model.District, error) {
	p := []model.District{}
	err := r.db.WithContext(ctx).Order("id asc").Find(&p).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *accountRepository) FindDistrictsByProvinceId(ctx context.Context, ProvinceId int) ([]model.District, error) {
	d := []model.District{}
	err := r.db.WithContext(ctx).Order("id asc").Where("province_id = ?", ProvinceId).Find(&d).Error
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (r *accountRepository) FirstSeller(ctx context.Context, req dtorepository.SellerDataRequest) (dtorepository.SellerDataResponse, error) {
	res := dtorepository.SellerDataResponse{}

	q := `
		select 
			a.id as "Id",
			a.shop_name as "Name",
			a.profile_picture as "ProfilePicture",
			aa.district as "District",
			a.shop_opening_hours as "StartOperatingHours",
			a.shop_closing_hours as "EndOperatingHours",
			'asia/jakarta' as "TimeZone",
			TRIM(BOTH '-' FROM 
					REGEXP_REPLACE(
						REGEXP_REPLACE(
							LOWER(a.shop_name), 
							'[^a-z0-9]+', '-', 'g'
						), 
						'-+', '-', 'g'
					)
				) as "ShopNameSlug",
			a.shop_description as "Description",
			coalesce(shop_stars.stars_average, 0) as "Stars"
		from accounts a
		left join account_addresses aa 
			on aa.account_id  = a.id
		left join (
			select
				p.seller_id,
				round(avg(por.rating), 1) as stars_average
			from product_orders po 
			left join product_order_details pod 
				on pod.product_order_id = po.id 
			left join product_order_reviews por
				on por.product_order_detail_id = pod.id 
			left join products p 
				on p.id = pod.product_id 
			group by p.seller_id
		) shop_stars on shop_stars.seller_id = a.id
		where TRIM(BOTH '-' FROM 
					REGEXP_REPLACE(
						REGEXP_REPLACE(
							LOWER(a.shop_name), 
							'[^a-z0-9]+', '-', 'g'
						), 
						'-+', '-', 'g'
					)
				) ilike ?
			OR a.id = ?
	`

	if err := r.db.WithContext(ctx).Raw(q, req.ShopName, req.ShopId).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) FindSellerBestSelling(ctx context.Context, req dtorepository.FindSellerBestSellingRequest) (dtorepository.FindSellerBestSellingResponse, error) {
	res := dtorepository.FindSellerBestSellingResponse{}
	products := []dtorepository.SellerProduct{}

	q := `
		select 
			p.id as "Id",
			p."name" as  "Name",
			product_lowest_price.lowest_price as "Price",
			product_image.url as "PictureUrl",
			coalesce(product_rating.rating, 0) as "Stars",
			p.created_at as "CreatedAt",
			coalesce(product_sold.quantity, 0) as "TotalSold",
			case 
				when category_level_3.level_3_id is not null then category_level_3.level_2_name
				when category_level_2.level_2_id is not null then category_level_2.level_2_name
			end as "Category",
			seller.shop_name,
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(p."name"), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) AS "ProductNameSlug",
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(seller.shop_name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) AS "ShopNameSlug"
		from products p 
		left join (
			select
				pvsc.product_id,
				min (
					case
						when pvsc.price > 0 then pvsc.price 
						else null
					end
				) as lowest_price
			from product_variant_selection_combinations pvsc 
			group by pvsc.product_id
		) product_lowest_price on product_lowest_price.product_id = p.id 
		left join lateral (
			select
				pi2.product_id,
				pi2.url 
			from product_images pi2 
				where pi2.product_id = p.id
			order by pi2.id asc
			limit 1
		) product_image on product_image.product_id = p.id 
		left join (
			select
				c.id as level_1_id,
				c."name" level_1_name
			from categories c
			where c."level" = 1
		) as category_level_1 on category_level_1.level_1_id = p.category_id 
		left join (
			select
				c.id as level_2_id,
				c."name" level_2_name,
				c2.id as level_1_id,
				c2."name" as level_1_name
			from categories c
			inner join categories c2 
				on c2.id = c.parent 
			where c."level" = 2
		) as category_level_2 on category_level_2.level_2_id = p.category_id 
		left join (
			select
				c.id as level_3_id,
				c."name" level_3_name,
				c2.id as level_2_id,
				c2."name" level_2_name,
				c3.id as level_1_id,
				c3."name" as level_1_name
			from categories c
			inner join categories c2 
				on c2.id = c.parent 
			inner join categories c3
				on c3.id = c2.parent 
			where c."level" = 3
		) as category_level_3 on category_level_3.level_3_id = p.category_id 
		left join (
			select
				pod.product_id,
				sum(pod.quantity) as quantity
			from product_order_details pod
			group by pod.product_id 
		) as product_sold on product_sold.product_id = p.id
		left join (
			select
				por.product_id,
				round(avg(por.rating), 1) as rating
			from product_order_reviews por
			group by por.product_id 
		) as product_rating on product_rating.product_id = p.id
		left join accounts seller
			on seller.id = p.seller_id
		where p.seller_id = ?
	`

	err := r.db.WithContext(ctx).Raw(q, req.SellerId).Scan(&products).Error
	if err != nil {
		return res, err
	}

	res.Products = products

	return res, nil
}

func (r *accountRepository) FindSellerShowcases(ctx context.Context, req dtorepository.FindSellerShowcasesRequest) ([]dtorepository.FindSellerShowcasesResponse, error) {
	res := []dtorepository.FindSellerShowcasesResponse{}

	q := `
		select
			ss.id as "ShowcaseId",
			ss.name as "ShowcaseName"
		from seller_showcases ss
		where ss.seller_id = ?
	`

	err := r.db.WithContext(ctx).Raw(q, req.SellerId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) FindSellerShowcaseProduct(ctx context.Context, req dtorepository.FindSellerShowcaseProductRequest) (dtousecase.GetSellerShowcaseProductResponse, error) {
	res := dtousecase.GetSellerShowcaseProductResponse{}
	sp := []dtousecase.SellerProduct{}
	type count struct {
		TotalItem int
	}
	c := count{}
	countVar := 1

	q := `
		select
			p.id as "Id",
			p."name" as  "Name",
			product_lowest_price.lowest_price as "Price",
			product_image.url as "PictureUrl",
			coalesce(product_sold.quantity, 0) as "TotalSold",
			coalesce(product_rating.rating, 0) as "Stars",
			p.created_at as "CreatedAt",
			seller.shop_name as "ShopName",
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(p."name"), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) AS "ProductNameSlug",
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(seller.shop_name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) AS "ShopNameSlug"
		from products p
			left join (
				select
					pvsc.product_id,
					min (
						case
							when pvsc.price > 0 then pvsc.price 
							else null
						end
					) as lowest_price
				from product_variant_selection_combinations pvsc 
				group by pvsc.product_id
			) product_lowest_price on product_lowest_price.product_id = p.id 
			left join lateral (
				select
					pi2.product_id,
					pi2.url 
				from product_images pi2 
				where pi2.product_id = p.id
				limit 1
			) product_image on product_image.product_id = p.id 
			left join (
				select
					pod.product_id,
					sum(pod.quantity) as quantity
				from product_order_details pod
				group by pod.product_id 
			) as product_sold on product_sold.product_id = p.id
			left join (
				select
					por.product_id,
					round(avg(por.rating), 1) as rating
				from product_order_reviews por
				group by por.product_id 
			) as product_rating on product_rating.product_id = p.id
			left join accounts seller
				on seller.id = p.seller_id
			left join seller_showcase_products ssp
				on ssp.product_id = p.id
	`

	q += fmt.Sprint(`
		where TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(seller.shop_name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) = $`, countVar)
	countVar++

	wq := r.db.WithContext(ctx).Raw(q, req.ShopName)

	if req.ShowcaseId != "0" {
		q += fmt.Sprint(" and ssp.seller_showcase_id = $", countVar)
		countVar++
		wq = wq.Raw(q, req.ShowcaseId)
	}

	if err := r.db.
		WithContext(ctx).
		Raw(`select count(*) as "TotalItem" from (?) t`, wq).
		Scan(&c).Error; err != nil {
		return res, err
	}

	q += fmt.Sprint(` limit $`, countVar)
	countVar++
	wq = wq.Raw(q, req.Limit)

	q += fmt.Sprint(` offset $`, countVar)
	countVar++
	wq = wq.Raw(q, (req.Page-1)*req.Limit)

	if err := wq.Scan(&sp).Error; err != nil {
		return res, err
	}

	res.SellerProducts = sp
	res.CurrentPage = req.Page
	res.Limit = req.Limit
	res.TotalPage = int(math.Ceil(float64(c.TotalItem) / float64(req.Limit)))
	res.TotalItem = c.TotalItem

	return res, nil
}

func (r *accountRepository) FindCategories(ctx context.Context) ([]dtorepository.Category, error) {
	res := []dtorepository.Category{}

	q := `
		select 
			c.id as "CategoryLevel1Id",
			c."name" as "CategoryLevel1Name",
			c2.id as "CategoryLevel2Id",
			c2."name" as "CategoryLevel2Name",
			c3.id as "CategoryLevel3Id",
			c3."name" as "CategoryLevel3Name"
		from categories c
		left join categories c2 
			on c2.parent = c.id 
			and c2."level" = 2
		left join categories c3 
			on c3.parent = c2.id 
			and c3."level" = 3
		where c."level" = 1
		order by c.id asc, c2.id asc, c3.id asc
	`

	if err := r.db.WithContext(ctx).Raw(q).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) FindCategoryByID(ctx context.Context, categoryID int) (dtorepository.Category, error) {
	res := dtorepository.Category{}

	q := `
	select 
		c.id as "CategoryLevel1Id",
		c."name" as "CategoryLevel1Name",
		c2.id as "CategoryLevel2Id",
		c2."name" as "CategoryLevel2Name",
		c3.id as "CategoryLevel3Id",
		c3."name" as "CategoryLevel3Name"
	from categories c
	left join categories c2 
		on c2.parent = c.id 
		and c2."level" = 2
	left join categories c3 
		on c3.parent = c2.id 
		and c3."level" = 3
	where c2.id = $1 or c3.id = $1
	order by c.id asc, c2.id asc, c3.id asc
	`

	if err := r.db.WithContext(ctx).Raw(q, categoryID).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *accountRepository) UpdateShopProfile(ctx context.Context, req dtousecase.UpdateShopProfileRequest) (dtousecase.UpdateShopProfileResponse, error) {
	res := dtousecase.UpdateShopProfileResponse{}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		a := model.Accounts{}

		if err := tx.Where("id = ?", req.ShopId).Find(&a).Error; err != nil {
			return err
		}

		a.ShopName = req.ShopName
		a.ShopDescription = req.ShopDescription
		a.ShopOpeningHours = req.OpeningHours
		a.ShopClosingHours = req.ClosingHours

		if err := tx.Updates(&a).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return res, err
	}

	res = dtousecase.UpdateShopProfileResponse(req)

	return res, nil
}
