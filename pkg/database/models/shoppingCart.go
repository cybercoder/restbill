package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId uint       `gorm:"not null;index"`
	Items  []CartItem `gorm:"foreignKey:CartId"`
}

type CartItem struct {
	gorm.Model
	CartId    uint           `gorm:"index"`
	ProductId uint           `gorm:"index"`
	Quantity  int            `gorm:"not null;default:1"`
	Price     float64        `gorm:"not null"`
	Discount  float64        `gorm:"default:0"`
	Cart      Cart           `gorm:"foreignKey:CartId"`
	Product   Product        `gorm:"foreignKey:ProductId"`
	Addons    datatypes.JSON `gorm:"type:json"`
}
