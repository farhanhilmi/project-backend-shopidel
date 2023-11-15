package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductOrders struct {
	ID            int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID     int             `gorm:"foreignKey:AccountID;type:bigint;not null"`
	CourierName   string          `gorm:"type:varchar"`
	ProductName   string          `gorm:"type:varchar"`
	DeliveryFee   decimal.Decimal `gorm:"type:decimal;not null"`
	Province      string          `gorm:"type:varchar"`
	District      string          `gorm:"type:varchar"`
	SubDistrict   string          `gorm:"type:varchar"`
	Kelurahan     string          `gorm:"type:varchar"`
	ZipCode       string          `gorm:"type:varchar"`
	AddressDetail string          `gorm:"type:text"`
	Status        string          `gorm:"type:varchar"`
	Notes         string          `gorm:"type:varchar;default:null"`
	CreatedAt     time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt     time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt     time.Time       `gorm:"type:timestamp;default:null"`
}

type Couriers struct {
	ID          int       `gorm:"primaryKey;not null,autoIncrement;serial" json:"id"`
	Name        string    `gorm:"type:varchar;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp" json:"-"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp" json:"-"`
	DeletedAt   time.Time `gorm:"type:timestamp;default:null" json:"-"`
}

type ProductOrderDetails struct {
	ID              int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductOrderID  int             `gorm:"foreignKey:ProductOrderID;type:bigint;not null"`
	ProductID       int             `gorm:"foreignKey:ProductID;type:bigint;not null"`
	Quantity        int             `gorm:"type:int;not null"`
	VariantName     string          `gorm:"type:varchar"`
	IndividualPrice decimal.Decimal `gorm:"type:decimal;not null"`
	CreatedAt       time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt       time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt       time.Time       `gorm:"type:timestamp;default:null"`
}

type ProductOrderSeller struct {
	ID                                   int
	ProductVariantSelectionCombinationID int
	ProductStock                         int
	ProductID                            int
	AccountID                            int
	SellerID                             int
	IndividualPrice                      decimal.Decimal
	Quantity                             int
	Status                               string
}

type ProductOrderHistories struct {
	ID                   int
	ProductOrderDetailID int
	ProductName          string
	CourierName          string
	VariantName          string
	Quantity             int
	Status               string
	ProductID            int
	IndividualPrice      decimal.Decimal
	PictureURL           string
	Feedback             string
	Rating               int
	ReviewID             int
	ReviewImageURL       string
	ShopName             string
	ReviewCreatedAt      time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}

type ProductOrderDetail struct {
	ID              int
	ProductName     string
	Quantity        int
	Status          string
	ProductID       int
	IndividualPrice decimal.Decimal
	PictureURL      string
	Feedback        string
	Rating          int
	ReviewID        int
	ShopName        string
	Province        string
	District        string
	ZipCode         string
	SubDistrict     string
	Kelurahan       string
	Detail          string
	DeliveryFee     decimal.Decimal
	ReviewCreatedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type ProductOrderReviews struct {
	ID                   int
	AccountID            int
	ProductID            int
	ProductOrderDetailID int
	Feedback             string
	Rating               int
	VariantName          string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}

type ProductOrderReviewImages struct {
	ID              int
	ProductReviewID int
	ImageURL        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
