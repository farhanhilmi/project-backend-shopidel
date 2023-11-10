package dtorepository

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	ProductID int
	AccountId int
}

type ProductResponse struct {
	ID          int
	Name        string
	Description string
	IsFavorite  bool
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

type IsProductFavoriteRequest struct {
	AccountId int
	ProductId int
}

type IsProductFavoriteResponse struct {
	AccountId  int
	ProductId  int
	IsFavorite bool
}
