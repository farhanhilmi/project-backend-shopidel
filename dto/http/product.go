package dtohttp

import (
	"time"

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

type ProductListResponse struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	District   string          `json:"district"`
	TotalSold  int             `json:"total_sold"`
	Price      decimal.Decimal `json:"price"`
	PictureURL string          `json:"picture_url"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"-"`
	DeletedAt  time.Time       `json:"-"`
}