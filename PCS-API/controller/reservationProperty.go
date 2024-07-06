package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

func reservationProperty(reservation *gin.RouterGroup) {
	reservation.GET("/property/allreservation/:id", service.ReservationGetAllOfAProperty)
	reservation.PUT("/property/annulation/:id", service.ReservationPropertyAnnulationWithAId)
	reservationPropertyGroup := reservation.Group("/property")
	reservationPropertyGroup.Use(middleware.BlockTypeMiddleware(models.TravelerType))
	{
		reservationPropertyGroup.POST("/validation/:id", service.ReservationValidationPaiement)
		reservationPropertyGroup.PUT("/report", service.ReservationPropertyReportReservation)
	}
}
