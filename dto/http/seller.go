package dtohttp

import "github.com/shopspring/decimal"

type UpdateSellerProfileBody struct {
	ShopName        string `json:"shop_name" binding:"required"`
	ShopDescription string `json:"shop_description" binding:"required"`
	OpeningHours    string `json:"opening_hours" binding:"required"`
	ClosingHours    string `json:"closing_hours" binding:"required"`
}

type UpdateSellerProfileResponse struct {
	ShopName        string `json:"shop_name" binding:"required"`
	ShopDescription string `json:"shop_description" binding:"required"`
	OpeningHours    string `json:"opening_hours" binding:"required"`
	ClosingHours    string `json:"closing_hours" binding:"required"`
}

type GetSellerProfileResponse struct {
	SellerName          string              `json:"seller_name"`
	SellerPictureUrl    string              `json:"seller_picture_url"`
	SellerDistrict      string              `json:"seller_district"`
	SellerOperatingHour SellerOperatingHour `json:"seller_operating_hour"`
	ShopNameSlug        string              `json:"shop_name_slug"`
	SellerStars         decimal.Decimal     `json:"seller_stars"`
	SellerDescription   string              `json:"seller_description"`
}

type SellerOperatingHour struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
