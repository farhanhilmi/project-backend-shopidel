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
	Origin      string `json:"origin" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	Weight      string `json:"weight" binding:"required"`
	Courier     string `json:"courier" binding:"required"`
}
