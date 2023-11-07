package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Accounts struct {
	ID                      int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	FullName                string    `gorm:"type:varchar;not null"`
	Username                string    `gorm:"type:varchar"`
	Email                   string    `gorm:"type:varchar;not null"`
	PhoneNumber             string    `gorm:"type:varchar"`
	Password                string    `gorm:"type:varchar"`
	ShopName                string    `gorm:"type:varchar;unique"`
	Gender                  string    `gorm:"type:varchar"`
	Birthdate               time.Time `gorm:"type:timestamp"`
	ProfilePicture          string    `gorm:"type:varchar"`
	WalletNumber            string    `gorm:"type:varchar"`
	WalletPin               string
	Balance                 decimal.Decimal `gorm:"type:decimal;default:0"`
	SallerBalance           decimal.Decimal `gorm:"type:decimal;default:0"`
	ForgetPasswordToken     string          `gorm:"type:varchar"`
	ForgetPasswordExpiredAt time.Time       `gorm:"type:timestamp"`
	CreatedAt               time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt               time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt               time.Time       `gorm:"type:timestamp;default:null"`
}

type UsedEmail struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID int       `gorm:"foreignKey:AccountID;type:bigint;not null"`
	Email     string    `gorm:"type:varchar;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type MyWalletTransactionHistories struct {
	ID             int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID      int             `gorm:"foreignKey:AccountID;type:bigint;not null"`
	Type           string          `gorm:"type:varchar"`
	Amount         decimal.Decimal `gorm:"type:decimal"`
	ProductOrderID int             `gorm:"foreignKey:AccountID;type:bigint;default:null"`
	CreatedAt      time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt      time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt      time.Time       `gorm:"type:timestamp;default:null"`
}

type SaleWalletTransactionHistories struct {
	ID             int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID      int             `gorm:"foreignKey:AccountID;type:bigint;not null"`
	Type           string          `gorm:"type:varchar"`
	Amount         decimal.Decimal `gorm:"type:decimal"`
	ProductOrderID int             `gorm:"foreignKey:AccountID;type:bigint;default:null"`
	CreatedAt      time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt      time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt      time.Time       `gorm:"type:timestamp;default:null"`
}

type AccountCarts struct {
	ID                                   int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID                            int       `gorm:"foreignKey:AccountID;type:bigint;not null"`
	ProductVariantSelectionCombinationId int       `gorm:"foreignKey:product_variant_selection_combinations;type:bigint;not null"`
	Quantity                             int       `gorm:"type:int;not null"`
	CreatedAt                            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt                            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt                            time.Time `gorm:"type:timestamp;default:null"`
}

type SellerCouriers struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID int       `gorm:"foreignKey:AccountID;type:bigint;not null"`
	CourierID int       `gorm:"foreignKey:CourierID;type:bigint;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}
