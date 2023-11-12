package model

import (
	"time"
)

type Province struct {
	ID                   int    `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name                 string `gorm:"type:varchar;not null"`
	RajaOngkirProvinceId int
	CreatedAt            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt            time.Time `gorm:"type:timestamp;default:null"`
}

type District struct {
	ID                   int    `gorm:"primaryKey;not null,autoIncrement;serial"`
	Name                 string `gorm:"type:varchar;not null"`
	ProvinceId           int    `gorm:"foreignKey:ProvinceId;not null"`
	RajaOngkirDistrictId string
	CreatedAt            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt            time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt            time.Time `gorm:"type:timestamp;default:null"`
}
