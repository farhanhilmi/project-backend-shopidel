package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductOrders struct {
	ID            int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	CourierID     int             `gorm:"foreignKey:CourierID;type:bigint;not null"`
	AccountID     int             `gorm:"foreignKey:AccountID;type:bigint;not null"`
	DeliveryFee   decimal.Decimal `gorm:"type:decimal;not null"`
	Province      string          `gorm:"type:varchar"`
	District      string          `gorm:"type:varchar"`
	SubDistrict   string          `gorm:"type:varchar"`
	Kelurahan     string          `gorm:"type:varchar"`
	ZipCode       string          `gorm:"type:varchar"`
	AddressDetail string          `gorm:"type:text"`
	Status        string          `gorm:"type:varchar"`
	CreatedAt     time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt     time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt     time.Time       `gorm:"type:timestamp;default:null"`
}

type Couriers struct {
	ID          int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name        string    `gorm:"type:varchar;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt   time.Time `gorm:"type:timestamp;default:null"`
}

type ProductOrderDetails struct {
	ID                                   int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductOrderID                       int             `gorm:"foreignKey:ProductOrderID;type:bigint;not null"`
	ProductVariantSelectionCombinationID int             `gorm:"foreignKey:ProductVariantSelectionCombinationID;type:bigint;not null"`
	Quantity                             int             `gorm:"type:int;not null"`
	IndividualPrice                      decimal.Decimal `gorm:"type:decimal;not null"`
	CreatedAt                            time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt                            time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt                            time.Time       `gorm:"type:timestamp;default:null"`
}
