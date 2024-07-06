// Initialise l'ensemble de l'API
package main

import (
	"PCS-API/controller"
	"PCS-API/middleware"
	"PCS-API/utils"
	"github.com/gin-gonic/gin"
)

// Initialise l'API
func main() {
	utils.LoadConfig()
	router := gin.Default()

	middleware.CORS(router)

	api := router.Group("/api")

	controller.Users(api)
	controller.Chat(api)
	controller.Sidebar(api)
	controller.Property(api)
	controller.Admin(api)
	controller.Traveler(api)
	controller.Property_image(api)
	controller.Service(api)
	controller.Reservation(api)
	controller.Ticket(api)

	err := router.Run(":" + utils.PortApp)
	if err != nil {
		return
	}
}
