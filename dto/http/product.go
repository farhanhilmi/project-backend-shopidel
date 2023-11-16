package dtohttp

import (
	"time"

	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID                int              `json:"id"`
	Name              string           `json:"name,omitempty"`
	Description       string           `json:"description,omitempty"`
	CategoryID        int              `json:"category_id,omitempty"`
	Category          CategoryResponse `json:"category,omitempty"`
	HazardousMaterial bool             `json:"hazardous_material,omitempty"`
	Weight            decimal.Decimal  `json:"wight,omitempty"`
	Size              decimal.Decimal  `json:"size,omitempty"`
	IsNew             bool             `json:"is_new,omitempty"`
	InternalSKU       string           `json:"internal_sku,omitempty"`
	ViewCount         int              `json:"view_count,omitempty"`
	IsActive          bool             `json:"is_active,omitempty"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	UpdatedAt         time.Time        `json:"updated_at,omitempty"`
	DeletedAt         time.Time        `json:"deleted_at,omitempty"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type CheckDeliveryFeeRequest struct {
	SellerID    int    `json:"seller_id" binding:"required"`
	Destination string `json:"destination_address_id" binding:"required"`
	Weight      string `json:"weight" binding:"required"`
	CourierID   int    `json:"courier_id" binding:"required"`
}

type UpdateCartRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type AddNewProductRequest struct {
	ProductName       string                            `json:"product_name" binding:"required"`
	Description       string                            `json:"description"`
	CategoryID        int                               `json:"category_id" binding:"required"`
	HazardousMaterial *bool                             `json:"hazardous_material" binding:"required"`
	IsNew             *bool                             `json:"is_new" binding:"required"`
	InternalSKU       string                            `json:"internal_sku" binding:"required"`
	Weight            decimal.Decimal                   `json:"weight" binding:"required"`
	Size              decimal.Decimal                   `json:"size" binding:"required"`
	IsActive          *bool                             `json:"is_active" binding:"required"`
	Variants          []dtousecase.AddNewProductVariant `json:"variants"`
}

// type AddNewProductVariant struct {
// 	VariantName  string          `json:"variant_name" binding:"required"`
// 	VariantValue string          `json:"variant_value" binding:"required"`
// 	Stock        int             `json:"stock" binding:"required"`
// 	Price        decimal.Decimal `json:"price" binding:"required"`
// 	// Images       *multipart.FileHeader
// }
