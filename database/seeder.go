package database

import (
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

func RunSeeder() {
	if config.GetEnv("ENV") == "dev" {
		dropTable()
		createTable()
		seeding()
	}
}

func dropTable() {
	sql := `
		drop table if exists 
			accounts, 
			categories,
			my_wallet_transaction_histories,
			product_images,
			product_variant_selection_combinations,
			product_variant_selections, 
			product_variants, 
			product_videos,
			products,
			used_emails,
			MyWalletTransactionHistories,
			product_orders,
			sale_wallet_transaction_histories,
			product_order_details,
			couriers,
			account_addresses,
			account_carts,
			seller_couriers,
			provinces,
			districts
			;
	`

	err := db.Exec(sql).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully delete tables")
}

func createTable() {
	err := db.AutoMigrate(
		&model.Accounts{},
		&model.UsedEmail{},
		&model.Category{},
		&model.Products{},
		&model.ProductImages{},
		&model.ProductVideos{},
		&model.ProductVariants{},
		&model.ProductVariantSelections{},
		&model.ProductVariantSelectionCombinations{},
		&model.MyWalletTransactionHistories{},
		&model.SaleWalletTransactionHistories{},
		&model.Couriers{},
		&model.Category{},
		&model.Products{},
		&model.ProductImages{},
		&model.ProductVideos{},
		&model.ProductVariants{},
		&model.ProductVariantSelections{},
		&model.ProductVariantSelectionCombinations{},
		&model.ProductOrders{},
		&model.ProductOrderDetails{},
		&model.AccountAddress{},
		&model.AccountCarts{},
		&model.SellerCouriers{},
		&model.Province{},
		&model.District{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully migrate tables")
}

func seeding() {
	provinces := []*model.Province{
		{
			Name:                 "Bali",
			RajaOngkirProvinceId: 1,
		},
		{
			Name:                 "Bangka Belitung",
			RajaOngkirProvinceId: 2,
		},
		{
			Name:                 "Banten",
			RajaOngkirProvinceId: 3,
		},
		{
			Name:                 "Bengkulu",
			RajaOngkirProvinceId: 4,
		},
		{
			Name:                 "DI Yogyakarta",
			RajaOngkirProvinceId: 5,
		},
		{
			Name:                 "DKI Jakarta",
			RajaOngkirProvinceId: 6,
		},
		{
			Name:                 "Gorontalo",
			RajaOngkirProvinceId: 7,
		},
		{
			Name:                 "Jambi",
			RajaOngkirProvinceId: 8,
		},
		{
			Name:                 "Jawa Barat",
			RajaOngkirProvinceId: 9,
		},
		{
			Name:                 "Jawa Tengah",
			RajaOngkirProvinceId: 10,
		},
		{
			Name:                 "Jawa Timur",
			RajaOngkirProvinceId: 11,
		},
		{
			Name:                 "Kalimantan Barat",
			RajaOngkirProvinceId: 12,
		},
		{
			Name:                 "Kalimantan Selatan",
			RajaOngkirProvinceId: 13,
		},
		{
			Name:                 "Kalimantan Tengah",
			RajaOngkirProvinceId: 14,
		},
		{
			Name:                 "Kalimantan Timur",
			RajaOngkirProvinceId: 15,
		},
		{
			Name:                 "Kalimantan Utara",
			RajaOngkirProvinceId: 16,
		},
		{
			Name:                 "Kepulauan Riau",
			RajaOngkirProvinceId: 17,
		},
		{
			Name:                 "Lampung",
			RajaOngkirProvinceId: 18,
		},
		{
			Name:                 "Maluku",
			RajaOngkirProvinceId: 19,
		},
		{
			Name:                 "Maluku Utara",
			RajaOngkirProvinceId: 20,
		},
		{
			Name:                 "Nanggroe Aceh Darussalam (NAD)",
			RajaOngkirProvinceId: 21,
		},
		{
			Name:                 "Nusa Tenggara Barat (NTB)",
			RajaOngkirProvinceId: 22,
		},
		{
			Name:                 "Nusa Tenggara Timur (NTT)",
			RajaOngkirProvinceId: 23,
		},
		{
			Name:                 "Papua",
			RajaOngkirProvinceId: 24,
		},
		{
			Name:                 "Papua Barat",
			RajaOngkirProvinceId: 25,
		},
		{
			Name:                 "Riau",
			RajaOngkirProvinceId: 26,
		},
		{
			Name:                 "Sulawesi Barat",
			RajaOngkirProvinceId: 27,
		},
		{
			Name:                 "Sulawesi Selatan",
			RajaOngkirProvinceId: 28,
		},
		{
			Name:                 "Sulawesi Tengah",
			RajaOngkirProvinceId: 29,
		},
		{
			Name:                 "Sulawesi Tenggara",
			RajaOngkirProvinceId: 30,
		},
		{
			Name:                 "Sulawesi Utara",
			RajaOngkirProvinceId: 31,
		},
		{
			Name:                 "Sumatera Barat",
			RajaOngkirProvinceId: 32,
		},
		{
			Name:                 "Sumatera Selatan",
			RajaOngkirProvinceId: 33,
		},
		{
			Name:                 "Sumatera Utara",
			RajaOngkirProvinceId: 34,
		},
	}

	if err := db.Create(provinces).Error; err != nil {
		panic(err)

	}

	districts := []*model.District{
		{
			Name:                 "Kabupaten Aceh Barat",
			ProvinceId:           21,
			RajaOngkirDistrictId: "1",
		},
		{
			Name:                 "Kabupaten Aceh Barat Daya",
			ProvinceId:           21,
			RajaOngkirDistrictId: "2",
		},
		{
			Name:                 "Kabupaten Aceh Besar",
			ProvinceId:           21,
			RajaOngkirDistrictId: "3",
		},
		{
			Name:                 "Kabupaten Aceh Jaya",
			ProvinceId:           21,
			RajaOngkirDistrictId: "4",
		},
		{
			Name:                 "Kabupaten Aceh Selatan",
			ProvinceId:           21,
			RajaOngkirDistrictId: "5",
		},
		{
			Name:                 "Kabupaten Aceh Singkil",
			ProvinceId:           21,
			RajaOngkirDistrictId: "6",
		},
		{
			Name:                 "Kabupaten Aceh Tamiang",
			ProvinceId:           21,
			RajaOngkirDistrictId: "7",
		},
		{
			Name:                 "Kabupaten Aceh Tengah",
			ProvinceId:           21,
			RajaOngkirDistrictId: "8",
		},
		{
			Name:                 "Kabupaten Aceh Tenggara",
			ProvinceId:           21,
			RajaOngkirDistrictId: "9",
		},
		{
			Name:                 "Kabupaten Aceh Timur",
			ProvinceId:           21,
			RajaOngkirDistrictId: "10",
		},
		{
			Name:                 "Kabupaten Aceh Utara",
			ProvinceId:           21,
			RajaOngkirDistrictId: "11",
		},
		{
			Name:                 "Kabupaten Agam",
			ProvinceId:           32,
			RajaOngkirDistrictId: "12",
		},
		{
			Name:                 "Kabupaten Alor",
			ProvinceId:           23,
			RajaOngkirDistrictId: "13",
		},
		{
			Name:                 "Kota Ambon",
			ProvinceId:           19,
			RajaOngkirDistrictId: "14",
		},
		{
			Name:                 "Kabupaten Asahan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "15",
		},
		{
			Name:                 "Kabupaten Asmat",
			ProvinceId:           24,
			RajaOngkirDistrictId: "16",
		},
		{
			Name:                 "Kabupaten Badung",
			ProvinceId:           1,
			RajaOngkirDistrictId: "17",
		},
		{
			Name:                 "Kabupaten Balangan",
			ProvinceId:           13,
			RajaOngkirDistrictId: "18",
		},
		{
			Name:                 "Kota Balikpapan",
			ProvinceId:           15,
			RajaOngkirDistrictId: "19",
		},
		{
			Name:                 "Kota Banda Aceh",
			ProvinceId:           21,
			RajaOngkirDistrictId: "20",
		},
		{
			Name:                 "Kota Bandar Lampung",
			ProvinceId:           18,
			RajaOngkirDistrictId: "21",
		},
		{
			Name:                 "Kabupaten Bandung",
			ProvinceId:           9,
			RajaOngkirDistrictId: "22",
		},
		{
			Name:                 "Kota Bandung",
			ProvinceId:           9,
			RajaOngkirDistrictId: "23",
		},
		{
			Name:                 "Kabupaten Bandung Barat",
			ProvinceId:           9,
			RajaOngkirDistrictId: "24",
		},
		{
			Name:                 "Kabupaten Banggai",
			ProvinceId:           29,
			RajaOngkirDistrictId: "25",
		},
		{
			Name:                 "Kabupaten Banggai Kepulauan",
			ProvinceId:           29,
			RajaOngkirDistrictId: "26",
		},
		{
			Name:                 "Kabupaten Bangka",
			ProvinceId:           2,
			RajaOngkirDistrictId: "27",
		},
		{
			Name:                 "Kabupaten Bangka Barat",
			ProvinceId:           2,
			RajaOngkirDistrictId: "28",
		},
		{
			Name:                 "Kabupaten Bangka Selatan",
			ProvinceId:           2,
			RajaOngkirDistrictId: "29",
		},
		{
			Name:                 "Kabupaten Bangka Tengah",
			ProvinceId:           2,
			RajaOngkirDistrictId: "30",
		},
		{
			Name:                 "Kabupaten Bangkalan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "31",
		},
		{
			Name:                 "Kabupaten Bangli",
			ProvinceId:           1,
			RajaOngkirDistrictId: "32",
		},
		{
			Name:                 "Kabupaten Banjar",
			ProvinceId:           13,
			RajaOngkirDistrictId: "33",
		},
		{
			Name:                 "Kota Banjar",
			ProvinceId:           9,
			RajaOngkirDistrictId: "34",
		},
		{
			Name:                 "Kota Banjarbaru",
			ProvinceId:           13,
			RajaOngkirDistrictId: "35",
		},
		{
			Name:                 "Kota Banjarmasin",
			ProvinceId:           13,
			RajaOngkirDistrictId: "36",
		},
		{
			Name:                 "Kabupaten Banjarnegara",
			ProvinceId:           10,
			RajaOngkirDistrictId: "37",
		},
		{
			Name:                 "Kabupaten Bantaeng",
			ProvinceId:           28,
			RajaOngkirDistrictId: "38",
		},
		{
			Name:                 "Kabupaten Bantul",
			ProvinceId:           5,
			RajaOngkirDistrictId: "39",
		},
		{
			Name:                 "Kabupaten Banyuasin",
			ProvinceId:           33,
			RajaOngkirDistrictId: "40",
		},
		{
			Name:                 "Kabupaten Banyumas",
			ProvinceId:           10,
			RajaOngkirDistrictId: "41",
		},
		{
			Name:                 "Kabupaten Banyuwangi",
			ProvinceId:           11,
			RajaOngkirDistrictId: "42",
		},
		{
			Name:                 "Kabupaten Barito Kuala",
			ProvinceId:           13,
			RajaOngkirDistrictId: "43",
		},
		{
			Name:                 "Kabupaten Barito Selatan",
			ProvinceId:           14,
			RajaOngkirDistrictId: "44",
		},
		{
			Name:                 "Kabupaten Barito Timur",
			ProvinceId:           14,
			RajaOngkirDistrictId: "45",
		},
		{
			Name:                 "Kabupaten Barito Utara",
			ProvinceId:           14,
			RajaOngkirDistrictId: "46",
		},
		{
			Name:                 "Kabupaten Barru",
			ProvinceId:           28,
			RajaOngkirDistrictId: "47",
		},
		{
			Name:                 "Kota Batam",
			ProvinceId:           17,
			RajaOngkirDistrictId: "48",
		},
		{
			Name:                 "Kabupaten Batang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "49",
		},
		{
			Name:                 "Kabupaten Batang Hari",
			ProvinceId:           8,
			RajaOngkirDistrictId: "50",
		},
		{
			Name:                 "Kota Batu",
			ProvinceId:           11,
			RajaOngkirDistrictId: "51",
		},
		{
			Name:                 "Kabupaten Batu Bara",
			ProvinceId:           34,
			RajaOngkirDistrictId: "52",
		},
		{
			Name:                 "Kota Bau-Bau",
			ProvinceId:           30,
			RajaOngkirDistrictId: "53",
		},
		{
			Name:                 "Kabupaten Bekasi",
			ProvinceId:           9,
			RajaOngkirDistrictId: "54",
		},
		{
			Name:                 "Kota Bekasi",
			ProvinceId:           9,
			RajaOngkirDistrictId: "55",
		},
		{
			Name:                 "Kabupaten Belitung",
			ProvinceId:           2,
			RajaOngkirDistrictId: "56",
		},
		{
			Name:                 "Kabupaten Belitung Timur",
			ProvinceId:           2,
			RajaOngkirDistrictId: "57",
		},
		{
			Name:                 "Kabupaten Belu",
			ProvinceId:           23,
			RajaOngkirDistrictId: "58",
		},
		{
			Name:                 "Kabupaten Bener Meriah",
			ProvinceId:           21,
			RajaOngkirDistrictId: "59",
		},
		{
			Name:                 "Kabupaten Bengkalis",
			ProvinceId:           26,
			RajaOngkirDistrictId: "60",
		},
		{
			Name:                 "Kabupaten Bengkayang",
			ProvinceId:           12,
			RajaOngkirDistrictId: "61",
		},
		{
			Name:                 "Kota Bengkulu",
			ProvinceId:           4,
			RajaOngkirDistrictId: "62",
		},
		{
			Name:                 "Kabupaten Bengkulu Selatan",
			ProvinceId:           4,
			RajaOngkirDistrictId: "63",
		},
		{
			Name:                 "Kabupaten Bengkulu Tengah",
			ProvinceId:           4,
			RajaOngkirDistrictId: "64",
		},
		{
			Name:                 "Kabupaten Bengkulu Utara",
			ProvinceId:           4,
			RajaOngkirDistrictId: "65",
		},
		{
			Name:                 "Kabupaten Berau",
			ProvinceId:           15,
			RajaOngkirDistrictId: "66",
		},
		{
			Name:                 "Kabupaten Biak Numfor",
			ProvinceId:           24,
			RajaOngkirDistrictId: "67",
		},
		{
			Name:                 "Kabupaten Bima",
			ProvinceId:           22,
			RajaOngkirDistrictId: "68",
		},
		{
			Name:                 "Kota Bima",
			ProvinceId:           22,
			RajaOngkirDistrictId: "69",
		},
		{
			Name:                 "Kota Binjai",
			ProvinceId:           34,
			RajaOngkirDistrictId: "70",
		},
		{
			Name:                 "Kabupaten Bintan",
			ProvinceId:           17,
			RajaOngkirDistrictId: "71",
		},
		{
			Name:                 "Kabupaten Bireuen",
			ProvinceId:           21,
			RajaOngkirDistrictId: "72",
		},
		{
			Name:                 "Kota Bitung",
			ProvinceId:           31,
			RajaOngkirDistrictId: "73",
		},
		{
			Name:                 "Kabupaten Blitar",
			ProvinceId:           11,
			RajaOngkirDistrictId: "74",
		},
		{
			Name:                 "Kota Blitar",
			ProvinceId:           11,
			RajaOngkirDistrictId: "75",
		},
		{
			Name:                 "Kabupaten Blora",
			ProvinceId:           10,
			RajaOngkirDistrictId: "76",
		},
		{
			Name:                 "Kabupaten Boalemo",
			ProvinceId:           7,
			RajaOngkirDistrictId: "77",
		},
		{
			Name:                 "Kabupaten Bogor",
			ProvinceId:           9,
			RajaOngkirDistrictId: "78",
		},
		{
			Name:                 "Kota Bogor",
			ProvinceId:           9,
			RajaOngkirDistrictId: "79",
		},
		{
			Name:                 "Kabupaten Bojonegoro",
			ProvinceId:           11,
			RajaOngkirDistrictId: "80",
		},
		{
			Name:                 "Kabupaten Bolaang Mongondow (Bolmong)",
			ProvinceId:           31,
			RajaOngkirDistrictId: "81",
		},
		{
			Name:                 "Kabupaten Bolaang Mongondow Selatan",
			ProvinceId:           31,
			RajaOngkirDistrictId: "82",
		},
		{
			Name:                 "Kabupaten Bolaang Mongondow Timur",
			ProvinceId:           31,
			RajaOngkirDistrictId: "83",
		},
		{
			Name:                 "Kabupaten Bolaang Mongondow Utara",
			ProvinceId:           31,
			RajaOngkirDistrictId: "84",
		},
		{
			Name:                 "Kabupaten Bombana",
			ProvinceId:           30,
			RajaOngkirDistrictId: "85",
		},
		{
			Name:                 "Kabupaten Bondowoso",
			ProvinceId:           11,
			RajaOngkirDistrictId: "86",
		},
		{
			Name:                 "Kabupaten Bone",
			ProvinceId:           28,
			RajaOngkirDistrictId: "87",
		},
		{
			Name:                 "Kabupaten Bone Bolango",
			ProvinceId:           7,
			RajaOngkirDistrictId: "88",
		},
		{
			Name:                 "Kota Bontang",
			ProvinceId:           15,
			RajaOngkirDistrictId: "89",
		},
		{
			Name:                 "Kabupaten Boven Digoel",
			ProvinceId:           24,
			RajaOngkirDistrictId: "90",
		},
		{
			Name:                 "Kabupaten Boyolali",
			ProvinceId:           10,
			RajaOngkirDistrictId: "91",
		},
		{
			Name:                 "Kabupaten Brebes",
			ProvinceId:           10,
			RajaOngkirDistrictId: "92",
		},
		{
			Name:                 "Kota Bukittinggi",
			ProvinceId:           32,
			RajaOngkirDistrictId: "93",
		},
		{
			Name:                 "Kabupaten Buleleng",
			ProvinceId:           1,
			RajaOngkirDistrictId: "94",
		},
		{
			Name:                 "Kabupaten Bulukumba",
			ProvinceId:           28,
			RajaOngkirDistrictId: "95",
		},
		{
			Name:                 "Kabupaten Bulungan (Bulongan)",
			ProvinceId:           16,
			RajaOngkirDistrictId: "96",
		},
		{
			Name:                 "Kabupaten Bungo",
			ProvinceId:           8,
			RajaOngkirDistrictId: "97",
		},
		{
			Name:                 "Kabupaten Buol",
			ProvinceId:           29,
			RajaOngkirDistrictId: "98",
		},
		{
			Name:                 "Kabupaten Buru",
			ProvinceId:           19,
			RajaOngkirDistrictId: "99",
		},
		{
			Name:                 "Kabupaten Buru Selatan",
			ProvinceId:           19,
			RajaOngkirDistrictId: "100",
		},
		{
			Name:                 "Kabupaten Buton",
			ProvinceId:           30,
			RajaOngkirDistrictId: "101",
		},
		{
			Name:                 "Kabupaten Buton Utara",
			ProvinceId:           30,
			RajaOngkirDistrictId: "102",
		},
		{
			Name:                 "Kabupaten Ciamis",
			ProvinceId:           9,
			RajaOngkirDistrictId: "103",
		},
		{
			Name:                 "Kabupaten Cianjur",
			ProvinceId:           9,
			RajaOngkirDistrictId: "104",
		},
		{
			Name:                 "Kabupaten Cilacap",
			ProvinceId:           10,
			RajaOngkirDistrictId: "105",
		},
		{
			Name:                 "Kota Cilegon",
			ProvinceId:           3,
			RajaOngkirDistrictId: "106",
		},
		{
			Name:                 "Kota Cimahi",
			ProvinceId:           9,
			RajaOngkirDistrictId: "107",
		},
		{
			Name:                 "Kabupaten Cirebon",
			ProvinceId:           9,
			RajaOngkirDistrictId: "108",
		},
		{
			Name:                 "Kota Cirebon",
			ProvinceId:           9,
			RajaOngkirDistrictId: "109",
		},
		{
			Name:                 "Kabupaten Dairi",
			ProvinceId:           34,
			RajaOngkirDistrictId: "110",
		},
		{
			Name:                 "Kabupaten Deiyai (Deliyai)",
			ProvinceId:           24,
			RajaOngkirDistrictId: "111",
		},
		{
			Name:                 "Kabupaten Deli Serdang",
			ProvinceId:           34,
			RajaOngkirDistrictId: "112",
		},
		{
			Name:                 "Kabupaten Demak",
			ProvinceId:           10,
			RajaOngkirDistrictId: "113",
		},
		{
			Name:                 "Kota Denpasar",
			ProvinceId:           1,
			RajaOngkirDistrictId: "114",
		},
		{
			Name:                 "Kota Depok",
			ProvinceId:           9,
			RajaOngkirDistrictId: "115",
		},
		{
			Name:                 "Kabupaten Dharmasraya",
			ProvinceId:           32,
			RajaOngkirDistrictId: "116",
		},
		{
			Name:                 "Kabupaten Dogiyai",
			ProvinceId:           24,
			RajaOngkirDistrictId: "117",
		},
		{
			Name:                 "Kabupaten Dompu",
			ProvinceId:           22,
			RajaOngkirDistrictId: "118",
		},
		{
			Name:                 "Kabupaten Donggala",
			ProvinceId:           29,
			RajaOngkirDistrictId: "119",
		},
		{
			Name:                 "Kota Dumai",
			ProvinceId:           26,
			RajaOngkirDistrictId: "120",
		},
		{
			Name:                 "Kabupaten Empat Lawang",
			ProvinceId:           33,
			RajaOngkirDistrictId: "121",
		},
		{
			Name:                 "Kabupaten Ende",
			ProvinceId:           23,
			RajaOngkirDistrictId: "122",
		},
		{
			Name:                 "Kabupaten Enrekang",
			ProvinceId:           28,
			RajaOngkirDistrictId: "123",
		},
		{
			Name:                 "Kabupaten Fakfak",
			ProvinceId:           25,
			RajaOngkirDistrictId: "124",
		},
		{
			Name:                 "Kabupaten Flores Timur",
			ProvinceId:           23,
			RajaOngkirDistrictId: "125",
		},
		{
			Name:                 "Kabupaten Garut",
			ProvinceId:           9,
			RajaOngkirDistrictId: "126",
		},
		{
			Name:                 "Kabupaten Gayo Lues",
			ProvinceId:           21,
			RajaOngkirDistrictId: "127",
		},
		{
			Name:                 "Kabupaten Gianyar",
			ProvinceId:           1,
			RajaOngkirDistrictId: "128",
		},
		{
			Name:                 "Kabupaten Gorontalo",
			ProvinceId:           7,
			RajaOngkirDistrictId: "129",
		},
		{
			Name:                 "Kota Gorontalo",
			ProvinceId:           7,
			RajaOngkirDistrictId: "130",
		},
		{
			Name:                 "Kabupaten Gorontalo Utara",
			ProvinceId:           7,
			RajaOngkirDistrictId: "131",
		},
		{
			Name:                 "Kabupaten Gowa",
			ProvinceId:           28,
			RajaOngkirDistrictId: "132",
		},
		{
			Name:                 "Kabupaten Gresik",
			ProvinceId:           11,
			RajaOngkirDistrictId: "133",
		},
		{
			Name:                 "Kabupaten Grobogan",
			ProvinceId:           10,
			RajaOngkirDistrictId: "134",
		},
		{
			Name:                 "Kabupaten Gunung Kidul",
			ProvinceId:           5,
			RajaOngkirDistrictId: "135",
		},
		{
			Name:                 "Kabupaten Gunung Mas",
			ProvinceId:           14,
			RajaOngkirDistrictId: "136",
		},
		{
			Name:                 "Kota Gunungsitoli",
			ProvinceId:           34,
			RajaOngkirDistrictId: "137",
		},
		{
			Name:                 "Kabupaten Halmahera Barat",
			ProvinceId:           20,
			RajaOngkirDistrictId: "138",
		},
		{
			Name:                 "Kabupaten Halmahera Selatan",
			ProvinceId:           20,
			RajaOngkirDistrictId: "139",
		},
		{
			Name:                 "Kabupaten Halmahera Tengah",
			ProvinceId:           20,
			RajaOngkirDistrictId: "140",
		},
		{
			Name:                 "Kabupaten Halmahera Timur",
			ProvinceId:           20,
			RajaOngkirDistrictId: "141",
		},
		{
			Name:                 "Kabupaten Halmahera Utara",
			ProvinceId:           20,
			RajaOngkirDistrictId: "142",
		},
		{
			Name:                 "Kabupaten Hulu Sungai Selatan",
			ProvinceId:           13,
			RajaOngkirDistrictId: "143",
		},
		{
			Name:                 "Kabupaten Hulu Sungai Tengah",
			ProvinceId:           13,
			RajaOngkirDistrictId: "144",
		},
		{
			Name:                 "Kabupaten Hulu Sungai Utara",
			ProvinceId:           13,
			RajaOngkirDistrictId: "145",
		},
		{
			Name:                 "Kabupaten Humbang Hasundutan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "146",
		},
		{
			Name:                 "Kabupaten Indragiri Hilir",
			ProvinceId:           26,
			RajaOngkirDistrictId: "147",
		},
		{
			Name:                 "Kabupaten Indragiri Hulu",
			ProvinceId:           26,
			RajaOngkirDistrictId: "148",
		},
		{
			Name:                 "Kabupaten Indramayu",
			ProvinceId:           9,
			RajaOngkirDistrictId: "149",
		},
		{
			Name:                 "Kabupaten Intan Jaya",
			ProvinceId:           24,
			RajaOngkirDistrictId: "150",
		},
		{
			Name:                 "Kota Jakarta Barat",
			ProvinceId:           6,
			RajaOngkirDistrictId: "151",
		},
		{
			Name:                 "Kota Jakarta Pusat",
			ProvinceId:           6,
			RajaOngkirDistrictId: "152",
		},
		{
			Name:                 "Kota Jakarta Selatan",
			ProvinceId:           6,
			RajaOngkirDistrictId: "153",
		},
		{
			Name:                 "Kota Jakarta Timur",
			ProvinceId:           6,
			RajaOngkirDistrictId: "154",
		},
		{
			Name:                 "Kota Jakarta Utara",
			ProvinceId:           6,
			RajaOngkirDistrictId: "155",
		},
		{
			Name:                 "Kota Jambi",
			ProvinceId:           8,
			RajaOngkirDistrictId: "156",
		},
		{
			Name:                 "Kabupaten Jayapura",
			ProvinceId:           24,
			RajaOngkirDistrictId: "157",
		},
		{
			Name:                 "Kota Jayapura",
			ProvinceId:           24,
			RajaOngkirDistrictId: "158",
		},
		{
			Name:                 "Kabupaten Jayawijaya",
			ProvinceId:           24,
			RajaOngkirDistrictId: "159",
		},
		{
			Name:                 "Kabupaten Jember",
			ProvinceId:           11,
			RajaOngkirDistrictId: "160",
		},
		{
			Name:                 "Kabupaten Jembrana",
			ProvinceId:           1,
			RajaOngkirDistrictId: "161",
		},
		{
			Name:                 "Kabupaten Jeneponto",
			ProvinceId:           28,
			RajaOngkirDistrictId: "162",
		},
		{
			Name:                 "Kabupaten Jepara",
			ProvinceId:           10,
			RajaOngkirDistrictId: "163",
		},
		{
			Name:                 "Kabupaten Jombang",
			ProvinceId:           11,
			RajaOngkirDistrictId: "164",
		},
		{
			Name:                 "Kabupaten Kaimana",
			ProvinceId:           25,
			RajaOngkirDistrictId: "165",
		},
		{
			Name:                 "Kabupaten Kampar",
			ProvinceId:           26,
			RajaOngkirDistrictId: "166",
		},
		{
			Name:                 "Kabupaten Kapuas",
			ProvinceId:           14,
			RajaOngkirDistrictId: "167",
		},
		{
			Name:                 "Kabupaten Kapuas Hulu",
			ProvinceId:           12,
			RajaOngkirDistrictId: "168",
		},
		{
			Name:                 "Kabupaten Karanganyar",
			ProvinceId:           10,
			RajaOngkirDistrictId: "169",
		},
		{
			Name:                 "Kabupaten Karangasem",
			ProvinceId:           1,
			RajaOngkirDistrictId: "170",
		},
		{
			Name:                 "Kabupaten Karawang",
			ProvinceId:           9,
			RajaOngkirDistrictId: "171",
		},
		{
			Name:                 "Kabupaten Karimun",
			ProvinceId:           17,
			RajaOngkirDistrictId: "172",
		},
		{
			Name:                 "Kabupaten Karo",
			ProvinceId:           34,
			RajaOngkirDistrictId: "173",
		},
		{
			Name:                 "Kabupaten Katingan",
			ProvinceId:           14,
			RajaOngkirDistrictId: "174",
		},
		{
			Name:                 "Kabupaten Kaur",
			ProvinceId:           4,
			RajaOngkirDistrictId: "175",
		},
		{
			Name:                 "Kabupaten Kayong Utara",
			ProvinceId:           12,
			RajaOngkirDistrictId: "176",
		},
		{
			Name:                 "Kabupaten Kebumen",
			ProvinceId:           10,
			RajaOngkirDistrictId: "177",
		},
		{
			Name:                 "Kabupaten Kediri",
			ProvinceId:           11,
			RajaOngkirDistrictId: "178",
		},
		{
			Name:                 "Kota Kediri",
			ProvinceId:           11,
			RajaOngkirDistrictId: "179",
		},
		{
			Name:                 "Kabupaten Keerom",
			ProvinceId:           24,
			RajaOngkirDistrictId: "180",
		},
		{
			Name:                 "Kabupaten Kendal",
			ProvinceId:           10,
			RajaOngkirDistrictId: "181",
		},
		{
			Name:                 "Kota Kendari",
			ProvinceId:           30,
			RajaOngkirDistrictId: "182",
		},
		{
			Name:                 "Kabupaten Kepahiang",
			ProvinceId:           4,
			RajaOngkirDistrictId: "183",
		},
		{
			Name:                 "Kabupaten Kepulauan Anambas",
			ProvinceId:           17,
			RajaOngkirDistrictId: "184",
		},
		{
			Name:                 "Kabupaten Kepulauan Aru",
			ProvinceId:           19,
			RajaOngkirDistrictId: "185",
		},
		{
			Name:                 "Kabupaten Kepulauan Mentawai",
			ProvinceId:           32,
			RajaOngkirDistrictId: "186",
		},
		{
			Name:                 "Kabupaten Kepulauan Meranti",
			ProvinceId:           26,
			RajaOngkirDistrictId: "187",
		},
		{
			Name:                 "Kabupaten Kepulauan Sangihe",
			ProvinceId:           31,
			RajaOngkirDistrictId: "188",
		},
		{
			Name:                 "Kabupaten Kepulauan Seribu",
			ProvinceId:           6,
			RajaOngkirDistrictId: "189",
		},
		{
			Name:                 "Kabupaten Kepulauan Siau Tagulandang Biaro (Sitaro)",
			ProvinceId:           31,
			RajaOngkirDistrictId: "190",
		},
		{
			Name:                 "Kabupaten Kepulauan Sula",
			ProvinceId:           20,
			RajaOngkirDistrictId: "191",
		},
		{
			Name:                 "Kabupaten Kepulauan Talaud",
			ProvinceId:           31,
			RajaOngkirDistrictId: "192",
		},
		{
			Name:                 "Kabupaten Kepulauan Yapen (Yapen Waropen)",
			ProvinceId:           24,
			RajaOngkirDistrictId: "193",
		},
		{
			Name:                 "Kabupaten Kerinci",
			ProvinceId:           8,
			RajaOngkirDistrictId: "194",
		},
		{
			Name:                 "Kabupaten Ketapang",
			ProvinceId:           12,
			RajaOngkirDistrictId: "195",
		},
		{
			Name:                 "Kabupaten Klaten",
			ProvinceId:           10,
			RajaOngkirDistrictId: "196",
		},
		{
			Name:                 "Kabupaten Klungkung",
			ProvinceId:           1,
			RajaOngkirDistrictId: "197",
		},
		{
			Name:                 "Kabupaten Kolaka",
			ProvinceId:           30,
			RajaOngkirDistrictId: "198",
		},
		{
			Name:                 "Kabupaten Kolaka Utara",
			ProvinceId:           30,
			RajaOngkirDistrictId: "199",
		},
		{
			Name:                 "Kabupaten Konawe",
			ProvinceId:           30,
			RajaOngkirDistrictId: "200",
		},
		{
			Name:                 "Kabupaten Konawe Selatan",
			ProvinceId:           30,
			RajaOngkirDistrictId: "201",
		},
		{
			Name:                 "Kabupaten Konawe Utara",
			ProvinceId:           30,
			RajaOngkirDistrictId: "202",
		},
		{
			Name:                 "Kabupaten Kotabaru",
			ProvinceId:           13,
			RajaOngkirDistrictId: "203",
		},
		{
			Name:                 "Kota Kotamobagu",
			ProvinceId:           31,
			RajaOngkirDistrictId: "204",
		},
		{
			Name:                 "Kabupaten Kotawaringin Barat",
			ProvinceId:           14,
			RajaOngkirDistrictId: "205",
		},
		{
			Name:                 "Kabupaten Kotawaringin Timur",
			ProvinceId:           14,
			RajaOngkirDistrictId: "206",
		},
		{
			Name:                 "Kabupaten Kuantan Singingi",
			ProvinceId:           26,
			RajaOngkirDistrictId: "207",
		},
		{
			Name:                 "Kabupaten Kubu Raya",
			ProvinceId:           12,
			RajaOngkirDistrictId: "208",
		},
		{
			Name:                 "Kabupaten Kudus",
			ProvinceId:           10,
			RajaOngkirDistrictId: "209",
		},
		{
			Name:                 "Kabupaten Kulon Progo",
			ProvinceId:           5,
			RajaOngkirDistrictId: "210",
		},
		{
			Name:                 "Kabupaten Kuningan",
			ProvinceId:           9,
			RajaOngkirDistrictId: "211",
		},
		{
			Name:                 "Kabupaten Kupang",
			ProvinceId:           23,
			RajaOngkirDistrictId: "212",
		},
		{
			Name:                 "Kota Kupang",
			ProvinceId:           23,
			RajaOngkirDistrictId: "213",
		},
		{
			Name:                 "Kabupaten Kutai Barat",
			ProvinceId:           15,
			RajaOngkirDistrictId: "214",
		},
		{
			Name:                 "Kabupaten Kutai Kartanegara",
			ProvinceId:           15,
			RajaOngkirDistrictId: "215",
		},
		{
			Name:                 "Kabupaten Kutai Timur",
			ProvinceId:           15,
			RajaOngkirDistrictId: "216",
		},
		{
			Name:                 "Kabupaten Labuhan Batu",
			ProvinceId:           34,
			RajaOngkirDistrictId: "217",
		},
		{
			Name:                 "Kabupaten Labuhan Batu Selatan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "218",
		},
		{
			Name:                 "Kabupaten Labuhan Batu Utara",
			ProvinceId:           34,
			RajaOngkirDistrictId: "219",
		},
		{
			Name:                 "Kabupaten Lahat",
			ProvinceId:           33,
			RajaOngkirDistrictId: "220",
		},
		{
			Name:                 "Kabupaten Lamandau",
			ProvinceId:           14,
			RajaOngkirDistrictId: "221",
		},
		{
			Name:                 "Kabupaten Lamongan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "222",
		},
		{
			Name:                 "Kabupaten Lampung Barat",
			ProvinceId:           18,
			RajaOngkirDistrictId: "223",
		},
		{
			Name:                 "Kabupaten Lampung Selatan",
			ProvinceId:           18,
			RajaOngkirDistrictId: "224",
		},
		{
			Name:                 "Kabupaten Lampung Tengah",
			ProvinceId:           18,
			RajaOngkirDistrictId: "225",
		},
		{
			Name:                 "Kabupaten Lampung Timur",
			ProvinceId:           18,
			RajaOngkirDistrictId: "226",
		},
		{
			Name:                 "Kabupaten Lampung Utara",
			ProvinceId:           18,
			RajaOngkirDistrictId: "227",
		},
		{
			Name:                 "Kabupaten Landak",
			ProvinceId:           12,
			RajaOngkirDistrictId: "228",
		},
		{
			Name:                 "Kabupaten Langkat",
			ProvinceId:           34,
			RajaOngkirDistrictId: "229",
		},
		{
			Name:                 "Kota Langsa",
			ProvinceId:           21,
			RajaOngkirDistrictId: "230",
		},
		{
			Name:                 "Kabupaten Lanny Jaya",
			ProvinceId:           24,
			RajaOngkirDistrictId: "231",
		},
		{
			Name:                 "Kabupaten Lebak",
			ProvinceId:           3,
			RajaOngkirDistrictId: "232",
		},
		{
			Name:                 "Kabupaten Lebong",
			ProvinceId:           4,
			RajaOngkirDistrictId: "233",
		},
		{
			Name:                 "Kabupaten Lembata",
			ProvinceId:           23,
			RajaOngkirDistrictId: "234",
		},
		{
			Name:                 "Kota Lhokseumawe",
			ProvinceId:           21,
			RajaOngkirDistrictId: "235",
		},
		{
			Name:                 "Kabupaten Lima Puluh Koto/Kota",
			ProvinceId:           32,
			RajaOngkirDistrictId: "236",
		},
		{
			Name:                 "Kabupaten Lingga",
			ProvinceId:           17,
			RajaOngkirDistrictId: "237",
		},
		{
			Name:                 "Kabupaten Lombok Barat",
			ProvinceId:           22,
			RajaOngkirDistrictId: "238",
		},
		{
			Name:                 "Kabupaten Lombok Tengah",
			ProvinceId:           22,
			RajaOngkirDistrictId: "239",
		},
		{
			Name:                 "Kabupaten Lombok Timur",
			ProvinceId:           22,
			RajaOngkirDistrictId: "240",
		},
		{
			Name:                 "Kabupaten Lombok Utara",
			ProvinceId:           22,
			RajaOngkirDistrictId: "241",
		},
		{
			Name:                 "Kota Lubuk Linggau",
			ProvinceId:           33,
			RajaOngkirDistrictId: "242",
		},
		{
			Name:                 "Kabupaten Lumajang",
			ProvinceId:           11,
			RajaOngkirDistrictId: "243",
		},
		{
			Name:                 "Kabupaten Luwu",
			ProvinceId:           28,
			RajaOngkirDistrictId: "244",
		},
		{
			Name:                 "Kabupaten Luwu Timur",
			ProvinceId:           28,
			RajaOngkirDistrictId: "245",
		},
		{
			Name:                 "Kabupaten Luwu Utara",
			ProvinceId:           28,
			RajaOngkirDistrictId: "246",
		},
		{
			Name:                 "Kabupaten Madiun",
			ProvinceId:           11,
			RajaOngkirDistrictId: "247",
		},
		{
			Name:                 "Kota Madiun",
			ProvinceId:           11,
			RajaOngkirDistrictId: "248",
		},
		{
			Name:                 "Kabupaten Magelang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "249",
		},
		{
			Name:                 "Kota Magelang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "250",
		},
		{
			Name:                 "Kabupaten Magetan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "251",
		},
		{
			Name:                 "Kabupaten Majalengka",
			ProvinceId:           9,
			RajaOngkirDistrictId: "252",
		},
		{
			Name:                 "Kabupaten Majene",
			ProvinceId:           27,
			RajaOngkirDistrictId: "253",
		},
		{
			Name:                 "Kota Makassar",
			ProvinceId:           28,
			RajaOngkirDistrictId: "254",
		},
		{
			Name:                 "Kabupaten Malang",
			ProvinceId:           11,
			RajaOngkirDistrictId: "255",
		},
		{
			Name:                 "Kota Malang",
			ProvinceId:           11,
			RajaOngkirDistrictId: "256",
		},
		{
			Name:                 "Kabupaten Malinau",
			ProvinceId:           16,
			RajaOngkirDistrictId: "257",
		},
		{
			Name:                 "Kabupaten Maluku Barat Daya",
			ProvinceId:           19,
			RajaOngkirDistrictId: "258",
		},
		{
			Name:                 "Kabupaten Maluku Tengah",
			ProvinceId:           19,
			RajaOngkirDistrictId: "259",
		},
		{
			Name:                 "Kabupaten Maluku Tenggara",
			ProvinceId:           19,
			RajaOngkirDistrictId: "260",
		},
		{
			Name:                 "Kabupaten Maluku Tenggara Barat",
			ProvinceId:           19,
			RajaOngkirDistrictId: "261",
		},
		{
			Name:                 "Kabupaten Mamasa",
			ProvinceId:           27,
			RajaOngkirDistrictId: "262",
		},
		{
			Name:                 "Kabupaten Mamberamo Raya",
			ProvinceId:           24,
			RajaOngkirDistrictId: "263",
		},
		{
			Name:                 "Kabupaten Mamberamo Tengah",
			ProvinceId:           24,
			RajaOngkirDistrictId: "264",
		},
		{
			Name:                 "Kabupaten Mamuju",
			ProvinceId:           27,
			RajaOngkirDistrictId: "265",
		},
		{
			Name:                 "Kabupaten Mamuju Utara",
			ProvinceId:           27,
			RajaOngkirDistrictId: "266",
		},
		{
			Name:                 "Kota Manado",
			ProvinceId:           31,
			RajaOngkirDistrictId: "267",
		},
		{
			Name:                 "Kabupaten Mandailing Natal",
			ProvinceId:           34,
			RajaOngkirDistrictId: "268",
		},
		{
			Name:                 "Kabupaten Manggarai",
			ProvinceId:           23,
			RajaOngkirDistrictId: "269",
		},
		{
			Name:                 "Kabupaten Manggarai Barat",
			ProvinceId:           23,
			RajaOngkirDistrictId: "270",
		},
		{
			Name:                 "Kabupaten Manggarai Timur",
			ProvinceId:           23,
			RajaOngkirDistrictId: "271",
		},
		{
			Name:                 "Kabupaten Manokwari",
			ProvinceId:           25,
			RajaOngkirDistrictId: "272",
		},
		{
			Name:                 "Kabupaten Manokwari Selatan",
			ProvinceId:           25,
			RajaOngkirDistrictId: "273",
		},
		{
			Name:                 "Kabupaten Mappi",
			ProvinceId:           24,
			RajaOngkirDistrictId: "274",
		},
		{
			Name:                 "Kabupaten Maros",
			ProvinceId:           28,
			RajaOngkirDistrictId: "275",
		},
		{
			Name:                 "Kota Mataram",
			ProvinceId:           22,
			RajaOngkirDistrictId: "276",
		},
		{
			Name:                 "Kabupaten Maybrat",
			ProvinceId:           25,
			RajaOngkirDistrictId: "277",
		},
		{
			Name:                 "Kota Medan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "278",
		},
		{
			Name:                 "Kabupaten Melawi",
			ProvinceId:           12,
			RajaOngkirDistrictId: "279",
		},
		{
			Name:                 "Kabupaten Merangin",
			ProvinceId:           8,
			RajaOngkirDistrictId: "280",
		},
		{
			Name:                 "Kabupaten Merauke",
			ProvinceId:           24,
			RajaOngkirDistrictId: "281",
		},
		{
			Name:                 "Kabupaten Mesuji",
			ProvinceId:           18,
			RajaOngkirDistrictId: "282",
		},
		{
			Name:                 "Kota Metro",
			ProvinceId:           18,
			RajaOngkirDistrictId: "283",
		},
		{
			Name:                 "Kabupaten Mimika",
			ProvinceId:           24,
			RajaOngkirDistrictId: "284",
		},
		{
			Name:                 "Kabupaten Minahasa",
			ProvinceId:           31,
			RajaOngkirDistrictId: "285",
		},
		{
			Name:                 "Kabupaten Minahasa Selatan",
			ProvinceId:           31,
			RajaOngkirDistrictId: "286",
		},
		{
			Name:                 "Kabupaten Minahasa Tenggara",
			ProvinceId:           31,
			RajaOngkirDistrictId: "287",
		},
		{
			Name:                 "Kabupaten Minahasa Utara",
			ProvinceId:           31,
			RajaOngkirDistrictId: "288",
		},
		{
			Name:                 "Kabupaten Mojokerto",
			ProvinceId:           11,
			RajaOngkirDistrictId: "289",
		},
		{
			Name:                 "Kota Mojokerto",
			ProvinceId:           11,
			RajaOngkirDistrictId: "290",
		},
		{
			Name:                 "Kabupaten Morowali",
			ProvinceId:           29,
			RajaOngkirDistrictId: "291",
		},
		{
			Name:                 "Kabupaten Muara Enim",
			ProvinceId:           33,
			RajaOngkirDistrictId: "292",
		},
		{
			Name:                 "Kabupaten Muaro Jambi",
			ProvinceId:           8,
			RajaOngkirDistrictId: "293",
		},
		{
			Name:                 "Kabupaten Muko Muko",
			ProvinceId:           4,
			RajaOngkirDistrictId: "294",
		},
		{
			Name:                 "Kabupaten Muna",
			ProvinceId:           30,
			RajaOngkirDistrictId: "295",
		},
		{
			Name:                 "Kabupaten Murung Raya",
			ProvinceId:           14,
			RajaOngkirDistrictId: "296",
		},
		{
			Name:                 "Kabupaten Musi Banyuasin",
			ProvinceId:           33,
			RajaOngkirDistrictId: "297",
		},
		{
			Name:                 "Kabupaten Musi Rawas",
			ProvinceId:           33,
			RajaOngkirDistrictId: "298",
		},
		{
			Name:                 "Kabupaten Nabire",
			ProvinceId:           24,
			RajaOngkirDistrictId: "299",
		},
		{
			Name:                 "Kabupaten Nagan Raya",
			ProvinceId:           21,
			RajaOngkirDistrictId: "300",
		},
		{
			Name:                 "Kabupaten Nagekeo",
			ProvinceId:           23,
			RajaOngkirDistrictId: "301",
		},
		{
			Name:                 "Kabupaten Natuna",
			ProvinceId:           17,
			RajaOngkirDistrictId: "302",
		},
		{
			Name:                 "Kabupaten Nduga",
			ProvinceId:           24,
			RajaOngkirDistrictId: "303",
		},
		{
			Name:                 "Kabupaten Ngada",
			ProvinceId:           23,
			RajaOngkirDistrictId: "304",
		},
		{
			Name:                 "Kabupaten Nganjuk",
			ProvinceId:           11,
			RajaOngkirDistrictId: "305",
		},
		{
			Name:                 "Kabupaten Ngawi",
			ProvinceId:           11,
			RajaOngkirDistrictId: "306",
		},
		{
			Name:                 "Kabupaten Nias",
			ProvinceId:           34,
			RajaOngkirDistrictId: "307",
		},
		{
			Name:                 "Kabupaten Nias Barat",
			ProvinceId:           34,
			RajaOngkirDistrictId: "308",
		},
		{
			Name:                 "Kabupaten Nias Selatan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "309",
		},
		{
			Name:                 "Kabupaten Nias Utara",
			ProvinceId:           34,
			RajaOngkirDistrictId: "310",
		},
		{
			Name:                 "Kabupaten Nunukan",
			ProvinceId:           16,
			RajaOngkirDistrictId: "311",
		},
		{
			Name:                 "Kabupaten Ogan Ilir",
			ProvinceId:           33,
			RajaOngkirDistrictId: "312",
		},
		{
			Name:                 "Kabupaten Ogan Komering Ilir",
			ProvinceId:           33,
			RajaOngkirDistrictId: "313",
		},
		{
			Name:                 "Kabupaten Ogan Komering Ulu",
			ProvinceId:           33,
			RajaOngkirDistrictId: "314",
		},
		{
			Name:                 "Kabupaten Ogan Komering Ulu Selatan",
			ProvinceId:           33,
			RajaOngkirDistrictId: "315",
		},
		{
			Name:                 "Kabupaten Ogan Komering Ulu Timur",
			ProvinceId:           33,
			RajaOngkirDistrictId: "316",
		},
		{
			Name:                 "Kabupaten Pacitan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "317",
		},
		{
			Name:                 "Kota Padang",
			ProvinceId:           32,
			RajaOngkirDistrictId: "318",
		},
		{
			Name:                 "Kabupaten Padang Lawas",
			ProvinceId:           34,
			RajaOngkirDistrictId: "319",
		},
		{
			Name:                 "Kabupaten Padang Lawas Utara",
			ProvinceId:           34,
			RajaOngkirDistrictId: "320",
		},
		{
			Name:                 "Kota Padang Panjang",
			ProvinceId:           32,
			RajaOngkirDistrictId: "321",
		},
		{
			Name:                 "Kabupaten Padang Pariaman",
			ProvinceId:           32,
			RajaOngkirDistrictId: "322",
		},
		{
			Name:                 "Kota Padang Sidempuan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "323",
		},
		{
			Name:                 "Kota Pagar Alam",
			ProvinceId:           33,
			RajaOngkirDistrictId: "324",
		},
		{
			Name:                 "Kabupaten Pakpak Bharat",
			ProvinceId:           34,
			RajaOngkirDistrictId: "325",
		},
		{
			Name:                 "Kota Palangka Raya",
			ProvinceId:           14,
			RajaOngkirDistrictId: "326",
		},
		{
			Name:                 "Kota Palembang",
			ProvinceId:           33,
			RajaOngkirDistrictId: "327",
		},
		{
			Name:                 "Kota Palopo",
			ProvinceId:           28,
			RajaOngkirDistrictId: "328",
		},
		{
			Name:                 "Kota Palu",
			ProvinceId:           29,
			RajaOngkirDistrictId: "329",
		},
		{
			Name:                 "Kabupaten Pamekasan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "330",
		},
		{
			Name:                 "Kabupaten Pandeglang",
			ProvinceId:           3,
			RajaOngkirDistrictId: "331",
		},
		{
			Name:                 "Kabupaten Pangandaran",
			ProvinceId:           9,
			RajaOngkirDistrictId: "332",
		},
		{
			Name:                 "Kabupaten Pangkajene Kepulauan",
			ProvinceId:           28,
			RajaOngkirDistrictId: "333",
		},
		{
			Name:                 "Kota Pangkal Pinang",
			ProvinceId:           2,
			RajaOngkirDistrictId: "334",
		},
		{
			Name:                 "Kabupaten Paniai",
			ProvinceId:           24,
			RajaOngkirDistrictId: "335",
		},
		{
			Name:                 "Kabupaten Manggarai Timur",
			ProvinceId:           28,
			RajaOngkirDistrictId: "336",
		},
		{
			Name:                 "Kota Pariaman",
			ProvinceId:           32,
			RajaOngkirDistrictId: "337",
		},
		{
			Name:                 "Kabupaten Parigi Moutong",
			ProvinceId:           29,
			RajaOngkirDistrictId: "338",
		},
		{
			Name:                 "Kabupaten Pasaman",
			ProvinceId:           32,
			RajaOngkirDistrictId: "339",
		},
		{
			Name:                 "Kabupaten Pasaman Barat",
			ProvinceId:           32,
			RajaOngkirDistrictId: "340",
		},
		{
			Name:                 "Kabupaten Paser",
			ProvinceId:           15,
			RajaOngkirDistrictId: "341",
		},
		{
			Name:                 "Kabupaten Pasuruan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "342",
		},
		{
			Name:                 "Kota Pasuruan",
			ProvinceId:           11,
			RajaOngkirDistrictId: "343",
		},
		{
			Name:                 "Kabupaten Pati",
			ProvinceId:           10,
			RajaOngkirDistrictId: "344",
		},
		{
			Name:                 "Kota Payakumbuh",
			ProvinceId:           32,
			RajaOngkirDistrictId: "345",
		},
		{
			Name:                 "Kabupaten Pegunungan Arfak",
			ProvinceId:           25,
			RajaOngkirDistrictId: "346",
		},
		{
			Name:                 "Kabupaten Pegunungan Bintang",
			ProvinceId:           24,
			RajaOngkirDistrictId: "347",
		},
		{
			Name:                 "Kabupaten Pekalongan",
			ProvinceId:           10,
			RajaOngkirDistrictId: "348",
		},
		{
			Name:                 "Kota Pekalongan",
			ProvinceId:           10,
			RajaOngkirDistrictId: "349",
		},
		{
			Name:                 "Kota Pekanbaru",
			ProvinceId:           26,
			RajaOngkirDistrictId: "350",
		},
		{
			Name:                 "Kabupaten Pelalawan",
			ProvinceId:           26,
			RajaOngkirDistrictId: "351",
		},
		{
			Name:                 "Kabupaten Pemalang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "352",
		},
		{
			Name:                 "Kota Pematang Siantar",
			ProvinceId:           33,
			RajaOngkirDistrictId: "353",
		},
		{
			Name:                 "Kabupaten Penajam Paser Utara",
			ProvinceId:           15,
			RajaOngkirDistrictId: "354",
		},
		{
			Name:                 "Kabupaten Pesawaran",
			ProvinceId:           18,
			RajaOngkirDistrictId: "355",
		},
		{
			Name:                 "Kabupaten Pesisir Barat",
			ProvinceId:           18,
			RajaOngkirDistrictId: "356",
		},
		{
			Name:                 "Kabupaten Pesisir Selatan",
			ProvinceId:           32,
			RajaOngkirDistrictId: "357",
		},
		{
			Name:                 "Kabupaten Pidie",
			ProvinceId:           21,
			RajaOngkirDistrictId: "358",
		},
		{
			Name:                 "Kabupaten Pidie Jaya",
			ProvinceId:           21,
			RajaOngkirDistrictId: "359",
		},
		{
			Name:                 "Kabupaten Pinrang",
			ProvinceId:           28,
			RajaOngkirDistrictId: "360",
		},
		{
			Name:                 "Kabupaten Pohuwato",
			ProvinceId:           7,
			RajaOngkirDistrictId: "361",
		},
		{
			Name:                 "Kabupaten Polewali Mandar",
			ProvinceId:           27,
			RajaOngkirDistrictId: "362",
		},
		{
			Name:                 "Kabupaten Ponorogo",
			ProvinceId:           11,
			RajaOngkirDistrictId: "363",
		},
		{
			Name:                 "Kabupaten Pontianak",
			ProvinceId:           12,
			RajaOngkirDistrictId: "364",
		},
		{
			Name:                 "Kota Pontianak",
			ProvinceId:           12,
			RajaOngkirDistrictId: "365",
		},
		{
			Name:                 "Kabupaten Poso",
			ProvinceId:           29,
			RajaOngkirDistrictId: "366",
		},
		{
			Name:                 "Kota Prabumulih",
			ProvinceId:           33,
			RajaOngkirDistrictId: "367",
		},
		{
			Name:                 "Kabupaten Pringsewu",
			ProvinceId:           18,
			RajaOngkirDistrictId: "368",
		},
		{
			Name:                 "Kabupaten Probolinggo",
			ProvinceId:           11,
			RajaOngkirDistrictId: "369",
		},
		{
			Name:                 "Kota Probolinggo",
			ProvinceId:           11,
			RajaOngkirDistrictId: "370",
		},
		{
			Name:                 "Kabupaten Pulang Pisau",
			ProvinceId:           14,
			RajaOngkirDistrictId: "371",
		},
		{
			Name:                 "Kabupaten Pulau Morotai",
			ProvinceId:           20,
			RajaOngkirDistrictId: "372",
		},
		{
			Name:                 "Kabupaten Puncak",
			ProvinceId:           24,
			RajaOngkirDistrictId: "373",
		},
		{
			Name:                 "Kabupaten Puncak Jaya",
			ProvinceId:           24,
			RajaOngkirDistrictId: "374",
		},
		{
			Name:                 "Kabupaten Purbalingga",
			ProvinceId:           10,
			RajaOngkirDistrictId: "375",
		},
		{
			Name:                 "Kabupaten Purwakarta",
			ProvinceId:           9,
			RajaOngkirDistrictId: "376",
		},
		{
			Name:                 "Kabupaten Purworejo",
			ProvinceId:           10,
			RajaOngkirDistrictId: "377",
		},
		{
			Name:                 "Kabupaten Raja Ampat",
			ProvinceId:           25,
			RajaOngkirDistrictId: "378",
		},
		{
			Name:                 "Kabupaten Rejang Lebong",
			ProvinceId:           4,
			RajaOngkirDistrictId: "379",
		},
		{
			Name:                 "Kabupaten Rembang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "380",
		},
		{
			Name:                 "Kabupaten Rokan Hilir",
			ProvinceId:           26,
			RajaOngkirDistrictId: "381",
		},
		{
			Name:                 "Kabupaten Rokan Hulu",
			ProvinceId:           26,
			RajaOngkirDistrictId: "382",
		},
		{
			Name:                 "Kabupaten Rote Ndao",
			ProvinceId:           23,
			RajaOngkirDistrictId: "383",
		},
		{
			Name:                 "Kota Sabang",
			ProvinceId:           21,
			RajaOngkirDistrictId: "384",
		},
		{
			Name:                 "Kabupaten Sabu Raijua",
			ProvinceId:           23,
			RajaOngkirDistrictId: "385",
		},
		{
			Name:                 "Kota Salatiga",
			ProvinceId:           10,
			RajaOngkirDistrictId: "386",
		},
		{
			Name:                 "Kota Samarinda",
			ProvinceId:           15,
			RajaOngkirDistrictId: "387",
		},
		{
			Name:                 "Kabupaten Sambas",
			ProvinceId:           12,
			RajaOngkirDistrictId: "388",
		},
		{
			Name:                 "Kabupaten Samosir",
			ProvinceId:           34,
			RajaOngkirDistrictId: "389",
		},
		{
			Name:                 "Kabupaten Sampang",
			ProvinceId:           11,
			RajaOngkirDistrictId: "390",
		},
		{
			Name:                 "Kabupaten Sanggau",
			ProvinceId:           12,
			RajaOngkirDistrictId: "391",
		},
		{
			Name:                 "Kabupaten Sarmi",
			ProvinceId:           24,
			RajaOngkirDistrictId: "392",
		},
		{
			Name:                 "Kabupaten Sarolangun",
			ProvinceId:           8,
			RajaOngkirDistrictId: "393",
		},
		{
			Name:                 "Kota Sawah Lunto",
			ProvinceId:           32,
			RajaOngkirDistrictId: "394",
		},
		{
			Name:                 "Kabupaten Sekadau",
			ProvinceId:           12,
			RajaOngkirDistrictId: "395",
		},
		{
			Name:                 "Kabupaten Selayar (Kepulauan Selayar)",
			ProvinceId:           28,
			RajaOngkirDistrictId: "396",
		},
		{
			Name:                 "Kabupaten Seluma",
			ProvinceId:           4,
			RajaOngkirDistrictId: "397",
		},
		{
			Name:                 "Kabupaten Semarang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "398",
		},
		{
			Name:                 "Kota Semarang",
			ProvinceId:           10,
			RajaOngkirDistrictId: "399",
		},
		{
			Name:                 "Kabupaten Seram Bagian Barat",
			ProvinceId:           19,
			RajaOngkirDistrictId: "400",
		},
		{
			Name:                 "Kabupaten Seram Bagian Timur",
			ProvinceId:           19,
			RajaOngkirDistrictId: "401",
		},
		{
			Name:                 "Kabupaten Serang",
			ProvinceId:           3,
			RajaOngkirDistrictId: "402",
		},
		{
			Name:                 "Kota Serang",
			ProvinceId:           3,
			RajaOngkirDistrictId: "403",
		},
		{
			Name:                 "Kabupaten Serdang Bedagai",
			ProvinceId:           34,
			RajaOngkirDistrictId: "404",
		},
		{
			Name:                 "Kabupaten Seruyan",
			ProvinceId:           14,
			RajaOngkirDistrictId: "405",
		},
		{
			Name:                 "Kabupaten Siak",
			ProvinceId:           26,
			RajaOngkirDistrictId: "406",
		},
		{
			Name:                 "Kota Sibolga",
			ProvinceId:           34,
			RajaOngkirDistrictId: "407",
		},
		{
			Name:                 "Kabupaten Sidenreng Rappang/Rapang",
			ProvinceId:           28,
			RajaOngkirDistrictId: "408",
		},
		{
			Name:                 "Kabupaten Sidoarjo",
			ProvinceId:           11,
			RajaOngkirDistrictId: "409",
		},
		{
			Name:                 "Kabupaten Sigi",
			ProvinceId:           29,
			RajaOngkirDistrictId: "410",
		},
		{
			Name:                 "Kabupaten Sijunjung (Sawah Lunto Sijunjung)",
			ProvinceId:           32,
			RajaOngkirDistrictId: "411",
		},
		{
			Name:                 "Kabupaten Sikka",
			ProvinceId:           23,
			RajaOngkirDistrictId: "412",
		},
		{
			Name:                 "Kabupaten Simalungun",
			ProvinceId:           34,
			RajaOngkirDistrictId: "413",
		},
		{
			Name:                 "Kabupaten Simeulue",
			ProvinceId:           21,
			RajaOngkirDistrictId: "414",
		},
		{
			Name:                 "Kota Singkawang",
			ProvinceId:           12,
			RajaOngkirDistrictId: "415",
		},
		{
			Name:                 "Kabupaten Sinjai",
			ProvinceId:           28,
			RajaOngkirDistrictId: "416",
		},
		{
			Name:                 "Kabupaten Sintang",
			ProvinceId:           12,
			RajaOngkirDistrictId: "417",
		},
		{
			Name:                 "Kabupaten Situbondo",
			ProvinceId:           11,
			RajaOngkirDistrictId: "418",
		},
		{
			Name:                 "Kabupaten Sleman",
			ProvinceId:           5,
			RajaOngkirDistrictId: "419",
		},
		{
			Name:                 "Kabupaten Solok",
			ProvinceId:           32,
			RajaOngkirDistrictId: "420",
		},
		{
			Name:                 "Kota Solok",
			ProvinceId:           32,
			RajaOngkirDistrictId: "421",
		},
		{
			Name:                 "Kabupaten Solok Selatan",
			ProvinceId:           32,
			RajaOngkirDistrictId: "422",
		},
		{
			Name:                 "Kabupaten Soppeng",
			ProvinceId:           28,
			RajaOngkirDistrictId: "423",
		},
		{
			Name:                 "Kabupaten Sorong",
			ProvinceId:           25,
			RajaOngkirDistrictId: "424",
		},
		{
			Name:                 "Kota Sorong",
			ProvinceId:           25,
			RajaOngkirDistrictId: "425",
		},
		{
			Name:                 "Kabupaten Sorong Selatan",
			ProvinceId:           25,
			RajaOngkirDistrictId: "426",
		},
		{
			Name:                 "Kabupaten Sragen",
			ProvinceId:           10,
			RajaOngkirDistrictId: "427",
		},
		{
			Name:                 "Kabupaten Subang",
			ProvinceId:           9,
			RajaOngkirDistrictId: "428",
		},
		{
			Name:                 "Kota Subulussalam",
			ProvinceId:           21,
			RajaOngkirDistrictId: "429",
		},
		{
			Name:                 "Kabupaten Sukabumi",
			ProvinceId:           9,
			RajaOngkirDistrictId: "430",
		},
		{
			Name:                 "Kota Sukabumi",
			ProvinceId:           9,
			RajaOngkirDistrictId: "431",
		},
		{
			Name:                 "Kabupaten Sukamara",
			ProvinceId:           14,
			RajaOngkirDistrictId: "432",
		},
		{
			Name:                 "Kabupaten Sukoharjo",
			ProvinceId:           10,
			RajaOngkirDistrictId: "433",
		},
		{
			Name:                 "Kabupaten Sumba Barat",
			ProvinceId:           23,
			RajaOngkirDistrictId: "434",
		},
		{
			Name:                 "Kabupaten Sumba Barat Daya",
			ProvinceId:           23,
			RajaOngkirDistrictId: "435",
		},
		{
			Name:                 "Kabupaten Sumba Tengah",
			ProvinceId:           23,
			RajaOngkirDistrictId: "436",
		},
		{
			Name:                 "Kabupaten Sumba Timur",
			ProvinceId:           23,
			RajaOngkirDistrictId: "437",
		},
		{
			Name:                 "Kabupaten Sumbawa",
			ProvinceId:           22,
			RajaOngkirDistrictId: "438",
		},
		{
			Name:                 "Kabupaten Sumbawa Barat",
			ProvinceId:           22,
			RajaOngkirDistrictId: "439",
		},
		{
			Name:                 "Kabupaten Sumedang",
			ProvinceId:           9,
			RajaOngkirDistrictId: "440",
		},
		{
			Name:                 "Kabupaten Sumenep",
			ProvinceId:           11,
			RajaOngkirDistrictId: "441",
		},
		{
			Name:                 "Kota Sungaipenuh",
			ProvinceId:           8,
			RajaOngkirDistrictId: "442",
		},
		{
			Name:                 "Kabupaten Supiori",
			ProvinceId:           24,
			RajaOngkirDistrictId: "443",
		},
		{
			Name:                 "Kota Surabaya",
			ProvinceId:           11,
			RajaOngkirDistrictId: "444",
		},
		{
			Name:                 "Kota Surakarta (Solo)",
			ProvinceId:           10,
			RajaOngkirDistrictId: "445",
		},
		{
			Name:                 "Kabupaten Tabalong",
			ProvinceId:           13,
			RajaOngkirDistrictId: "446",
		},
		{
			Name:                 "Kabupaten Tabanan",
			ProvinceId:           1,
			RajaOngkirDistrictId: "447",
		},
		{
			Name:                 "Kabupaten Takalar",
			ProvinceId:           28,
			RajaOngkirDistrictId: "448",
		},
		{
			Name:                 "Kabupaten Tambrauw",
			ProvinceId:           25,
			RajaOngkirDistrictId: "449",
		},
		{
			Name:                 "Kabupaten Tana Tidung",
			ProvinceId:           16,
			RajaOngkirDistrictId: "450",
		},
		{
			Name:                 "Kabupaten Tana Toraja",
			ProvinceId:           28,
			RajaOngkirDistrictId: "451",
		},
		{
			Name:                 "Kabupaten Tanah Bumbu",
			ProvinceId:           13,
			RajaOngkirDistrictId: "452",
		},
		{
			Name:                 "Kabupaten Tanah Datar",
			ProvinceId:           32,
			RajaOngkirDistrictId: "453",
		},
		{
			Name:                 "Kabupaten Tanah Laut",
			ProvinceId:           13,
			RajaOngkirDistrictId: "454",
		},
		{
			Name:                 "Kabupaten Tangerang",
			ProvinceId:           3,
			RajaOngkirDistrictId: "455",
		},
		{
			Name:                 "Kota Tangerang",
			ProvinceId:           3,
			RajaOngkirDistrictId: "456",
		},
		{
			Name:                 "Kota Tangerang Selatan",
			ProvinceId:           3,
			RajaOngkirDistrictId: "457",
		},
		{
			Name:                 "Kabupaten Tanggamus",
			ProvinceId:           18,
			RajaOngkirDistrictId: "458",
		},
		{
			Name:                 "Kota Tanjung Balai",
			ProvinceId:           34,
			RajaOngkirDistrictId: "459",
		},
		{
			Name:                 "Kabupaten Tanjung Jabung Barat",
			ProvinceId:           8,
			RajaOngkirDistrictId: "460",
		},
		{
			Name:                 "Kabupaten Tanjung Jabung Timur",
			ProvinceId:           8,
			RajaOngkirDistrictId: "461",
		},
		{
			Name:                 "Kota Tanjung Pinang",
			ProvinceId:           17,
			RajaOngkirDistrictId: "462",
		},
		{
			Name:                 "Kabupaten Tapanuli Selatan",
			ProvinceId:           34,
			RajaOngkirDistrictId: "463",
		},
		{
			Name:                 "Kabupaten Tapanuli Tengah",
			ProvinceId:           34,
			RajaOngkirDistrictId: "464",
		},
		{
			Name:                 "Kabupaten Tapanuli Utara",
			ProvinceId:           34,
			RajaOngkirDistrictId: "465",
		},
		{
			Name:                 "Kabupaten Tapin",
			ProvinceId:           13,
			RajaOngkirDistrictId: "466",
		},
		{
			Name:                 "Kota Tarakan",
			ProvinceId:           16,
			RajaOngkirDistrictId: "467",
		},
		{
			Name:                 "Kabupaten Tasikmalaya",
			ProvinceId:           9,
			RajaOngkirDistrictId: "468",
		},
		{
			Name:                 "Kota Tasikmalaya",
			ProvinceId:           9,
			RajaOngkirDistrictId: "469",
		},
		{
			Name:                 "Kota Tebing Tinggi",
			ProvinceId:           34,
			RajaOngkirDistrictId: "470",
		},
		{
			Name:                 "Kabupaten Tebo",
			ProvinceId:           8,
			RajaOngkirDistrictId: "471",
		},
		{
			Name:                 "Kabupaten Tegal",
			ProvinceId:           10,
			RajaOngkirDistrictId: "472",
		},
		{
			Name:                 "Kota Tegal",
			ProvinceId:           10,
			RajaOngkirDistrictId: "473",
		},
		{
			Name:                 "Kabupaten Teluk Bintuni",
			ProvinceId:           25,
			RajaOngkirDistrictId: "474",
		},
		{
			Name:                 "Kabupaten Teluk Wondama",
			ProvinceId:           25,
			RajaOngkirDistrictId: "475",
		},
		{
			Name:                 "Kabupaten Temanggung",
			ProvinceId:           10,
			RajaOngkirDistrictId: "476",
		},
		{
			Name:                 "Kota Ternate",
			ProvinceId:           20,
			RajaOngkirDistrictId: "477",
		},
		{
			Name:                 "Kota Tidore Kepulauan",
			ProvinceId:           20,
			RajaOngkirDistrictId: "478",
		},
		{
			Name:                 "Kabupaten Timor Tengah Selatan",
			ProvinceId:           23,
			RajaOngkirDistrictId: "479",
		},
		{
			Name:                 "Kabupaten Timor Tengah Utara",
			ProvinceId:           23,
			RajaOngkirDistrictId: "480",
		},
		{
			Name:                 "Kabupaten Toba Samosir",
			ProvinceId:           34,
			RajaOngkirDistrictId: "481",
		},
		{
			Name:                 "Kabupaten Tojo Una-Una",
			ProvinceId:           29,
			RajaOngkirDistrictId: "482",
		},
		{
			Name:                 "Kabupaten Toli-Toli",
			ProvinceId:           29,
			RajaOngkirDistrictId: "483",
		},
		{
			Name:                 "Kabupaten Tolokara",
			ProvinceId:           24,
			RajaOngkirDistrictId: "484",
		},
		{
			Name:                 "Kota Tomohon",
			ProvinceId:           31,
			RajaOngkirDistrictId: "485",
		},
		{
			Name:                 "Kabupaten Toraja Utara",
			ProvinceId:           28,
			RajaOngkirDistrictId: "486",
		},
		{
			Name:                 "Kabupaten Trenggalek",
			ProvinceId:           11,
			RajaOngkirDistrictId: "487",
		},
		{
			Name:                 "Kota Tual",
			ProvinceId:           19,
			RajaOngkirDistrictId: "488",
		},
		{
			Name:                 "Kabupaten Tuban",
			ProvinceId:           11,
			RajaOngkirDistrictId: "489",
		},
		{
			Name:                 "Kabupaten Tulang Bawang",
			ProvinceId:           18,
			RajaOngkirDistrictId: "490",
		},
		{
			Name:                 "Kabupaten Tulang Bawang Barat",
			ProvinceId:           18,
			RajaOngkirDistrictId: "491",
		},
		{
			Name:                 "Kabupaten Tulungagung",
			ProvinceId:           11,
			RajaOngkirDistrictId: "492",
		},
		{
			Name:                 "Kabupaten Wajo",
			ProvinceId:           28,
			RajaOngkirDistrictId: "493",
		},
		{
			Name:                 "Kabupaten Wakatobi",
			ProvinceId:           30,
			RajaOngkirDistrictId: "494",
		},
		{
			Name:                 "Kabupaten Waropen",
			ProvinceId:           24,
			RajaOngkirDistrictId: "495",
		},
		{
			Name:                 "Kabupaten Way Kanan",
			ProvinceId:           18,
			RajaOngkirDistrictId: "496",
		},
		{
			Name:                 "Kabupaten Wonogiri",
			ProvinceId:           10,
			RajaOngkirDistrictId: "497",
		},
		{
			Name:                 "Kabupaten Wonosobo",
			ProvinceId:           10,
			RajaOngkirDistrictId: "498",
		},
		{
			Name:                 "Kabupaten Yahukimo",
			ProvinceId:           24,
			RajaOngkirDistrictId: "499",
		},
		{
			Name:                 "Kabupaten Yalimo",
			ProvinceId:           24,
			RajaOngkirDistrictId: "500",
		},
		{
			Name:                 "Kota Yogyakarta",
			ProvinceId:           5,
			RajaOngkirDistrictId: "501",
		},
	}

	if err := db.Create(districts).Error; err != nil {
		panic(err)

	}

	accounts := []*model.Accounts{
		{
			Username:      "testing",
			FullName:      "My Testing Account",
			Email:         "testing@mail.com",
			PhoneNumber:   "08982873823",
			Password:      "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
			WalletNumber:  "4200000000001",
			Gender:        "male",
			ShopName:      "XYZ SHOP",
			Birthdate:     time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
			Balance:       decimal.NewFromInt(0),
			WalletPin:     "123456",
			SellerBalance: decimal.NewFromInt(90000),
		},
		{
			Username:     "satoni",
			FullName:     "Ahmad Satoni",
			Email:        "satoni@mail.com",
			PhoneNumber:  "089828738222",
			Password:     "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
			WalletNumber: "4200000000002",
			Gender:       "male",
			Birthdate:    time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
			Balance:      decimal.NewFromInt(0),
		},
		{
			Username:      "satrianusa",
			FullName:      "Satria Nusa",
			Email:         "satria@mail.com",
			PhoneNumber:   "089345433823",
			Password:      "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
			WalletNumber:  "4200000000003",
			Gender:        "male",
			ShopName:      "Satria Shop",
			Birthdate:     time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
			Balance:       decimal.NewFromInt(0),
			WalletPin:     "123456",
			SellerBalance: decimal.NewFromInt(0),
		},
	}
	err := db.Create(accounts).Error

	if err != nil {
		panic(err)
	}

	categories := []*model.Category{
		{
			Name:  "Elektronik",
			Level: 1,
		},
		{
			Name:   "TV & Aksesoris",
			Level:  2,
			Parent: 1,
		},
		{
			Name:   "TV",
			Level:  3,
			Parent: 2,
		},
		{
			Name:   "Kelistrikan",
			Level:  2,
			Parent: 1,
		},
		{
			Name:   "Saklar",
			Level:  3,
			Parent: 2,
		},
	}

	err = db.Create(categories).Error

	if err != nil {
		panic(err)
	}

	products := []*model.Products{
		{
			Name:              "Minyak Goreng Refill Rose Brand 2L",
			Description:       "Minyak Goreng Rose Brand terbuat dari kelapa sawit pilihan berkualitas, diproses secara modern dengan teknologi tinggi secara higienis untuk membuat semya masakan menjadi lebih gurih dan lezat. Minyak Goreng Rose Brand mengandung BETA Karoten, omega 9, vitamin A dan E yang baik untuk tubuh.",
			CategoryID:        5,
			HazardousMaterial: false,
			Weight:            decimal.NewFromInt(22),
			Size:              decimal.NewFromInt(30),
			IsNew:             true,
			InternalSKU:       "OAKO OEasEF",
			ViewCount:         0,
			IsActive:          true,
			SellerID:          1,
		},
		{
			Name:              "Schneider Electric Leona Saklar Lampu - 2 Gang 2 Arah - LNA0600321",
			Description:       "Desain stylish dan minimalis untuk semua desain rumah Leona memiliki karakter berbentuk melingkar di setiap ujungnya. Desain yang tak lekang waktu dan sesuai untuk segala jenis rumah, serta memiliki berbagai varian untuk berbagai jenis kebutuhan, mulai dari saklar lampu, stop kontak schuko, tv, telepon dan data, hingga peredup lampu (dimmer). 2 Cara Pemasangan, sistem pencakar atau sekrup Saklar lampu dan stop kontak Leona hadir dengan 2 pilihan cara pemasangan. Dengan sistem pencakar dan sekrup yang memungkinkan untuk proyek renovasi maupun rumah baru. Sistem pencakar yang terlindungi, menjamin kekuatan dan daya cengkeram pada inbowdoost. Harga yang terjangkau dan kualitas terbaik Saklar lampu dan stop kontak Leona terbuat dari bahan polycarbonate berkualitas dan lebih aman karena lebih tahan panas, serta diperkuat dengan frame modul dari bahan logam yang menjamin kualitas, kekuatan dan tahan lebih lama.",
			CategoryID:        5,
			HazardousMaterial: false,
			Weight:            decimal.NewFromInt(22),
			Size:              decimal.NewFromInt(30),
			IsNew:             true,
			InternalSKU:       "OAKO OEKFOEF",
			ViewCount:         0,
			IsActive:          true,
			SellerID:          1,
		},
		{
			Name:              "Magsafe 2 Charger macbook 45w l 60w AIR l PRO - 45W",
			Description:       "MAGSAFE 2 LAGI PROMO MINGGU INI SILAHKAN ORDER SEBELUM HARGA KEMBALI NORMAL!!",
			CategoryID:        3,
			HazardousMaterial: false,
			Weight:            decimal.NewFromInt(22),
			Size:              decimal.NewFromInt(30),
			IsNew:             true,
			InternalSKU:       "OAKO",
			ViewCount:         0,
			IsActive:          true,
			SellerID:          3,
		},
	}

	err = db.Create(products).Error

	if err != nil {
		panic(err)
	}

	productImages := []*model.ProductImages{
		{
			ProductID: 2,
			URL:       "https://down-id.img.susercontent.com/file/1e16d71744f0b71db776f915facb6df9",
		},
		{
			ProductID: 3,
			URL:       "https://down-id.img.susercontent.com/file/id-11134207-7r991-lnif9zpjj4au82",
		},
		{
			ProductID: 3,
			URL:       "https://down-id.img.susercontent.com/file/2a4e6f610e903fe5dcce459b76a9081f",
		},
	}

	err = db.Create(productImages).Error

	if err != nil {
		panic(err)
	}

	productVideos := []*model.ProductVideos{
		{
			ProductID: 3,
			URL:       "https://www.youtube.com/embed/tR05rgXCFdk?si=jhsuqvYv8cBhPiu3",
		},
	}

	err = db.Create(productVideos).Error

	if err != nil {
		panic(err)
	}

	productVariants := []*model.ProductVariants{
		{
			ProductID: 1,
			Name:      "default_reserved_keyword",
		},
		{
			ProductID: 2,
			Name:      "Color",
		},
		{
			ProductID: 2,
			Name:      "Bahan",
		},
		{
			ProductID: 3,
			Name:      "Pilih produk",
		},
	}

	err = db.Create(productVariants).Error
	if err != nil {
		panic(err)
	}

	productVariantSelections := []*model.ProductVariantSelections{
		{
			ProductVariantID: 1,
			Name:             "default_reserved_keyword",
		},
		{
			ProductVariantID: 2,
			Name:             "Merah",
		},
		{
			ProductVariantID: 2,
			Name:             "Biru",
		},
		{
			ProductVariantID: 3,
			Name:             "Metal",
		},
		{
			ProductVariantID: 3,
			Name:             "Wood",
		},
		{
			ProductVariantID: 4,
			Name:             "45W",
		},
		{
			ProductVariantID: 4,
			Name:             "60W",
		},
	}

	err = db.Create(productVariantSelections).Error

	if err != nil {
		panic(err)
	}

	productVariantSelectionCombinations := []*model.ProductVariantSelectionCombinations{
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

	err = db.Create(productVariantSelectionCombinations).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&[]model.Couriers{
		{
			Name:        "jne",
			Description: "layanan JNE courier",
		},
		{
			Name:        "tiki",
			Description: "layanan TIKI courier",
		},
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&model.ProductOrders{
		CourierID:     1,
		AccountID:     2,
		DeliveryFee:   decimal.NewFromInt(15000),
		Province:      "Jawa Barat",
		District:      "Kabupaten Bandung",
		SubDistrict:   "Bojongsoang",
		Kelurahan:     "Sukapura",
		ZipCode:       "40851",
		AddressDetail: "Jl Telekomunikasi No 1 Bojongsoang",
		Status:        constant.StatusWaitingSellerConfirmation,
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&[]model.ProductOrderDetails{
		{
			ProductOrderID:                       1,
			ProductVariantSelectionCombinationID: 1,
			Quantity:                             2,
			IndividualPrice:                      decimal.NewFromInt(20000),
		},
		{
			ProductOrderID:                       1,
			ProductVariantSelectionCombinationID: 3,
			Quantity:                             1,
			IndividualPrice:                      decimal.NewFromInt(50000),
		},
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&model.SaleWalletTransactionHistories{
		ProductOrderID: 1,
		AccountID:      1,
		Type:           constant.SaleMoneyIncomeType,
		Amount:         decimal.NewFromInt(90000),
		From:           "4200913923913",
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&[]model.AccountAddress{
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
	}).Error

	if err != nil {
		panic(err)
	}

	accountCarts := []*model.AccountCarts{
		{
			AccountID:                            2,
			ProductVariantSelectionCombinationId: 1,
			Quantity:                             2,
		},
		{
			AccountID:                            2,
			ProductVariantSelectionCombinationId: 3,
			Quantity:                             1,
		},
		{
			AccountID:                            2,
			ProductVariantSelectionCombinationId: 7,
			Quantity:                             1,
		},
	}

	err = db.Create(accountCarts).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&[]model.SellerCouriers{
		{
			AccountID: 1,
			CourierID: 1,
		},
		{
			AccountID: 1,
			CourierID: 2,
		},
	}).Error

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully seed tables")
}
