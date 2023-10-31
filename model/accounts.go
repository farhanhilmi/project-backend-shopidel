package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Accounts struct {
	ID                      int             `gorm:"primaryKey;not null,autoIncrement;serial"`
	FullName                string          `gorm:"type:varchar;not null"`
	Username                string          `gorm:"type:varchar"`
	Email                   string          `gorm:"type:email;not null"`
	PhoneNumber             string          `gorm:"type:varchar"`
	Password                string          `gorm:"type:varchar"`
	ShopName                string          `gorm:"type:varchar"`
	Gender                  string          `gorm:"type:varchar"`
	Birthdate               time.Time       `gorm:"type:timestamp"`
	ProfilePicture          string          `gorm:"type:varchar"`
	WalletNumber            string          `gorm:"type:varchar"`
	WalletPin               string          `gorm:"type:int"`
	Balance                 decimal.Decimal `gorm:"type:decimal;default:0"`
	ForgetPasswordToken     string          `gorm:"type:varchar"`
	ForgetPasswordExpiredAt time.Time       `gorm:"type:timestamp"`
}
