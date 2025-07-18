package models

import (
	"gorm.io/gorm"
)

type ProductPrice struct {
	gorm.Model
	ProductID  uint     `gorm:"index"`
	CurrencyID uint     `gorm:"index"`
	Amount     float64  `gorm:"not null"`
	Currency   Currency `gorm:"foreignKey:CurrencyID"`
}

type AddonPrice struct {
	gorm.Model
	AddonID    uint     `gorm:"index"`
	CurrencyID uint     `gorm:"index"`
	Amount     float64  `gorm:"not null"`
	Currency   Currency `gorm:"foreignKey:CurrencyID"`
}
