package models

import "gorm.io/gorm"

type Category struct {
	Name        string `gorm:"unique;not null"` // "Virtual Machines", "CDN", "Storage"
	Description string
	gorm.Model
}

type Product struct {
	CategoryID  uint   `gorm:"index"`
	Name        string `gorm:"not null"`
	Description string
	Category    Category       `gorm:"foreignKey:CategoryID"`
	Price       []ProductPrice `gorm:"foreignKey:ProductID"`
	gorm.Model
}

type Addon struct {
	CategoryID  uint   `gorm:"index"`
	Name        string `gorm:"unique;not null"`
	Description string
	Category    Category     `gorm:"foreignKey:CategoryID"`
	Price       []AddonPrice `gorm:"foreignKey:AddonID"`
	gorm.Model
}
