package usecase

import (
	"context"
	"errors"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type ProductUsecase interface {
	GetProductDetail(ctx context.Context, req dtousecase.GetProductDetailRequest) (*dtousecase.GetProductDetailResponse, error)
	GetProductDetailV2(ctx context.Context, req dtousecase.GetProductDetailRequestV2) (*dtousecase.GetProductDetailResponse, error)
	AddToFavorite(ctx context.Context, req dtousecase.FavoriteProduct) (*dtousecase.FavoriteProduct, error)
	GetProductFavorites(ctx context.Context, req dtousecase.ProductFavoritesParams) ([]dtousecase.GetFavoriteProductListResponse, *dtogeneral.PaginationData, error)
	GetProducts(ctx context.Context, req dtousecase.ProductListParam) (*[]dtorepository.ProductListResponse, *dtogeneral.PaginationData, error)
	GetProductReviews(ctx context.Context, req dtousecase.GetProductReviewsRequest) (dtousecase.GetProductReviewsResponse, error)
	GetProductPictures(ctx context.Context, req dtousecase.GetProductPicturesRequest) (*dtousecase.GetProductPicturesResponse, error)
	GetProductDetailRecomendedProducts(ctx context.Context, req dtousecase.GetProductDetailRecomendedProductRequest) (*dtousecase.GetProductDetailRecomendedProductResponse, error)
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

func (u *productUsecase) GetProducts(ctx context.Context, req dtousecase.ProductListParam) (*[]dtorepository.ProductListResponse, *dtogeneral.PaginationData, error) {
	uReq := dtorepository.ProductListParam{
		CategoryId: req.CategoryId,
		AccountID:  req.AccountID,
		SortBy:     req.SortBy,
		Search:     req.Search,
		Sort:       req.Sort,
		MinRating:  req.MinRating,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		District:   req.District,
		Limit:      req.Limit,
		Page:       req.Page,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}

	listProduct, totalItems, err := u.productRepository.FindProducts(ctx, uReq)
	if err != nil {
		return &listProduct, nil, err
	}

	pagination := dtogeneral.PaginationData{
		TotalItem:   int(totalItems),
		TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}

	return &listProduct, &pagination, nil
}

func (u *productUsecase) AddToFavorite(ctx context.Context, req dtousecase.FavoriteProduct) (*dtousecase.FavoriteProduct, error) {
	_, err := u.productRepository.First(ctx, dtorepository.ProductRequest{
		ProductID: req.ProductID,
	})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}

	_, err = u.productRepository.FindProductFavorites(ctx, dtorepository.FavoriteProduct{ProductID: req.ProductID, AccountID: req.AccountID})
	if err != nil && !errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if !errors.Is(err, util.ErrNoRecordFound) {
		favorite, err := u.productRepository.RemoveProductFavorite(ctx, dtorepository.FavoriteProduct{ProductID: req.ProductID, AccountID: req.AccountID})
		if errors.Is(err, util.ErrNoRecordFound) {
			return nil, util.ErrProductNotFound
		}
		if err != nil {
			return nil, err
		}
		return &dtousecase.FavoriteProduct{
			ID:        favorite.ID,
			ProductID: favorite.ProductID,
			AccountID: favorite.AccountID,
		}, nil
	}

	favorite, err := u.productRepository.AddProductFavorite(ctx, dtorepository.FavoriteProduct{
		ProductID: req.ProductID,
		AccountID: req.AccountID,
	})
	if err != nil {
		return nil, err
	}

	return &dtousecase.FavoriteProduct{
		ID:        favorite.ID,
		ProductID: favorite.ProductID,
		AccountID: favorite.AccountID,
	}, nil
}

func (u *productUsecase) GetProductFavorites(ctx context.Context, req dtousecase.ProductFavoritesParams) ([]dtousecase.GetFavoriteProductListResponse, *dtogeneral.PaginationData, error) {
	res := []dtousecase.GetFavoriteProductListResponse{}

	products, totalItems, err := u.productRepository.FindAllProductFavorites(ctx, dtorepository.ProductFavoritesParams{
		AccountID: req.AccountID,
		SortBy:    req.SortBy,
		Search:    req.Search,
		Sort:      req.Sort,
		Limit:     req.Limit,
		Page:      req.Page,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return res, nil, err
	}

	pagination := dtogeneral.PaginationData{
		TotalItem:   int(totalItems),
		TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}

	return products, &pagination, nil
}

func (u *productUsecase) GetProductDetail(ctx context.Context, req dtousecase.GetProductDetailRequest) (*dtousecase.GetProductDetailResponse, error) {
	res := &dtousecase.GetProductDetailResponse{}

	rReq := dtorepository.ProductRequest{
		ProductID: req.ProductId,
		AccountId: req.AccountId,
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
	res.IsFavorite = rRes.IsFavorite

	return res, nil
}

func (u *productUsecase) GetProductDetailV2(ctx context.Context, req dtousecase.GetProductDetailRequestV2) (*dtousecase.GetProductDetailResponse, error) {
	res := &dtousecase.GetProductDetailResponse{}

	rReq := dtorepository.ProductRequestV2{
		AccountId:   req.AccountId,
		ShopName:    req.ShopName,
		ProductName: req.ProductName,
	}
	rRes, err := u.productRepository.FirstV2(ctx, rReq)
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
	res.IsFavorite = rRes.IsFavorite

	return res, nil
}

func (u *productUsecase) GetProductPictures(ctx context.Context, req dtousecase.GetProductPicturesRequest) (*dtousecase.GetProductPicturesResponse, error) {
	res := &dtousecase.GetProductPicturesResponse{}

	rRes, err := u.productRepository.FindImages(ctx, req.ProductId)
	if err != nil {
		return res, err
	}

	for _, picture := range rRes.ProductPictures {
		res.PicturesUrl = append(res.PicturesUrl, picture.PictureUrl)
	}

	return res, nil
}

func (u *productUsecase) GetProductDetailRecomendedProducts(ctx context.Context, req dtousecase.GetProductDetailRecomendedProductRequest) (*dtousecase.GetProductDetailRecomendedProductResponse, error) {
	res := &dtousecase.GetProductDetailRecomendedProductResponse{}

	product, err := u.productRepository.First(ctx, dtorepository.ProductRequest{ProductID: req.ProductId})
	if err != nil {
		return res, err
	}

	anotherProducts, err := u.productRepository.FindSellerAnotherProducts(ctx, product.SellerId)
	if err != nil {
		return res, err
	}

	res.AnotherProducts = anotherProducts

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
	m2 := map[string]map[string]string{}
	p := map[string]string{}

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
				if m2[data.VariantName2] != nil {
					if m2[data.VariantName2][data.SelectionName2] == "" {
						m2[data.VariantName2][data.SelectionName2] = data.SelectionName2
					}
				} else {
					m2[data.VariantName2] = map[string]string{
						data.SelectionName2: data.SelectionName2,
					}
				}
			}
		}
	}

	for _, data := range req.Variants {
		if p[data.SelectionName1] == "" {
			p[data.SelectionName1] = data.ImageURL
		}
	}

	for key, value := range m {
		vos := []string{}
		variantsPicture := []string{}

		for key2 := range value {
			vos = append(vos, key2)
			if p[key2] != "" {
				variantsPicture = append(variantsPicture, p[key2])
			}
		}

		res = append(res, dtousecase.VariantOption{
			VariantOptionName: key,
			Childs:            vos,
			Pictures:          variantsPicture,
		})
	}

	for key, value := range m2 {
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

func (u *productUsecase) GetProductReviews(ctx context.Context, req dtousecase.GetProductReviewsRequest) (dtousecase.GetProductReviewsResponse, error) {
	res := dtousecase.GetProductReviewsResponse{}
	rRes, err := u.productRepository.FindProductReviews(ctx, req)
	if err != nil {
		return res, err
	}

	for _, data := range rRes.Reviews {
		review := dtousecase.ProductReview{
			Id:                 data.Id,
			CustomerName:       data.CustomerName,
			CustomerPictureUrl: data.CustomerPictureUrl,
			Stars:              data.Stars,
			Comment:            data.Comment,
			Variant:            data.Variant,
			CreatedAt:          data.CreatedAt,
		}

		pictures, err := u.productRepository.FindProductReviewPictures(ctx, data.Id)
		if err != nil {
			return res, err
		}

		for _, picture := range pictures {
			review.Pictures = append(review.Pictures, picture.Url)
		}

		res.Reviews = append(res.Reviews, review)
	}

	res.Limit = req.Limit
	res.CurrentPage = req.Page
	res.TotalItem = rRes.TotalItem
	res.TotalPage = rRes.TotalPage

	return res, nil
}
