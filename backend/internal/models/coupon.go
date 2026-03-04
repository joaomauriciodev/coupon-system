package models

import "time"

type Coupon struct {
	ID            uint   `gorm:"primaryKey"`
	Code          string `gorm:"uniqueIndex"`
	DiscountType  string
	DiscountValue float64
	MaxUses       int
	CurrentUses   int
	MinOrderValue float64
	ExpiresAt     time.Time
	Active        bool
	CreatedAt     time.Time
}
