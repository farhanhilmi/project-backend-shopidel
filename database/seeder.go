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
			account_carts;
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
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully migrate tables")
}

func seeding() {
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
			SallerBalance: decimal.NewFromInt(90000),
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
			SallerBalance: decimal.NewFromInt(0),
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
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&[]model.AccountAddress{
		{
			AccountID: 2,
			Province: "DKI Jakarta",
			District: "Jakarta Selatan",
			SubDistrict: "Sub District 1",
			Kelurahan: "lurahan skuy living",
			ZipCode: "12230",
			RajaOngkirDistrictId: "153",
			Detail: "mambo jambo",
		},
		{
			AccountID: 2,
			Province: "DKI Jakarta",
			District: "Jakarta Timur",
			SubDistrict: "Sub District 2",
			Kelurahan: "lurahan teranjay",
			ZipCode: "14738",
			RajaOngkirDistrictId: "154",
			Detail: "rujak cireng",
			IsBuyerDefault: true,
		},
		{
			AccountID: 2,
			Province: "DKI Jakarta",
			District: "Jakarta Barat",
			SubDistrict: "Sub District 3",
			Kelurahan: "lurahan skuy living",
			RajaOngkirDistrictId: "151",
			ZipCode: "15405",
			Detail: "es jambu",
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

	fmt.Println("successfully seed tables")
}
