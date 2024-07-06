package controller

import (
	"PCS-API/middleware"
	"github.com/gin-gonic/gin"
)

func Reservation(api *gin.RouterGroup) {
	reservation := api.Group("/reservation")
	reservation.Use(middleware.AuthMiddleware())
	reservationProperty(reservation)
	reservationCheckout(reservation)
}
