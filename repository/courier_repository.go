package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type courierRepository struct {
	db *gorm.DB
}

type CourierRepository interface {
	FindByName(ctx context.Context, req dtorepository.CourierData) (dtorepository.CourierData, error)
}

func NewCourierRepository(db *gorm.DB) CourierRepository {
	return &courierRepository{
		db: db,
	}
}

func (r *courierRepository) FindByName(ctx context.Context, req dtorepository.CourierData) (dtorepository.CourierData, error) {
	courier := model.Couriers{}
	res := dtorepository.CourierData{}

	err := r.db.WithContext(ctx).Where("name = ?", req.Name).First(&courier).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	res.ID = courier.ID
	res.Name = courier.Name
	res.Description = courier.Description
	res.CreatedAt = courier.CreatedAt
	res.DeletedAt = courier.DeletedAt
	res.UpdatedAt = courier.UpdatedAt

	return res, nil
}
