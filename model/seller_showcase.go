package model

import (
	"time"
)

type SellerShowcase struct {
	ID               int `gorm:"primaryKey;not null,autoIncrement;serial"`
	SellerId         int
	Name             string
	ShowcaseProducts []SellerShowcaseProduct
	CreatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt        time.Time `gorm:"type:timestamp;default:null"`
}

type SellerShowcaseProduct struct {
	ID               int `gorm:"primaryKey;not null,autoIncrement;serial"`
	SellerShowcaseId int
	SellerShowcase   SellerShowcase
	ProductId        int
	Product          Products
	CreatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt        time.Time `gorm:"type:timestamp;default:null"`
}
