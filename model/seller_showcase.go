package model

import (
	"time"
)

type SellerShowcase struct {
	ID        int `gorm:"primaryKey;not null,autoIncrement;serial"`
	SellerId  int
	Name      string
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null"`
}

type SellerShowcaseProduct struct {
	ID               int `gorm:"primaryKey;not null,autoIncrement;serial"`
	SellerShowcaseId int
	ProductId        int
	CreatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt        time.Time `gorm:"type:timestamp;default:null"`
}
