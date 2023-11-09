package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ProductVariantSelectionCombinations = []*model.ProductVariantSelectionCombinations{
	{
		ProductID:                  1,
		ProductVariantSelectionID1: 1,
		Price:                      decimal.NewFromInt(5000000),
		Stock:                      10,
		PictureURL:                 "https://down-id.img.susercontent.com/file/68171f9daf6be781832415086d2c18e2",
	},
	{
		ProductID:                  2,
		ProductVariantSelectionID1: 2,
		ProductVariantSelectionID2: 4,
		Price:                      decimal.NewFromInt(2000000),
		Stock:                      2,
		PictureURL:                 "https://down-id.img.susercontent.com/file/id-11134207-7r98p-lnb0pqj257k6b4",
	},
	{
		ProductID:                  2,
		ProductVariantSelectionID1: 3,
		ProductVariantSelectionID2: 4,
		Price:                      decimal.NewFromInt(2500000),
		Stock:                      5,
		PictureURL:                 "https://down-id.img.susercontent.com/file/68171f9daf6be781832415086d2c18e2",
	},
	{
		ProductID:                  2,
		ProductVariantSelectionID1: 2,
		ProductVariantSelectionID2: 5,
		Price:                      decimal.NewFromInt(2500000),
		Stock:                      5,
		PictureURL:                 "https://down-id.img.susercontent.com/file/68171f9daf6be781832415086d2c18e2",
	},
	{
		ProductID:                  2,
		ProductVariantSelectionID1: 3,
		ProductVariantSelectionID2: 5,
		Price:                      decimal.NewFromInt(2500000),
		Stock:                      5,
		PictureURL:                 "https://down-id.img.susercontent.com/file/68171f9daf6be781832415086d2c18e2",
	},
	{
		ProductID:                  3,
		ProductVariantSelectionID1: 6,
		Price:                      decimal.NewFromInt(5000000),
		Stock:                      10,
		PictureURL:                 "https://down-id.img.susercontent.com/file/68171f9daf6be781832415086d2c18e2",
	},
	{
		ProductID:                  3,
		ProductVariantSelectionID1: 7,
		Price:                      decimal.NewFromInt(5000000),
		Stock:                      10,
		PictureURL:                 "https://down-id.img.susercontent.com/file/68171f9daf6be781832415086d2c18e2",
	},
}
