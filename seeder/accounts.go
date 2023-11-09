package seeder

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var Accounts = []*model.Accounts{
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
