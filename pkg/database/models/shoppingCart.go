package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint       `gorm:"not null"`
	Items     []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CartItem struct {
	gorm.Model
	CartId    uint           `gorm:"not null"`
	ProductId uint           `gorm:"not null"`
	Quantity  int            `gorm:"not null;default:1"`
	Price     float64        `gorm:"not null"`
	Discount  float64        `gorm:"default:0"`
	Cart      Cart           `gorm:"foreignKey:CartID"`
	Product   Product        `gorm:"foreignKey:ProductID"`
	Addons    datatypes.JSON `gorm:"type:json"`
}
