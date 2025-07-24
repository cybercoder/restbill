package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId uint       `gorm:"not null;index"`
	Items  []CartItem `gorm:"foreignKey:CartId"`
}

type CartItem struct {
	gorm.Model
	CartId    uint             `gorm:"index"`
	ProductId uint             `gorm:"index"`
	Quantity  uint             `gorm:"not null;default:1"`
	Price     float64          `gorm:"not null"`
	Discount  float64          `gorm:"default:0"`
	Cart      Cart             `gorm:"foreignKey:CartId"`
	Product   Product          `gorm:"foreignKey:ProductId"`
	Addons    []CartItemAddons `gorm:"foreignKey:CartItemId"`
}

type CartItemAddons struct {
	gorm.Model
	CartItemId uint     `gorm:"index"`
	AddonId    uint     `gorm:"index"`
	Quantity   uint     `gorm:"not null;default:1"`
	Price      float64  `gorm:"not null"`
	Discount   float64  `gorm:"default:0"`
	CartItem   CartItem `gorm:"foreignKey:CartItemId"`
	Addon      Addon    `gorm:"foreignKey:AddonId"`
}
