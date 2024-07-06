package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

func Service(api *gin.RouterGroup) {
	api.GET("/service/all", service.ServiceGetAll)
	serviceManagement := api.Group("/service/management")
	serviceManagement.Use(middleware.AuthMiddleware())
	serviceManagement.Use(middleware.BlockTypeMiddleware(models.ProviderType))
	{
		serviceManagement.POST("", service.ServiceCreateNewService)
		serviceManagement.PUT("/:id", service.ServiceUpdate)
		serviceManagement.DELETE("/:id", service.ServiceDelete)
	}
}
