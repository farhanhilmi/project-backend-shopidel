package usecase

import (
	"context"
	"errors"
	"log"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type ProductUsecase interface {
	GetProduct(ctx context.Context, req dtousecase.ProductRequest) (*dtousecase.ProductResponse, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

type ProductUsecaseConfig struct {
	ProductRepository repository.ProductRepository
}

func NewProductUsecase(config ProductUsecaseConfig) ProductUsecase {
	au := &productUsecase{}
	if config.ProductRepository != nil {
		au.productRepository = config.ProductRepository
	}

	return au
}

func (u *productUsecase) GetProduct(ctx context.Context, req dtousecase.ProductRequest) (*dtousecase.ProductResponse, error) {
	res := dtousecase.ProductResponse{}

	product, err := u.productRepository.FindByID(ctx, dtorepository.ProductRequest{ProductID: req.ProductID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}
	log.Println("PRODUCT", product)

	res.ID = product.ID
	res.Category = product.Category
	res.Description = product.Description
	res.HazardousMaterial = product.HazardousMaterial
	res.InternalSKU = product.InternalSKU
	res.IsActive = product.IsActive
	res.IsNew = product.IsNew
	res.Size = product.Size
	res.ViewCount = product.ViewCount
	res.Name = product.Name
	res.Weight = product.Weight
	res.CreatedAt = product.CreatedAt
	res.DeletedAt = product.DeletedAt
	res.UpdatedAt = product.UpdatedAt

	return &res, nil
}
