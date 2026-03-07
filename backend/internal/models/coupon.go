package models

import "time"

type Coupon struct {
	ID            uint      `json:"id"`
	Code          string    `json:"code"`
	DiscountType  string    `json:"discount_type"`
	DiscountValue float64   `json:"discount_value"`
	MaxUses       int       `json:"max_uses"`
	CurrentUses   int       `json:"current_uses"`
	MinOrderValue float64   `json:"min_order_value"`
	ExpiresAt     time.Time `json:"expires_at"`
	Active        bool      `json:"active"`
	CreatedAt     time.Time `json:"created_at"`
}
