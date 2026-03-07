package main

import (
	"coupon-system/internal/database"
	"coupon-system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	router := gin.Default()

	couponHandler := handlers.NewCouponHandler(db)

	router.GET("/coupons", couponHandler.GetCoupons)
	router.POST("/coupons", couponHandler.CreateCoupon)

	router.Run(":8080")

}
