package usecase

import (
	"context"
	"errors"
	"fmt"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
)

type ProductUsecase interface {
	GetProductDetail(ctx context.Context, req dtousecase.GetProductDetailRequest) (*dtousecase.GetProductDetailResponse, error)
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

func (u *productUsecase) GetProductDetail(ctx context.Context, req dtousecase.GetProductDetailRequest) (*dtousecase.GetProductDetailResponse, error) {
	res := &dtousecase.GetProductDetailResponse{}

	rReq := dtorepository.ProductRequest{
		ProductID: req.ProductId,
	}
	rRes, err := u.productRepository.First(ctx, rReq)
	if err != nil {
		return res, err
	}
	if rRes.ID == 0 {
		return res, errors.New("product not found")
	}

	rReq2 := dtorepository.FindProductVariantRequest{
		ProductId: rRes.ID,
	}
	rRes2, err := u.productRepository.FindProductVariant(ctx, rReq2)
	if err != nil {
		return res, err
	}

	variants, err := u.convertProductVariants(ctx, rRes.Name, rRes2)
	if err != nil {
		return res, err
	}

	res.Id = rRes.ID
	res.ProductName = rRes.Name
	res.Description = rRes.Description
	res.Variants = variants

	return res, nil
}

func (u *productUsecase) convertProductVariants(ctx context.Context, productName string, req dtorepository.FindProductVariantResponse) ([]dtousecase.ProductVariant, error) {
	res := []dtousecase.ProductVariant{}
	fmt.Println(req)
	for _, data := range req.Variants {
		pv := dtousecase.ProductVariant{}
		pv.VariantId = data.VariantId
		pv.Price = data.Price
		pv.Stock = data.Stock

		if data.SelectionName1 == "default_reserved_keyword" {
			pv.VariantName = productName
		} else {
			pv.VariantName = productName + " - " + data.SelectionName1

			if data.SelectionId2 != 0 {
				pv.VariantName = pv.VariantName + ", " + data.SelectionName2
			}

			pv.Selections = append(pv.Selections, dtousecase.ProductSelection{SelectionName: data.SelectionName1})
			if data.SelectionId2 != 0 {
				pv.Selections = append(pv.Selections, dtousecase.ProductSelection{SelectionName: data.SelectionName2})
			}
		}

		res = append(res, pv)
	}
	fmt.Println(res)
	return res, nil
}
