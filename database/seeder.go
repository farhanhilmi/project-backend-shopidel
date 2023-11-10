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
			seller_couriers,
			favorite_products,
			product_order_reviews,
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
		&model.FavoriteProducts{},
		&model.ProductOrderReviews{},
		&model.Province{},
		&model.District{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully migrate tables")
}

func seeding() {
	seeders := []any{
		seeder.Accounts,
		seeder.Categories,
		seeder.Products,
		seeder.ProductImages,
		seeder.ProductVideos,
		seeder.ProductVariants,
		seeder.ProductVariantSelections,
		seeder.ProductVariantSelectionCombinations,
		seeder.Couriers,
		seeder.ProductOrders,
		seeder.ProductOrderDetails,
		seeder.SaleWalletTransactionHistories,
		seeder.AccountAddress,
		seeder.AccountCarts,
		seeder.SellerCouriers,
		seeder.ProductOrderReviews,
		seeder.Provinces,
		seeder.Districts,
	}

	for _, seed := range seeders {
		if err := db.Create(seed).Error; err != nil {
			panic(err)
		}
	}

	fmt.Println("successfully seed tables")
}
