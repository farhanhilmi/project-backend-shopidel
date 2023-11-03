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

type productOrdersRepository struct {
	db *gorm.DB
}

type ProductOrdersRepository interface {
	FindByIDAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
	FindByStatusAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
	UpdateOrderStatusByIDAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
	FindByID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
}

func NewProductOrdersRepository(db *gorm.DB) ProductOrdersRepository {
	return &productOrdersRepository{
		db: db,
	}
}

func (r *productOrdersRepository) FindByID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Where("id = ?", req.ID).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, err
}

func (r *productOrdersRepository) FindByIDAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountID).Where("id = ?", req.ID).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, err
}

func (r *productOrdersRepository) FindByStatusAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountID).Where("status = ?", req.Status).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, err
}

func (r *productOrdersRepository) UpdateOrderStatusByIDAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&order).Where("id = ?", req.ID).Where("account_id = ?", req.AccountID).Update("status", req.Status).Error

	if err != nil {
		return res, err
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, nil
}
