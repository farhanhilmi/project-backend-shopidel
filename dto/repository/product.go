package dtorepository

import "github.com/shopspring/decimal"

type ProductRequest struct {
	ProductID int
}

type ProductResponse struct {
	ID          int
	Name        string
	Description string
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
