package dtorepository

import (
	"time"

	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	ProductID int
	AccountId int
}

type ProductRequestV2 struct {
	AccountId   int
	ShopName    string
	ProductName string
}

type ProductResponse struct {
	ID                int
	Name              string
	Description       string
	IsFavorite        bool
	SellerId          int
	CategoryID        int
	HazardousMaterial *bool
	IsNew             *bool
	InternalSKU       string
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsActive          *bool
	VideoURL          string
	Stars             decimal.Decimal
	Sold              int
}

type ProductImages struct {
	URL string
}

type ProductLowestHighestPriceRequest struct {
	ProductId int
}

type ProductLowestHighestPriceResponse struct {
	LowestPrice  decimal.Decimal
	HighestPrice decimal.Decimal
}

type ProductTotalSoldRequest struct {
	ProductId int
}

type ProductTotalSoldResponse struct {
	TotalSold int
}

type ProductAverageStarsRequest struct {
	ProductId int
}

type ProductAverageStarsResponse struct {
	AverageStars decimal.Decimal
}

type ProductVariantRequest struct {
	ProductId  int
	VariantId1 int
	VariantId2 int
}

type ProductSelection struct {
	Id          int
	Name        string
	IsAvailable bool
}

type FindProductVariantRequest struct {
	ProductId int
}

type FindProductVariantResponse struct {
	Variants []ProductVariant
}

type ProductVariant struct {
	ID             int
	VariantId      int
	SelectionId1   int
	SelectionId2   int
	SelectionName1 string
	SelectionName2 string
	VariantName1   string
	VariantName2   string
	ImageURL       string
	Price          decimal.Decimal
	Stock          int
}

type ProductCombinationVariantListRequest struct {
	ID []int
}

type ProductCombinationVariantRequest struct {
	ID int
}

type ProductCombinationVariantRespponse struct {
	ID              int
	IndividualPrice decimal.Decimal
	VariantName1    string
	VariantName2    string
	Stock           int
	ProductID       int
}

type UpdateCart struct {
	ProductID int
	Quantity  int
}

type ProductCart struct {
	ID        int
	ProductID int
	Quantity  int
	AccountID int
}

type FavoriteProduct struct {
	ID        int
	ProductID int
	AccountID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type FavoriteProductResponse struct {
	ID          int
	ProductID   int
	AccountID   int
	Name        string
	Description string
	Price       decimal.Decimal
	PictureURL  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type ProductFavoritesParams struct {
	AccountID int
	SortBy    string
	Search    string
	Sort      string
	Limit     int
	Page      int
	StartDate string
	EndDate   string
}

type ProductListResponse struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	District        string          `json:"district"`
	TotalSold       int             `json:"total_sold"`
	Price           decimal.Decimal `json:"price"`
	Rating          float64         `json:"rating"`
	PictureURL      string          `json:"picture_url"`
	CategoryId      int             `json:"category_id"`
	CategoryName    string          `json:"category_name"`
	ShopName        string          `json:"shop_name"`
	ProductNameSlug string          `json:"product_name_slug"`
	ShopNameSlug    string          `json:"shop_name_slug"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `json:"deleted_at"`
}

type ProductListSellerResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type ProductListParam struct {
	CategoryId string
	AccountID  int
	SellerID   int
	SortBy     string
	Search     string
	Sort       string
	MinRating  int
	MinPrice   int
	MaxPrice   int
	District   string
	Limit      int
	Page       int
	StartDate  string
	EndDate    string
}
type IsProductFavoriteRequest struct {
	AccountId int
	ProductId int
}

type IsProductFavoriteResponse struct {
	AccountId  int
	ProductId  int
	IsFavorite bool
}

type FindProductPicturesResponse struct {
	ProductPictures []ProductPicture
}

type ProductPicture struct {
	PictureUrl string
}

type TopCategoriesResponse struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	PictureURL string `json:"picture_url"`
}

type AddNewProductRequest struct {
	ProductID         int
	SellerID          int
	ProductName       string
	Description       string
	CategoryID        int
	HazardousMaterial *bool
	IsNew             *bool
	InternalSKU       string
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsActive          *bool
	Variants          []dtousecase.AddNewProductVariant
	Images            []string
	DeletedImages     []string
	VideoURL          string
	ProductVariants   []dtousecase.ProductVariants
}

type UpdateProductRequest struct {
	ProductID         int
	SellerID          int
	ProductName       string
	Description       string
	CategoryID        int
	HazardousMaterial *bool
	IsNew             *bool
	InternalSKU       string
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsActive          *bool
	Variants          []dtousecase.UpdateProductVariant
	Images            []string
	DeletedImages     []string
	VideoURL          string
	ProductVariants   []dtousecase.ProductVariants
}

type AddNewProductResponse struct {
	ID                int
	Name              string
	Description       string
	CategoryID        int
	HazardousMaterial *bool
	IsNew             *bool
	InternalSKU       string
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsActive          *bool
	VideoURL          string
}

type ProductVariantType struct {
	VariantName  string
	VariantValue string
}

type GetProductReviewsResponse struct {
	Reviews     []ProductReview
	TotalPage   int
	TotalItem   int
	CurrentPage int
	Limit       int
}

type ProductReview struct {
	Id                 int
	CustomerName       string
	CustomerPictureUrl string
	Stars              decimal.Decimal
	Comment            string
	Variant            string
	CreatedAt          string
}

type RemoveProduct struct {
	ID        int
	SellerID  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
