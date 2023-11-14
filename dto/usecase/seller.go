package dtousecase

import "github.com/shopspring/decimal"

type GetSellerProfileRequest struct {
	ShopName string
}

type GetSellerProfileResponse struct {
	SellerName          string              `json:"seller_name"`
	SellerPictureUrl    string              `json:"seller_picture_url"`
	SellerDistrict      string              `json:"seller_district"`
	SellerOperatingHour SellerOperatingHour `json:"seller_operating_hour"`
}

type GetSellerProductsRequest struct {
	ShopName string
}

type GetSellerProductsResponse struct {
	SellerProducts []SellerProduct `json:"seller_products"`
}

type GetSellerCategoriesRequest struct {
	ShopName string
}

type GetSellerCategoriesResponse struct {
	Categories []SellerCategory
}

type GetSellerCategoryProductRequest struct {
	ShopName   string
	CategoryId string
}

type GetSellerCategoryProductResponse struct {
	SellerProducts []SellerProduct `json:"seller_products"`
}

type SellerOperatingHour struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type SellerProduct struct {
	Name       string          `json:"name"`
	Price      decimal.Decimal `json:"price"`
	PictureUrl string          `json:"picture_url"`
	Stars      decimal.Decimal `json:"stars"`
	TotalSold  int             `json:"total_sold"`
	CreatedAt  string          `json:"created_at"`
	Category   string          `json:"category"`
	ShopName   string          `json:"shop_name"`
}

type SellerCategory struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
