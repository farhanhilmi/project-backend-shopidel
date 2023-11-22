package dtorepository

import "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"

type FindShowcaseRequest struct {
	ShopId int
	Page   int
}

type FindShowcaseResponse struct {
	ShopPromotions []model.SellerShowcase
	CurrentPage    int
	TotalPages     int
	TotalItems     int
	Limit          int
}
