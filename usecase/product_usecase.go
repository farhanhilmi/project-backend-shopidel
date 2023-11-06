package usecase

import (
	"context"
	"errors"

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

	options, err := u.convertVariantOptions(ctx, rRes2)
	if err != nil {
		return res, err
	}

	res.Id = rRes.ID
	res.ProductName = rRes.Name
	res.Description = rRes.Description
	res.Variants = variants
	res.VariantOptions = options

	return res, nil
}

func (u *productUsecase) convertProductVariants(ctx context.Context, productName string, req dtorepository.FindProductVariantResponse) ([]dtousecase.ProductVariant, error) {
	res := []dtousecase.ProductVariant{}

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

			pv.Selections = append(pv.Selections, dtousecase.ProductSelection{SelectionName: data.SelectionName1, SelectionVariantName: data.VariantName1})
			if data.SelectionId2 != 0 {
				pv.Selections = append(pv.Selections, dtousecase.ProductSelection{SelectionName: data.SelectionName2, SelectionVariantName: data.VariantName2})
			}
		}

		res = append(res, pv)
	}

	return res, nil
}

func (u *productUsecase) convertVariantOptions(ctx context.Context, req dtorepository.FindProductVariantResponse) ([]dtousecase.VariantOption, error) {
	res := []dtousecase.VariantOption{}
	m := map[string]map[string]string{}

	for _, data := range req.Variants {
		if data.SelectionName1 != "default_reserved_keyword" {
			if m[data.VariantName1] != nil {
				if m[data.VariantName1][data.SelectionName1] == "" {
					m[data.VariantName1][data.SelectionName1] = data.SelectionName1
				}
			} else {
				m[data.VariantName1] = map[string]string{
					data.SelectionName1: data.SelectionName1,
				}
			}

			if data.SelectionId2 != 0 {
				if m[data.VariantName2] != nil {
					if m[data.VariantName2][data.SelectionName2] == "" {
						m[data.VariantName2][data.SelectionName2] = data.SelectionName2
					}
				} else {
					m[data.VariantName2] = map[string]string{
						data.SelectionName2: data.SelectionName2,
					}
				}
			}
		}
	}

	for key, value := range m {
		vos := []string{}

		for key2 := range value {
			vos = append(vos, key2)
		}

		res = append(res, dtousecase.VariantOption{
			VariantOptionName: key,
			Childs:            vos,
		})
	}

	return res, nil
}
