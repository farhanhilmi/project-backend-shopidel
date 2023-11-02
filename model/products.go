package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Products struct {
	ID                int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name              string          `gorm:"type:varchar;not null"`
	Description       string          `gorm:"type:text;"`
	CategoryID        int             `gorm:"foreignKey:CategoryID;type:bigint;not null"`
	HazardousMaterial bool            `gorm:"type:bolean;not null"`
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

type Category struct {
	ID        int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name      string    `gorm:"type:varchar;not null"`
	Level     int       `gorm:"type:int;not null"`
	Parent    int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type ProductVariantSelectionCombinations struct {
	ID                         int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	ProductID                  int             `gorm:"foreignKey:ProductID;type:bigint;not null"`
	ProductVariantSelectionID1 int             `gorm:"foreignKey:ProductVariantSelectionID1;type:bigint;not null"`
	ProductVariantSelectionID2 int             `gorm:"foreignKey:ProductVariantSelectionID2;type:bigint"`
	Price                      decimal.Decimal `gorm:"type:decimal;not null"`
	Stock                      int             `gorm:"type:int"`
	PictureURL                 string          `gorm:"type:varchar"`
	CreatedAt                  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt                  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt                  time.Time       `gorm:"type:timestamp;default:null"`
}
