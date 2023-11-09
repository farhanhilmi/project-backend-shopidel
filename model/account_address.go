package model

import "time"

type AccountAddress struct {
	ID                   int       `gorm:"primaryKey;not null,autoIncrement;serial"`
	AccountID            int       `gorm:"type:bigint;not null"`
	Province             string    `gorm:"type:varchar;not null"`
	District             string    `gorm:"type:varchar;not null"`
	RajaOngkirDistrictId string    `gorm:"type:varchar;not null"`
	SubDistrict          string    `gorm:"type:varchar;not null"`
	Kelurahan            string    `gorm:"type:varchar;not null"`
	ZipCode              string    `gorm:"type:varchar;not null"`
	Detail               string    `gorm:"type:text"`
	IsBuyerDefault       bool      `gorm:"type:boolean"`
	IsSellerDefault      bool      `gorm:"type:boolean"`
	CreatedAt            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt            time.Time `gorm:"type:timestamp;default:null"`
}
