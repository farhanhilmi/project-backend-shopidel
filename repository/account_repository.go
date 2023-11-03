package repository

import (
	"context"
	"errors"

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
	RefundBalance(ctx context.Context, tx *gorm.DB, req dtorepository.MyWalletRequest) (dtorepository.WalletResponse, error)
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
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
		}}).Model(&account).Where("id = ?", req.UserID).Update("balance", gorm.Expr("balance + ?", req.Balance)).Error

	if err != nil {
		tx.Rollback()
		return dtorepository.WalletResponse{}, err
	}

	_, err = r.walletTransactionHistories.CreateWithTx(ctx, tx, dtorepository.MyWalletTransactionHistoriesRequest{
		AccountID: req.UserID,
		Amount:    req.Balance,
		Type:      "Refund",
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

	_, err = r.walletTransactionHistories.CreateWithTx(ctx, tx, dtorepository.MyWalletTransactionHistoriesRequest{
		AccountID: req.UserID,
		Amount:    req.Amount,
		Type:      req.Type,
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

	err := r.db.WithContext(ctx).Where("username = ?", req.Username).First(&account).Error

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
