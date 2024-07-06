package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Ticket # Controller
//
// Ticket réceptionne toutes les requêtes ayant pour endpoint '/Ticket'
// Il les envoie aux fonctions services liés
func Ticket(api *gin.RouterGroup) {
	api.GET("/ticket", service.TicketGetAll)
	ticketGroup := api.Group("/ticket")
	ticketGroup.Use(middleware.AuthMiddleware())
	{
		ticketGroup.PUT("/:id", service.TicketUpdateById)
		ticketGroup.GET("/", service.TicketGetAll)
		ticketGroup.POST("/", service.TicketCreate)
	}
}
