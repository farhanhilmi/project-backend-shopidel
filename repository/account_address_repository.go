package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type accountAddressRepository struct {
	db *gorm.DB
}

type AccountAddressRepository interface {
	FindBuyerAddressByID(ctx context.Context, req dtorepository.AccountAddressRequest) (dtorepository.AccountAddressResponse, error)
}

func NewAccountAddressRepository(db *gorm.DB) AccountAddressRepository {
	return &accountAddressRepository{
		db: db,
	}
}

func (r *accountAddressRepository) FindBuyerAddressByID(ctx context.Context, req dtorepository.AccountAddressRequest) (dtorepository.AccountAddressResponse, error) {
	accountAddress := model.AccountAddress{}
	res := dtorepository.AccountAddressResponse{}

	err := r.db.WithContext(ctx).Where("id = ?", req.ID).First(&accountAddress).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	if err != nil {
		return res, err
	}

	res.ID = accountAddress.ID
	res.AccountID = accountAddress.AccountID
	res.Province = accountAddress.Province
	res.SubDistrict = accountAddress.SubDistrict
	res.District = accountAddress.District
	res.Kelurahan = accountAddress.Kelurahan
	res.ZipCode = accountAddress.ZipCode
	res.Detail = accountAddress.Detail
	res.IsBuyerDefault = accountAddress.IsBuyerDefault
	res.IsSellerDefault = accountAddress.IsSellerDefault
	return res, err
}
