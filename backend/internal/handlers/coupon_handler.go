package handlers

import (
	"coupon-system/internal/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	DB *sql.DB
}

func NewCouponHandler(db *sql.DB) *CouponHandler {
	return &CouponHandler{DB: db}
}

func (h *CouponHandler) CreateCoupon(c *gin.Context) {
	var coupon models.Coupon

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupon.Active = true

	query := `
	INSERT INTO coupons
	(code, discount_type, discount_value, max_uses, min_order_value, expires_at)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id, created_at
	`
	err := h.DB.QueryRow(
		query,
		coupon.Code,
		coupon.DiscountType,
		coupon.DiscountValue,
		coupon.MaxUses,
		coupon.MinOrderValue,
		coupon.ExpiresAt,
	).Scan(&coupon.ID, &coupon.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, coupon)
}

func (h *CouponHandler) GetCoupons(c *gin.Context) {
	rows, err := h.DB.Query(`
		SELECT id, code, discount_type, discount_value, max_uses,
		current_uses, min_order_value, expires_at, active, created_at
		FROM coupons
	`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	defer rows.Close()

	var coupons []models.Coupon

	for rows.Next() {
		var coupon models.Coupon

		err := rows.Scan(
			&coupon.ID,
			&coupon.Code,
			&coupon.DiscountType,
			&coupon.DiscountValue,
			&coupon.MaxUses,
			&coupon.CurrentUses,
			&coupon.MinOrderValue,
			&coupon.ExpiresAt,
			&coupon.Active,
			&coupon.CreatedAt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		coupons = append(coupons, coupon)
	}

	c.JSON(http.StatusOK, coupons)
}

func (h *CouponHandler) ValidateCoupon(c *gin.Context) {
	var req models.CouponValidationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		SELECT id, code, discount_type, discount_value,
		max_uses, current_uses, min_order_value, expires_at, active
		FROM coupons
		WHERE code = $1
	`
	var coupon models.Coupon

	err := h.DB.QueryRow(query, req.Code).Scan(
		&coupon.ID,
		&coupon.Code,
		&coupon.DiscountType,
		&coupon.DiscountValue,
		&coupon.MaxUses,
		&coupon.CurrentUses,
		&coupon.MinOrderValue,
		&coupon.ExpiresAt,
		&coupon.Active,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"valid":   false,
			"message": "coupon not found",
		})
		return
	}

	if !coupon.Active {
		c.JSON(http.StatusNotFound, gin.H{
			"valid":   false,
			"message": "coupon is not active",
		})
		return
	}

	if time.Now().After(coupon.ExpiresAt) {
		c.JSON(http.StatusNotFound, gin.H{
			"valid":   false,
			"message": "coupon has expired",
		})
		return
	}

	if req.OrderValue < coupon.MinOrderValue {
		c.JSON(http.StatusNotFound, gin.H{
			"valid":   false,
			"message": "order value does not meet minimum requirement",
		})
		return
	}

	var discount float64

	if coupon.DiscountType == "percentage" {
		discount = req.OrderValue * (coupon.DiscountValue / 100)
	} else {
		discount = coupon.DiscountValue
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":    true,
		"discount": discount,
	})
}
