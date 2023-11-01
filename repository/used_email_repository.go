package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type usedEmailRepository struct {
	db *gorm.DB
}

type UsedEmailRepository interface {
	FindByEmail(ctx context.Context, req dtorepository.UsedEmailRequest) (dtorepository.UsedEmailResponse, error)
	CreateEmail(ctx context.Context, tx *gorm.DB, req dtorepository.UsedEmailRequest) (dtorepository.UsedEmailResponse, error)
}

func NewUsedEmailRepository(db *gorm.DB) UsedEmailRepository {
	return &usedEmailRepository{
		db: db,
	}
}

func (r *usedEmailRepository) FindByEmail(ctx context.Context, req dtorepository.UsedEmailRequest) (dtorepository.UsedEmailResponse, error) {
	usedEmail := model.UsedEmail{}
	res := dtorepository.UsedEmailResponse{}

	err := r.db.WithContext(ctx).Where("email = ?", req.Email).First(&usedEmail).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.Email = usedEmail.Email
	res.ID = usedEmail.ID
	res.CreatedAt = usedEmail.CreatedAt
	res.DeletedAt = usedEmail.DeletedAt
	res.UpdatedAt = usedEmail.UpdatedAt
	res.AccountID = usedEmail.AccountID

	return res, err
}

func (r *usedEmailRepository) CreateEmail(ctx context.Context, tx *gorm.DB, req dtorepository.UsedEmailRequest) (dtorepository.UsedEmailResponse, error) {
	res := dtorepository.UsedEmailResponse{}
	
	usedEmail := model.UsedEmail{
		AccountID: req.AccountID,
		Email: req.Email,
	}

	err := tx.WithContext(ctx).Model(&usedEmail).Create(&usedEmail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.AccountID = usedEmail.AccountID
	res.Email = usedEmail.Email

	return res, err
}