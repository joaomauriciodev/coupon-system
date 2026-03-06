package handlers

import (
	"coupon-system/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CouponHandler struct {
	DB *gorm.DB
}

func NewCouponHandler(db *gorm.DB) *CouponHandler {
	return &CouponHandler{DB: db}
}

func (h *CouponHandler) CreateCoupon(c *gin.Context) {
	var coupon models.Coupon

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupon.Active = true

	result := h.DB.Create(&coupon)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Coupon code already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coupon"})
		return
	}

	c.JSON(http.StatusCreated, coupon)
}

func (h *CouponHandler) GetCoupons(c *gin.Context) {
	var coupons []models.Coupon

	h.DB.Find(&coupons)

	c.JSON(http.StatusOK, coupons)
}
