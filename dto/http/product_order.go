package dtohttp

import (
	"time"

	"github.com/shopspring/decimal"
)

type CanceledOrderRequest struct {
	ProductID int    `json:"product_id"`
	AccountID int    `json:"account_id"`
	Notes     string `json:"notes" binding:"required"`
}

type ProductOrderResponse struct {
	ID            int             `json:"id,omitempty"`
	CourierID     int             `json:"courier_id,omitempty"`
	AccountID     int             `json:"account_id,omitempty"`
	DeliveryFee   decimal.Decimal `json:"delivery_id,omitempty"`
	Province      string          `json:"province,omitempty"`
	District      string          `json:"district,omitempty"`
	SubDistrict   string          `json:"sub_district,omitempty"`
	Kelurahan     string          `json:"kelurahan,omitempty"`
	ZipCode       string          `json:"zip_code,omitempty"`
	AddressDetail string          `json:"address_detail,omitempty"`
	Notes         string          `json:"notes,omitempty"`
	Status        string          `json:"status,omitempty"`
	CreatedAt     time.Time       `json:"created_at,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty"`
	DeletedAt     time.Time       `json:"deleted_at,omitempty"`
}

type ProductOrderReceiveResponse struct {
	ID     int    `json:"id,omitempty"`
	Notes  string `json:"notes,omitempty"`
	Status string `json:"status,omitempty"`
}
