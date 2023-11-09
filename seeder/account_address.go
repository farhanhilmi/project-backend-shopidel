package seeder

import "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"

var AccountAddress = []model.AccountAddress{
	{
		AccountID:            2,
		Province:             "DKI Jakarta",
		District:             "Jakarta Selatan",
		SubDistrict:          "Setiabudi",
		Kelurahan:            "Setiabudi",
		ZipCode:              "12230",
		RajaOngkirDistrictId: "153",
		Detail:               "Sopo Del Tower, Jalan Mega Kuningan Barat III Lot 10.1-6, RT.03/RW.03",
		IsBuyerDefault:       true,
		IsSellerDefault:      false,
	},
	{
		AccountID:            2,
		Province:             "DKI Jakarta",
		District:             "Jakarta Timur",
		SubDistrict:          "Jatinegara",
		Kelurahan:            "Cipinang Besar Sel",
		ZipCode:              "14738",
		RajaOngkirDistrictId: "154",
		Detail:               "Jl. Jend. Basuki Rachmat No.1A",
	},
	{
		AccountID:            2,
		Province:             "DKI Jakarta",
		District:             "Jakarta Barat",
		SubDistrict:          "Kembangan",
		Kelurahan:            "Lingkar Luar",
		RajaOngkirDistrictId: "151",
		ZipCode:              "11610",
		Detail:               "Puri Mansion Estate, Jl. Puri",
		IsBuyerDefault:       false,
		IsSellerDefault:      true,
	},
	{
		AccountID:            1,
		Province:             "Jawa Barat",
		District:             "Kabupaten Bandung",
		SubDistrict:          "Bojongsoang",
		Kelurahan:            "Sukapura",
		ZipCode:              "40851",
		Detail:               "Jl Telekomunikasi No 1 Bojongsoang",
		RajaOngkirDistrictId: "10",
		IsBuyerDefault:       false,
		IsSellerDefault:      true,
	},
}
