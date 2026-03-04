package main

import (
	"coupon-system/internal/database"
	"coupon-system/internal/handlers"
	"coupon-system/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	db.AutoMigrate(&models.Coupon{})

	router := gin.Default()

	couponHandler := handlers.NewCouponHandler(db)

	router.GET("/coupons", couponHandler.GetCoupons)
	router.POST("/coupons", couponHandler.CreateCoupon)

	router.Run(":8080")

}
