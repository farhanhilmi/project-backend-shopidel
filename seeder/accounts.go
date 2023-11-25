package seeder

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

var Accounts = []*model.Accounts{
	{
		ID:            1,
		Username:      "testing",
		FullName:      "My Testing Account",
		Email:         "testing@mail.com",
		PhoneNumber:   "+6224278394834",
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
		ID:           2,
		Username:     "satoni",
		FullName:     "Ahmad Satoni",
		Email:        "satoni@yopmail.com",
		PhoneNumber:  "+6224278394835",
		Password:     "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
		WalletNumber: "4200000000002",
		Gender:       "male",
		Birthdate:    time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
		Balance:      decimal.NewFromInt(0),
	},
	{
		ID:            3,
		Username:      "satrianusa",
		FullName:      "Satria Nusa",
		Email:         "satria@mail.com",
		PhoneNumber:   "+6224278394836",
		Password:      "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
		WalletNumber:  "4200000000003",
		Gender:        "male",
		ShopName:      "Satria Shop",
		Birthdate:     time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
		Balance:       decimal.NewFromInt(0),
		WalletPin:     "123456",
		SellerBalance: decimal.NewFromInt(0),
	},
	{
		ID:              4,
		FullName:        "Gita Purnama",
		Username:        "gitapurnama",
		Email:           "gitapurnama@mail.com",
		PhoneNumber:     "+6224278394837",
		Password:        "$2a$14$ggRGSX9uKrEfapylGVadWee/P1yCOKduFFqnzNdq7U3ble5nxtNqC",
		ShopName:        "Jejak Trendi",
		Gender:          "male",
		Birthdate:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
		ProfilePicture:  "https://mangathrill.com/wp-content/uploads/2019/07/Portgas.D..Ace_.full_.5794251280x720.png",
		WalletNumber:    "4200000000004",
		WalletPin:       "123456",
		Balance:         decimal.NewFromInt(0),
		SellerBalance:   decimal.NewFromInt(0),
		ShopDescription: "We are the most trusted fashion shop",
		ShopOpeningHours: func() time.Time {
			tm, _ := time.Parse("15:04", "08:00")

			return tm
		}(),
		ShopClosingHours: func() time.Time {
			tm, _ := time.Parse("15:04", "20:00")

			return tm
		}(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
