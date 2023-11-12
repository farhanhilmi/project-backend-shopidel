package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Accounts struct {
	ID                      int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	FullName                string          `gorm:"type:varchar;not null"`
	Username                string          `gorm:"type:varchar;unique"`
	Email                   string          `gorm:"type:varchar;unique;not null"`
	PhoneNumber             string          `gorm:"type:varchar;unique;default:null"`
	Password                string          `gorm:"type:varchar;not null"`
	ShopName                string          `gorm:"type:varchar;unique;default:null"`
	Gender                  string          `gorm:"type:varchar"`
	Birthdate               time.Time       `gorm:"type:timestamp;default:null"`
	ProfilePicture          string          `gorm:"type:varchar;default:null"`
	WalletNumber            string          `gorm:"type:varchar;unique;default:null"`
	WalletPin               string          `gorm:"type:varchar;default:null"`
	Balance                 decimal.Decimal `gorm:"type:decimal;default:0"`
	SellerBalance           decimal.Decimal `gorm:"type:decimal;default:0"`
	ForgetPasswordToken     string          `gorm:"type:varchar;default:null"`
	ForgetPasswordExpiredAt time.Time       `gorm:"type:timestamp;default:null"`
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
	ID             int             `gorm:"primaryKey;not null,autoIncrement;serial" json:"id"`
	AccountID      int             `gorm:"foreignKey:AccountID;type:bigint;not null" json:"-"`
	Type           string          `gorm:"type:varchar" json:"type"`
	From           string          `gorm:"type:varchar;default:null" json:"from"`
	To             string          `gorm:"type:varchar;default:null" json:"to"`
	Amount         decimal.Decimal `gorm:"type:decimal" json:"amount"`
	ProductOrderID int             `gorm:"foreignKey:AccountID;type:bigint;default:null" json:"-"`
	CreatedAt      time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp" json:"-"`
	DeletedAt      time.Time       `gorm:"type:timestamp;default:null" json:"-"`
}

type SaleWalletTransactionHistories struct {
	ID             int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID      int             `gorm:"foreignKey:AccountID;type:bigint;not null"`
	Type           string          `gorm:"type:varchar"`
	From           string          `gorm:"type:varchar;default:null"`
	To             string          `gorm:"type:varchar;default:null"`
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
