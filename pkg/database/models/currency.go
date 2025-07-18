package models

import "gorm.io/gorm"

type Currency struct {
	Code   string `gorm:"size:3;uniqueIndex;not null"` // IRR, EUR, USD, GBP etc.
	Symbol string `gorm:"size:5"`                      // ﷼, €, $ etc.
	Name   string
	gorm.Model
}
