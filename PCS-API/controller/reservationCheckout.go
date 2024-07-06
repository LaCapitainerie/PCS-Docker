package controller

import (
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Reservation is a function to group all reservation routes
// @Summary Group all reservation routes
// @Tags Reservation
// @Accept json
// @Produce json
// @Param api path string true "API"
// @Param reservation path string true "Reservation"
// @Param id path string true "ID"
// @Param quantity path string true "Quantity"
// @Router /api/reservation/checkout/session/{id}/{quantity} [post]
func reservationCheckout(reservation *gin.RouterGroup) {
	reservationCheckoutGroup := reservation.Group("/checkout")
	{
		// :id is the stripe price id
		// :quantity is the quantity of the product
		reservationCheckoutGroup.POST("/session/:id/:quantity", service.CheckoutCreateSession)
	}
}
