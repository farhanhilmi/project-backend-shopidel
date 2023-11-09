package database

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/seeder"
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
			seller_couriers
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
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully migrate tables")
}

func seeding() {

	err := db.Create(seeder.Accounts).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.Categories).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(&seeder.Products).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductImages).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductVideos).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductVariants).Error
	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductVariantSelections).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductVariantSelectionCombinations).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.Couriers).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductOrders).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.ProductOrderDetails).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.SaleWalletTransactionHistories).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.AccountAddress).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.AccountCarts).Error

	if err != nil {
		panic(err)
	}

	err = db.Create(seeder.SellerCouriers).Error

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully seed tables")
}
