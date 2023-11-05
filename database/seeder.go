package database

import (
	"fmt"
	"log"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
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
			MyWalletTransactionHistories;
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
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully migrate tables")
}

func seeding() {
	err := db.Create(&model.Accounts{
		Username:       "testing",
		FullName:       "My Testing Account",
		Email:          "testing@mail.com",
		PhoneNumber:    "08982873823",
		Password:       "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
		WalletNumber:   "7770000000001",
		Gender:         "male",
		Birthdate:      time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
		ProfilePicture: "https://cdn4.iconfinder.com/data/icons/web-ui-color/128/Account-512.png",
		Balance:        decimal.NewFromInt(0),
		WalletPin:      "123456",
	}).Error

	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = db.Create(&model.Accounts{
		Username:       "satoni",
		FullName:       "Ahmad Satoni",
		Email:          "satoni@mail.com",
		PhoneNumber:    "089828738222",
		Password:       "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
		WalletNumber:   "7770000000002",
		Gender:         "male",
		Birthdate:      time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
		ProfilePicture: "https://cdn4.iconfinder.com/data/icons/web-ui-color/128/Account-512.png",
		Balance:        decimal.NewFromInt(0),
	}).Error

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

	fmt.Println("successfully seed tables")
}
