package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type GetSellerProfileRequest struct {
	ShopName string
}

type GetSellerProfileResponse struct {
	SellerName          string              `json:"seller_name"`
	SellerPictureUrl    string              `json:"seller_picture_url"`
	SellerDistrict      string              `json:"seller_district"`
	SellerOperatingHour SellerOperatingHour `json:"seller_operating_hour"`
	ShopNameSlug        string              `json:"shop_name_slug"`
	SellerStars         string              `json:"seller_stars"`
}

type GetSellerProductsRequest struct {
	ShopName string
}

type GetSellerProductsResponse struct {
	SellerProducts []SellerProduct `json:"seller_products"`
}

type GetSellerShowcasesRequest struct {
	ShopName string
}

type GetSellerShowcasesResponse struct {
	Showcases []SellerShowcase
}

type GetSellerShowcaseProductRequest struct {
	ShopName   string
	ShowcaseId string
	Page       int
	Limit      int
}

type RemoveProduct struct {
	ID        int
	SellerID  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type GetSellerShowcaseProductResponse struct {
	SellerProducts []SellerProduct `json:"seller_products"`
}

type SellerOperatingHour struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type SellerProduct struct {
	Name            string          `json:"name"`
	Price           decimal.Decimal `json:"price"`
	PictureUrl      string          `json:"picture_url"`
	Stars           decimal.Decimal `json:"stars"`
	TotalSold       int             `json:"total_sold"`
	CreatedAt       string          `json:"created_at"`
	Category        string          `json:"category"`
	ShopName        string          `json:"shop_name"`
	ProductNameSlug string          `json:"product_name_slug"`
	ShopNameSlug    string          `json:"shop_name_slug"`
}

type SellerShowcase struct {
	ShowcaseId   int    `json:"showcase_id"`
	ShowcaseName string `json:"showcase_name"`
}
