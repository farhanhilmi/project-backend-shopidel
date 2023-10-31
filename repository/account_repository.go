package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

type AccountRepository interface {
	GetDetail(ctx context.Context) (*dto.AccountResponse, error)
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) GetDetail(ctx context.Context) (*dto.AccountResponse, error) {
	getBalance, _ := decimal.NewFromString("100500")

	data := &dto.AccountResponse{
		ID: 1,
		Username: "username",
		FullName: "fullname",
		Email: "email@gmail.com",
		PhoneNumber: "+6298374829",
		Gender: "male",
		Birthdate: "02/10/2000",
		ProfilePicture: "https://img.wattpad.com/98983ce58966193d9d0a4a74c2a00cb0542d47c8/68747470733a2f2f73332e616d617a6f6e6177732e636f6d2f776174747061642d6d656469612d736572766963652f53746f7279496d6167652f684b6c38305779434836646179773d3d2d313030393331303738362e3136353933666238663832303237363835313535333137373639382e6a7067",
		Balance: getBalance,
	}

	return data, nil
}
