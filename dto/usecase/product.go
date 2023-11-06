package dtousecase

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	ProductID  int
	Name       string
	CategoryID int
	IsActive   bool
}

type ProductResponse struct {
	ID                int
	Name              string
	Description       string
	CategoryID        int
	Category          model.Category
	HazardousMaterial bool
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsNew             bool
	InternalSKU       string
	ViewCount         int
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

type GetProductDetailRequest struct {
	ProductId  int
	Variant1Id int
	Variant2Id int
}

type GetProductDetailResponse struct {
	Id          int              `json:"id"`
	ProductName string           `json:"name"`
	Description string           `json:"description"`
	Stars       decimal.Decimal  `json:"stars"`
	Sold        int              `json:"sold"`
	Available   int              `json:"available"`
	Images      []string         `json:"images"`
	Variants    []ProductVariant `json:"variants,omitempty"`
}

type ProductVariant struct {
	VariantId   int                `json:"variant_id"`
	VariantName string             `json:"variant_name"`
	Selections  []ProductSelection `json:"selections,omitempty"`
	Stock       int                `json:"stock"`
	Price       decimal.Decimal    `json:"price"`
}

type ProductSelection struct {
	SelectionVariantName string `json:"selection_variant_name,omitempty"`
	SelectionName        string `json:"selection_name,omitempty"`
}
