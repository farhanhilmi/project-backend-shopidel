package model

import (
	"time"
)

type SellerPageSelectedCategory struct {
	ID         int `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountId  int
	CategoryId int
	CreatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt  time.Time `gorm:"type:timestamp;default:null"`
}
