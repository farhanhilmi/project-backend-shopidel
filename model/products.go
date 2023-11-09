package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Products struct {
	ID          int    `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name        string `gorm:"type:varchar;not null"`
	Description string `gorm:"type:text;"`
	CategoryID  int    `gorm:"foreignKey:CategoryID;type:bigint;not null"`
	SellerID    int    `gorm:"foreignKey:SellerID;type:bigint;not null"`
	Category    Category
	// ProductImages     ProductImages
	// ProductVideos     ProductVideos
	// Variants          []Variant
	// Images            []ProductImages
	HazardousMaterial bool            `gorm:"type:boolean;not null"`
	Weight            decimal.Decimal `gorm:"type:decimal;not null"`
	Size              decimal.Decimal `gorm:"type:decimal;not null"`
	IsNew             bool            `gorm:"type:boolean;not null"`
	InternalSKU       string          `gorm:"type:varchar;"`
	ViewCount         int             `gorm:"type:int;default:0"`
	IsActive          bool            `gorm:"type:boolean;not null"`
	CreatedAt         time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt         time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt         time.Time       `gorm:"type:timestamp;default:null"`
}

type ProductVariants struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductID int       `gorm:"foreignKey:ProductID;type:bigint;not null"`
	Name      string    `gorm:"type:varchar;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type ProductVariantSelections struct {
	ID               int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductVariantID int       `gorm:"foreignKey:ProductVariantID;type:bigint;not null"`
	Name             string    `gorm:"type:varchar;not null"`
	CreatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt        time.Time `gorm:"type:timestamp;default:null"`
}

type ProductVariantSelectionCombinations struct {
	ID                         int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductID                  int             `gorm:"foreignKey:ProductID;type:bigint;not null"`
	ProductVariantSelectionID1 int             `gorm:"foreignKey:ProductVariantSelectionID1;type:bigint"`
	ProductVariantSelectionID2 int             `gorm:"foreignKey:ProductVariantSelectionID2;type:bigint"`
	Price                      decimal.Decimal `gorm:"type:decimal;not null"`
	Stock                      int             `gorm:"type:int;not null"`
	PictureURL                 string          `gorm:"type:varchar;not null"`
	CreatedAt                  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt                  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt                  time.Time       `gorm:"type:timestamp;default:null"`
}

type Category struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name      string    `gorm:"type:varchar;not null"`
	Level     int       `gorm:"type:int;not null"`
	Parent    int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type ProductImages struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductID int       `gorm:"foreignKey:ProductID;type:bigint;not null"`
	URL       string    `gorm:"type:varchar;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type ProductVideos struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductID int       `gorm:"foreignKey:ProductID;type:bigint;not null"`
	URL       string    `gorm:"type:varchar;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type ProductCombinationVariant struct {
	ID    int `gorm:"serial"`
	Stock int `gorm:"type:int;not null"`
}

type FavoriteProducts struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductID int       `gorm:"foreignKey:ProductID;type:bigint;not null"`
	AccountID int       `gorm:"foreignKey:AccountID;type:bigint;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type FavoriteProductList struct {
	ID          int             `json:"id"`
	ProductID   int             `json:"product_id"`
	AccountID   int             `json:"account_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	PictureURL  string          `json:"picture_url"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"-"`
	DeletedAt   time.Time       `json:"-"`
}
