package handlers

import (
	"coupon-system/internal/models"
	"net/http"

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

	h.DB.Create(&coupon)

	c.JSON(http.StatusCreated, coupon)
}

func (h *CouponHandler) GetCoupons(c *gin.Context) {
	var coupons []models.Coupon

	h.DB.Find(&coupons)

	c.JSON(http.StatusOK, coupons)
}
