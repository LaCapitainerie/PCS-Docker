package controller

import (
	"PCS-API/middleware"
	"PCS-API/models"
	"PCS-API/service"

	"github.com/gin-gonic/gin"
)

// Admin # Controller
//
// Admin réceptionne toutes les requêtes ayant pour endpoint '/Admin'
// Il les envoie aux fonctions services liés
func Admin(api *gin.RouterGroup) {
	api.GET("/admin", service.GetAllAdmin)
	administrationGroup := api.Group("/administration")
	administrationGroup.Use(middleware.AuthMiddleware())
	administrationGroup.Use(middleware.BlockTypeMiddleware(models.AdminType))
	{
		administrationGroup.POST("/login", service.LoginAdmin)
		administrationGroup.POST("/register", service.AdminRegister)
	}
}
