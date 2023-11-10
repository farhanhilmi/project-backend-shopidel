package dtousecase

import "github.com/shopspring/decimal"

type GetSellerProfileRequest struct {
	SellerId int
}

type GetSellerProfileResponse struct {
	SellerId                 int                      `json:"seller_id"`
	SellerName               string                   `json:"seller_name"`
	SellerPictureUrl         string                   `json:"seller_picture_url"`
	SellerDistrict           string                   `json:"seller_district"`
	SellerOperatingHour      SellerOperatingHour      `json:"seller_operating_hour"`
	SellerProducts           []SellerProduct          `json:"seller_products"`
	SellerSelectedCategories []SellerSelectedCategory `json:"seller_selected_categories"`
}

type SellerOperatingHour struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type SellerProduct struct {
	Name           string          `json:"name"`
	Price          decimal.Decimal `json:"price"`
	PictureUrl     string          `json:"picture_url"`
	Stars          decimal.Decimal `json:"stars"`
	TotalSold      int             `json:"total_sold"`
	CreatedAt      string          `json:"created_at"`
	CategoryLevel1 string          `json:"category_level_1"`
	CategoryLevel2 string          `json:"category_level_2"`
	CategoryLevel3 string          `json:"category_level_3"`
}

type SellerSelectedCategory struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
