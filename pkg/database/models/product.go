package models

import "gorm.io/gorm"

type Category struct {
	Name        string `gorm:"unique;not null"` // "Virtual Machines", "CDN", "Storage"
	Description string
	gorm.Model
	Products []Product `gorm:"foreignKey:CategoryId"`
	Addons   []Addon   `gorm:"foreignKey:CategoryId"`
}

type Product struct {
	ID          uint   `gorm:"primary_key"`
	CategoryId  uint   `gorm:"index"`
	Name        string `gorm:"not null"`
	Description string
	Category    Category       `gorm:"foreignKey:CategoryId"`
	Price       []ProductPrice `gorm:"foreignKey:ProductId"`
	gorm.Model
}

type Addon struct {
	ID          uint   `gorm:"primary_key"`
	CategoryId  uint   `gorm:"index"`
	Name        string `gorm:"unique;not null"`
	Description string
	Category    Category     `gorm:"foreignKey:CategoryId"`
	Price       []AddonPrice `gorm:"foreignKey:AddonId"`
	gorm.Model
}
