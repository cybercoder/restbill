package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UserId        uint
	Status        string
	TotalAmount   float64
	Tax           float64
	TaxRate       float64
	Credit        float64
	PaymentMethod string
	Notes         string
	Items         []InvoiceItem `gorm:"foreignKey:InvoiceId"`
	DueDate       time.Time
	DatePaid      time.Time
	ExpiresAt     time.Time
}

type InvoiceItem struct {
	gorm.Model
	InvoiceId   uint          `gorm:"index"`
	ParentId    uint          `gorm:"index"`
	Invoice     Invoice       `gorm:"foreignKey:InvoiceId"`
	Items       []InvoiceItem `gorm:"foreignKey:ParentId"`
	Description string
	Amount      float64
	Taxed       bool `gorm:"default:true"`
}
