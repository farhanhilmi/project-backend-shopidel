package repository

import (
	"context"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}
type ProductRepository interface {
	FindByID(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindByID(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error) {

	res := dtorepository.ProductResponse{}

	return res, nil
}
