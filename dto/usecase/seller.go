package dtousecase

import (
	"time"

	"github.com/shopspring/decimal"
)

type GetSellerProfileRequest struct {
	ShopId   int
	ShopName string
}

type GetSellerProfileResponse struct {
	SellerName          string
	SellerPictureUrl    string
	SellerDistrict      string
	SellerOperatingHour SellerOperatingHour
	ShopNameSlug        string
	SellerStars         decimal.Decimal
	SellerDescription   string
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

type WithdrawBalance struct {
	OrderID  int             `json:"order_id"`
	SellerID int             `json:"-"`
	Balance  decimal.Decimal `json:"balance"`
}

type GetSellerShowcaseProductResponse struct {
	SellerProducts []SellerProduct `json:"seller_products"`
	Limit          int
	CurrentPage    int
	TotalItem      int
	TotalPage      int
}

type SellerOperatingHour struct {
	Start time.Time
	End   time.Time
}

type GetSellerShowcaseProducts struct {
	SellerProducts []SellerProduct
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

type UpdateShopProfileRequest struct {
	ShopId          int
	ShopName        string
	ShopDescription string
	OpeningHours    time.Time
	ClosingHours    time.Time
}

type UpdateShopProfileResponse struct {
	ShopId          int
	ShopName        string
	ShopDescription string
	OpeningHours    time.Time
	ClosingHours    time.Time
}
