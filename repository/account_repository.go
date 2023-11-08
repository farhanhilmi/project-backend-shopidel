package repository

import (
	"context"
	"errors"
	"fmt"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type accountRepository struct {
	db                         *gorm.DB
	usedEmailRepo              usedEmailRepository
	walletTransactionHistories walletTransactionHistoryRepository
	saleTransactionHistories   saleWalletTransactionHistoryRepository
}
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

func (r *accountRepository) GetAddresses(ctx context.Context, req dtorepository.AddressRequest) (*[]dtorepository.AddressResponse, error) {
	res := []dtorepository.AddressResponse{}
	addresses := []model.AccountAddress{}

	err := r.db.WithContext(ctx).Find(&addresses).Where("account_id = ?", req.UserId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &res, util.ErrNoRecordFound
	}

	for _, address := range addresses {
		convertedFullAddress := fmt.Sprintf("%s, %s, %s, %s, %s.",
			address.Detail,
			address.Kelurahan,
			address.SubDistrict,
			address.District,
			address.Province,
		)
		res = append(res, dtorepository.AddressResponse{
			ID:              address.ID,
			FullAddress:     convertedFullAddress,
			IsBuyerDefault:  address.IsBuyerDefault,
			IsSellerDefault: address.IsSellerDefault,
		})
	}

	return &res, nil
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
	a.ProfilePicture = req.ProfilePicture

	err = tx.WithContext(ctx).Save(&a).Error

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
	res.ProfilePicture = a.ProfilePicture

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
	a.ProfilePicture = req.ProfilePicture

	err = r.db.WithContext(ctx).Save(&a).Error

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
		From:      "5550000012345",
	})

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
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
			pvsc.picture_url as "ProductUrl",
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
			left join accounts seller
				on seller.id = p.seller_id 
		where ac.account_id = 2
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

	c := model.AccountCarts{}

	err = r.db.WithContext(ctx).Where("account_id = $1 and product_variant_selection_combination_id = $2", req.AccountId, req.ProductVariantCombinationId).First(&c).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if c.ID == 0 {
		c = model.AccountCarts{
			AccountID:                            req.AccountId,
			ProductVariantSelectionCombinationId: req.ProductVariantCombinationId,
			Quantity:                             req.Quantity,
		}

		err = r.db.WithContext(ctx).Create(&c).Error
		if err != nil {
			return res, err
		}
	} else {
		c.Quantity += req.Quantity

		err = r.db.WithContext(ctx).Save(&c).Error
		if err != nil {
			return res, err
		}
	}

	res.AccountId = req.AccountId
	res.Quantity = c.Quantity
	res.ProductVariantCombinationId = c.ProductVariantSelectionCombinationId

	return res, nil
}
