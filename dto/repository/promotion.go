package dtorepository

import "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"

type FindShopPromotionsRequest struct {
	ShopId int
	Page   int
}

type FindShopPromotionsResponse struct {
	ShopPromotions []model.ShopPromotion
	CurrentPage    int
	TotalPages     int
	TotalItems     int
	Limit          int
}
