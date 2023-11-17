package usecase

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type SellerUsecase interface {
	GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error)
	GetBestSelling(ctx context.Context, req dtousecase.GetSellerProductsRequest) (dtousecase.GetSellerProductsResponse, error)
	GetCategories(ctx context.Context, req dtousecase.GetSellerCategoriesRequest) (dtousecase.GetSellerCategoriesResponse, error)
	GetCategoryProducts(ctx context.Context, req dtousecase.GetSellerCategoryProductRequest) (dtousecase.GetSellerCategoryProductResponse, error)
	AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error)
	UploadPhoto(ctx context.Context, req dtousecase.UploadNewPhoto) (dtousecase.UploadNewPhoto, error)
}

type sellerUsecase struct {
	accountRepository repository.AccountRepository
	productRepository repository.ProductRepository
}

type SellerUsecaseConfig struct {
	AccountRepository repository.AccountRepository
	ProductRepository repository.ProductRepository
}

func NewSellerUsecase(config SellerUsecaseConfig) SellerUsecase {
	au := &sellerUsecase{}

	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
	}
	if config.ProductRepository != nil {
		au.productRepository = config.ProductRepository
	}

	return au
}

func (u *sellerUsecase) GetProfile(ctx context.Context, req dtousecase.GetSellerProfileRequest) (dtousecase.GetSellerProfileResponse, error) {
	res := dtousecase.GetSellerProfileResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	res.SellerDistrict = seller.District
	res.SellerName = seller.Name
	res.SellerOperatingHour = dtousecase.SellerOperatingHour{
		Start: seller.StartOperatingHours,
		End:   seller.EndOperatingHours,
	}
	res.SellerPictureUrl = seller.ProfilePicture

	return res, nil
}

func (u *sellerUsecase) GetBestSelling(ctx context.Context, req dtousecase.GetSellerProductsRequest) (dtousecase.GetSellerProductsResponse, error) {
	res := dtousecase.GetSellerProductsResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	rRes, err := u.accountRepository.FindSellerBestSelling(ctx, dtorepository.FindSellerBestSellingRequest{SellerId: seller.Id})
	if err != nil {
		return res, err
	}

	for _, data := range rRes.Products {
		p := dtousecase.SellerProduct{
			Name:       data.Name,
			Price:      data.Price,
			PictureUrl: data.PictureUrl,
			Stars:      data.Stars,
			TotalSold:  data.TotalSold,
			CreatedAt:  data.CreatedAt,
			Category:   data.Category,
			ShopName:   data.ShopName,
		}

		res.SellerProducts = append(res.SellerProducts, p)
	}

	return res, nil
}

func (u *sellerUsecase) GetCategories(ctx context.Context, req dtousecase.GetSellerCategoriesRequest) (dtousecase.GetSellerCategoriesResponse, error) {
	res := dtousecase.GetSellerCategoriesResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	rRes, err := u.accountRepository.FindSellerCategories(ctx, dtorepository.FindSellerCategoriesRequest{SellerId: seller.Id})
	if err != nil {
		return res, err
	}

	for _, data := range rRes {
		res.Categories = append(res.Categories, dtousecase.SellerCategory{CategoryId: data.CategoryId, CategoryName: data.CategoryName})
	}

	return res, nil
}

func (u *sellerUsecase) GetCategoryProducts(ctx context.Context, req dtousecase.GetSellerCategoryProductRequest) (dtousecase.GetSellerCategoryProductResponse, error) {
	res := dtousecase.GetSellerCategoryProductResponse{}

	seller, err := u.accountRepository.FirstSeller(ctx, dtorepository.SellerDataRequest{ShopName: req.ShopName})
	if err != nil {
		return res, err
	}

	if seller.Id == 0 {
		return res, util.ErrSellerNotFound
	}

	products, err := u.accountRepository.FindSellerCategoryProduct(ctx, dtorepository.FindSellerCategoryProductRequest{ShopName: req.ShopName, CategoryId: req.CategoryId})
	if err != nil {
		return res, err
	}

	res.SellerProducts = products

	return res, nil
}

func (u *sellerUsecase) AddNewProduct(ctx context.Context, req dtousecase.AddNewProductRequest) (dtousecase.AddNewProductResponse, error) {
	res := dtousecase.AddNewProductResponse{}

	productVariants := []dtousecase.ProductVariants{}

	if len(req.Variants) == 1 {
		switch {
		case req.Variants[0].Variant1.Name == "":
			productVariants = []dtousecase.ProductVariants{
				{
					Name: constant.DefaultReservedKeyword,
				},
			}
			req.Variants[0].Variant1.Name = constant.DefaultReservedKeyword
			req.Variants[0].Variant1.Value = constant.DefaultReservedKeyword
		case req.Variants[0].Variant2.Name == "":
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
			}
		default:
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
				{
					Name: req.Variants[0].Variant2.Name,
				},
			}
		}
	}

	if len(req.Variants) > 1 {
		if req.Variants[0].Variant2.Name == "" {
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
			}
		} else {
			productVariants = []dtousecase.ProductVariants{
				{
					Name: req.Variants[0].Variant1.Name,
				},
				{
					Name: req.Variants[0].Variant2.Name,
				},
			}
		}
	}

	product, err := u.productRepository.AddNewProduct(ctx, dtorepository.AddNewProductRequest{
		SellerID:          req.SellerID,
		ProductName:       req.ProductName,
		Description:       req.Description,
		CategoryID:        req.CategoryID,
		HazardousMaterial: req.HazardousMaterial,
		IsNew:             req.IsNew,
		InternalSKU:       req.InternalSKU,
		Weight:            req.Weight,
		IsActive:          req.IsActive,
		Size:              req.Size,
		Variants:          req.Variants,
		ProductVariants:   productVariants,
		Images:            req.Images,
		VideoURL:          req.VideoURL,
	})
	if err != nil {
		return res, err
	}

	//

	// for _, v := range req.Variants {
	// 	file, err := os.Open(fmt.Sprintf("./imageuploads/%v.png", v.ImageID))
	// 	if err != nil {
	// 		fmt.Println("Error opening file:", err)
	// 		return res, err
	// 	}

	// 	defer file.Close()

	// 	currentTime := time.Now().UnixNano()
	// 	fileExtension := path.Ext(file.Name())
	// 	originalFilename := file.Name()[:len(file.Name())-len(fileExtension)]
	// 	newFilename := fmt.Sprintf("%s_%d", originalFilename, currentTime)
	// 	fileName := strings.Split(newFilename, "./imageuploads/")

	// 	imageUrl, err := util.UploadToCloudinary(file, fileName[0])
	// 	if err != nil {
	// 		return res, err
	// 	}

	// 	v.ImageURL = imageUrl
	// }

	res.ProductName = product.Name

	return res, nil
}

func (u *sellerUsecase) UploadPhoto(ctx context.Context, req dtousecase.UploadNewPhoto) (dtousecase.UploadNewPhoto, error) {
	res := dtousecase.UploadNewPhoto{}

	err := os.MkdirAll("./imageuploads", os.ModePerm)
	if err != nil {
		return res, err
	}

	dst, err := os.Create(fmt.Sprintf("./imageuploads/%s%s", req.ImageID, filepath.Ext(req.ImageHeader.Filename)))
	if err != nil {
		return res, err
	}

	defer dst.Close()

	_, err = io.Copy(dst, req.Image)
	if err != nil {
		return res, err
	}

	res.ImageID = req.ImageID
	return res, nil
}
