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
			seller_page_selected_categories,
			districts,
			provinces,
			product_order_reviews,
			favorite_products,
			seller_couriers,
			account_carts,
			account_addresses,
			couriers,
			product_order_details,
			sale_wallet_transaction_histories,
			product_orders,
			MyWalletTransactionHistories,
			used_emails,
			products,
			product_videos,
			product_variants,
			product_variant_selections,
			product_variant_selection_combinations,
			product_images,
			my_wallet_transaction_histories,
			categories,
			accounts
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
		&model.SellerPageSelectedCategory{},
	)

	if err != nil {
		panic(err)
	}

	err = db.Exec(`
		ALTER TABLE accounts ALTER COLUMN wallet_number SET DEFAULT TO_CHAR(nextval('accounts_id_seq'::regclass),'"420"fm0000000000');
	`).Error
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
		seeder.SellerPageSelectedCategory,
	}

	for _, seed := range seeders {
		if err := db.Create(seed).Error; err != nil {
			panic(err)
		}
	}

	err := db.Exec(`
		ALTER SEQUENCE accounts_id_seq RESTART WITH 1000;
		ALTER SEQUENCE products_id_seq RESTART WITH 1000;
		ALTER SEQUENCE product_variant_selections_id_seq RESTART WITH 1000;
		ALTER SEQUENCE Product_variant_selection_combinations_id_seq RESTART WITH 1000;
		ALTER SEQUENCE product_variants_id_seq RESTART WITH 1000;
		ALTER SEQUENCE product_orders_id_seq RESTART WITH 1000;
	`).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully seed tables")
}
