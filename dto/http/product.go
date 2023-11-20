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
	ProductName       string                            `form:"product_name" binding:"required"`
	Description       string                            `form:"description"`
	CategoryID        int                               `form:"category_id" binding:"required"`
	HazardousMaterial *bool                             `form:"hazardous_material" binding:"required"`
	IsNew             *bool                             `form:"is_new" binding:"required"`
	InternalSKU       string                            `form:"internal_sku"`
	Weight            decimal.Decimal                   `form:"weight" binding:"required"`
	Size              decimal.Decimal                   `form:"size" binding:"required"`
	IsActive          *bool                             `form:"is_active" binding:"required"`
	Variants          []dtousecase.AddNewProductVariant `form:"variants[]" binding:"required"`
	VideoURL          string                            `form:"video_url"`
}

type UploadNewPhoto struct {
	ImageID string `form:"image_id" binding:"required"`
}
