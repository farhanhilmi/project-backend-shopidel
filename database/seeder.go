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
		drop table if exists accounts;
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
	)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&model.UsedEmail{},
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

	fmt.Println("successfully seed tables")
}
