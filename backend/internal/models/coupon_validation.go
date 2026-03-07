package models

type CouponValidationRequest struct {
	Code       string  `json:"code"`
	OrderValue float64 `json:"order_value"`
}
