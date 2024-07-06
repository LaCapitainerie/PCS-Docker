package controller

import (
	"PCS-API/middleware"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Property # Controller
//
// Property réceptionne toutes les requêtes ayant pour endpoint '/property'
// Il les envoie aux fonctions services liés
func Property(api *gin.RouterGroup) {
	property := api.Group("/property")
	property.Use(middleware.AuthMiddleware())
	{
		property.GET("", service.GetAllProperty)
		property.POST("", service.PostAProperty)
		property.DELETE("/:id", service.PropertyDeleteById)
		property.GET("/:id", service.GetPropertyById)
		property.PUT("/:id", service.PutPropertyById)
	}
}
