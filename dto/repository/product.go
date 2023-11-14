package dtorepository

import (
	"time"

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
	ID          int
	Name        string
	Description string
	IsFavorite  bool
	SellerId    int
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
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	District     string          `json:"district"`
	TotalSold    int             `json:"total_sold"`
	Price        decimal.Decimal `json:"price"`
	Rating       int             `json:"rating"`
	PictureURL   string          `json:"picture_url"`
	CategoryId   int             `json:"category_id"`
	CategoryName string          `json:"category_name"`
	ShopName     string          `json:"shop_name"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    time.Time       `json:"deleted_at"`
}

type ProductListParam struct {
	CategoryId string
	AccountID  int
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
