package seeder

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var ProductOrders = []model.ProductOrders{
	{
		ID:            1,
		CourierID:     1,
		AccountID:     2,
		DeliveryFee:   decimal.NewFromInt(15000),
		Province:      "Jawa Barat",
		District:      "Kabupaten Bandung",
		SubDistrict:   "Bojongsoang",
		Kelurahan:     "Sukapura",
		ZipCode:       "40851",
		AddressDetail: "Jl Telekomunikasi No 1 Bojongsoang",
		Status:        constant.StatusOrderOnProcess,
	},
	{
		ID:            2,
		CourierID:     1,
		AccountID:     2,
		DeliveryFee:   decimal.NewFromInt(20000),
		Province:      "Jawa Barat",
		District:      "Kabupaten Bandung",
		SubDistrict:   "Bojongsoang",
		Kelurahan:     "Sukapura",
		ZipCode:       "40851",
		AddressDetail: "Jl Telekomunikasi No 1 Bojongsoang",
		Status:        constant.StatusOrderDelivered,
	},
	{
		ID:            3,
		CourierID:     1,
		AccountID:     2,
		DeliveryFee:   decimal.NewFromInt(17000),
		Province:      "Jawa Barat",
		District:      "Kabupaten Bandung",
		SubDistrict:   "Bojongsoang",
		Kelurahan:     "Sukapura",
		ZipCode:       "40851",
		AddressDetail: "Jl Telekomunikasi No 1 Bojongsoang",
		Status:        constant.StatusCanceled,
	},
	{
		ID:            4,
		CourierID:     1,
		AccountID:     2,
		DeliveryFee:   decimal.NewFromInt(20000),
		Province:      "Jawa Barat",
		District:      "Kabupaten Bandung",
		SubDistrict:   "Bojongsoang",
		Kelurahan:     "Sukapura",
		ZipCode:       "40851",
		AddressDetail: "Jl Telekomunikasi No 1 Bojongsoang",
		Status:        constant.StatusOrderCompleted,
	},
}
