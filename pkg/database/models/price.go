package models

import (
	"gorm.io/gorm"
)

type ProductPrice struct {
	gorm.Model
	ProductId  uint     `gorm:"index"`
	CurrencyId uint     `gorm:"index"`
	Amount     float64  `gorm:"not null"`
	Currency   Currency `gorm:"foreignKey:CurrencyId"`
}

type AddonPrice struct {
	gorm.Model
	AddonId    uint     `gorm:"index"`
	CurrencyId uint     `gorm:"index"`
	Amount     float64  `gorm:"not null"`
	Currency   Currency `gorm:"foreignKey:CurrencyId"`
}
